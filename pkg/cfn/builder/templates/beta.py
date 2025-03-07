import copy
import json
import boto3
import os
import logging
import time
import cfnresponse  # AWS-provided helper for sending responses

# Set up logging
logger = logging.getLogger()
logger.setLevel(logging.INFO)


def validate_input(event):
    """
    Validate that all required fields are present in the event.
    """
    required_fields = ['Name', 'RoleArn', 'ResourcesVpcConfig',
                       'IAMPrincipalArn', 'STSRoleArn']
    for field in required_fields:
        if field not in event['ResourceProperties']:
            raise ValueError(f"Missing required field: {field}")


def delete_cluster(eks_client, cluster_name):
    """
    Delete an EKS cluster.
    """
    logger.info(f"Deleting EKS cluster: {cluster_name}")
    eks_client.delete_cluster(name=cluster_name)
    logger.info(f"EKS cluster deleted: {cluster_name}")


def convert_keys_to_lowercase_first_letter(d):
    """
    Convert the first character of dictionary keys to lowercase.
    """
    if not isinstance(d, dict):
        return d  # Return as-is if it's not a dictionary

    new_dict = {}
    for key, value in d.items():
        # Convert the first character of the key to lowercase
        new_key = key[:1].lower() + key[1:] if key else key

        # Recursively process nested dictionaries
        if isinstance(value, dict):
            new_dict[new_key] = convert_keys_to_lowercase_first_letter(value)
        # Process lists only if they contain dictionaries
        elif isinstance(value, list):
            new_dict[new_key] = [
                convert_keys_to_lowercase_first_letter(item) if isinstance(item, dict) else item
                for item in value
            ]
        else:
            new_dict[new_key] = value
    return new_dict


def replace_boolean_strings(d):
    """
    Replace string representations of booleans with actual booleans.
    """
    if isinstance(d, (dict, list)):
        iterable = d.items() if isinstance(d, dict) else enumerate(d)
        for k, v in iterable:
            if isinstance(v, str) and v.lower() in {"true", "false"}:
                d[k] = v.lower() == "true"
            elif isinstance(v, (dict, list)):
                replace_boolean_strings(v)


def wait_for_cluster_creation(eks_client, cluster_name):
    """
    Wait for the EKS cluster to become ACTIVE.
    """
    while True:
        response = eks_client.describe_cluster(name=cluster_name)
        status = response['cluster']['status']
        if status == 'ACTIVE':
            logger.info(f"EKS cluster {cluster_name} is now ACTIVE.")
            return response['cluster']
        elif status == 'FAILED':
            raise Exception(f"EKS cluster {cluster_name} creation failed.")
        else:
            logger.info(f"EKS cluster {cluster_name} status: {status}. Waiting...")
            time.sleep(10)  # Wait 10 seconds before polling again


def create_access_entry(eks_client, principal_arn, username, cluster_name):
    """
    Create an access entry for an IAM principal in an EKS cluster.
    """
    logger.info(f"Creating access entry in EKS cluster: {cluster_name}")
    response = eks_client.create_access_entry(
        clusterName=cluster_name,
        principalArn=principal_arn,
        username=username
    )
    logger.info("Access entry created successfully:")
    logger.info("Access entry resposne: " + json.dumps(response, default=str))

    # Associate the admin access policy
    policy_arn = 'arn:aws:eks::aws:cluster-access-policy/AmazonEKSClusterAdminPolicy'
    response = eks_client.associate_access_policy(
        clusterName=cluster_name,
        principalArn=principal_arn,
        policyArn=policy_arn,
        accessScope={
            'type': 'cluster',  # Scope type (cluster or namespace)
            'namespaces': []    # Leave empty for cluster-wide access
        }
    )


def handler(event, context):
    """
    Lambda function handler for CloudFormation custom resource.
    """
    try:
        logger.info("Received event: " + json.dumps(event, default=str))

        # Validate that the invocation is from CloudFormation
        if 'RequestType' not in event:
            raise ValueError("Invalid invocation source. This Lambda function can only be invoked by CloudFormation.")

        eks_endpoint = os.environ.get('AWS_ENDPOINT_URL_EKS', 'https://api.beta.us-west-2.wesley.amazonaws.com')
        eks_client = boto3.client('eks', endpoint_url=eks_endpoint)

        cluster_name = event['ResourceProperties']['Name']

        # Handle Delete event
        if event['RequestType'] == 'Delete':
            delete_cluster(eks_client, cluster_name)
            cfnresponse.send(event, context, cfnresponse.SUCCESS, {"Message": "Resource deleted"})
            return {
                'PhysicalResourceId': event['PhysicalResourceId']
            }

        # Validate input
        validate_input(event)

        iam_principal_arn = event['ResourceProperties']['IAMPrincipalArn']
        sts_role_arn = event['ResourceProperties']['STSRoleArn']

        # Prepare the create_cluster request payload from the custom resource properties
        create_cluster_payload = convert_keys_to_lowercase_first_letter(
            copy.deepcopy(event['ResourceProperties']))

        # cleanup create_cluster_payload as not all fields can be sent to EKS create_cluster API
        del create_cluster_payload['serviceToken']
        del create_cluster_payload['iAMPrincipalArn']
        del create_cluster_payload['sTSRoleArn']
        replace_boolean_strings(create_cluster_payload)

        # create and wait for the eks cluster
        cluster_details, response = create_cluster(eks_client, cluster_name, create_cluster_payload)

        # Create an access entry for the EKS cluster
        create_access_entry(eks_client, iam_principal_arn, sts_role_arn, cluster_name)

        # Extract required attributes
        eventData = {
            "Arn": cluster_details['arn'],
            "PhysicalResourceId": cluster_details['arn'],
            "ClusterName": cluster_details['name'],
            "ClusterSecurityGroupId": cluster_details['resourcesVpcConfig']['clusterSecurityGroupId'],
            "CertificateAuthorityData": cluster_details['certificateAuthority']['data'],
            "Endpoint": cluster_details['endpoint'],
        }

        cfnresponse.send(event, context, cfnresponse.SUCCESS,
                         eventData)
        # Return the cluster ARN as the PhysicalResourceId
        return {
            'PhysicalResourceId': response['cluster']['arn'],
            "ClusterName": response['cluster']['name'],
            'Data': json.dumps(response, default=str)
        }
    except Exception as e:
        logger.error("Error: " + str(e))
        cfnresponse.send(event, context, cfnresponse.FAILED, {"Message": str(e)})


def create_cluster(eks_client, cluster_name, create_cluster_payload):
    # Create the EKS cluster
    logger.info("Creating EKS cluster with payload: " + json.dumps(create_cluster_payload, default=str))
    response = eks_client.create_cluster(**create_cluster_payload)
    logger.info("EKS cluster created: " + json.dumps(response, default=str))
    # Wait for the cluster to become ACTIVE
    cluster_details = wait_for_cluster_creation(eks_client, cluster_name)
    return cluster_details, response
