import copy
import json
import boto3
import os
import logging
import cfnresponse  # AWS-provided helper for sending responses

# Set up logging
logger = logging.getLogger()
logger.setLevel(logging.INFO)


def validate_input(event):
    required_fields = ['Name', 'RoleArn', 'ResourcesVpcConfig']
    for field in required_fields:
        if field not in event['ResourceProperties']:
            raise ValueError(f"Missing required field: {field}")


def delete_cluster(cluster_name):
    eks_endpoint = os.environ.get('AWS_ENDPOINT_URL_EKS', 'https://api.beta.us-west-2.wesley.amazonaws.com')
    eks_client = boto3.client('eks', endpoint_url=eks_endpoint)
    logger.info(f"Deleting EKS cluster: {cluster_name}")
    eks_client.delete_cluster(name=cluster_name)
    logger.info(f"EKS cluster deleted: {cluster_name}")


def convert_keys_to_lowercase_first_letter(d):
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
    if isinstance(d, (dict, list)):
        iterable = d.items() if isinstance(d, dict) else enumerate(d)
        for k, v in iterable:
            if isinstance(v, str) and v.lower() in {"true", "false"}:
                d[k] = v.lower() == "true"
            elif isinstance(v, (dict, list)):
                replace_boolean_strings(v)


def handler(event, context):
    try:
        logger.info("Received event: " + json.dumps(event, default=str))

        # Validate that the invocation is from CloudFormation
        if 'RequestType' not in event:
            raise ValueError("Invalid invocation source. This Lambda function can only be invoked by CloudFormation.")

        # Handle Delete event
        if event['RequestType'] == 'Delete':
            cluster_name = event['ResourceProperties']['Name']
            delete_cluster(cluster_name)
            cfnresponse.send(event, context, cfnresponse.SUCCESS, {"Message": "Resource deleted"})
            return {
                'PhysicalResourceId': event['PhysicalResourceId']
            }

        # Validate input
        validate_input(event)

        # Prepare the create_cluster request payload from the custom resource properties
        create_cluster_payload = convert_keys_to_lowercase_first_letter(
            copy.deepcopy(event['ResourceProperties']))

        # delete ServiceToken from create_cluster_payload
        del create_cluster_payload['serviceToken']
        replace_boolean_strings(create_cluster_payload)

        # Create the EKS cluster
        eks_endpoint = os.environ.get('AWS_ENDPOINT_URL_EKS', 'https://api.beta.us-west-2.wesley.amazonaws.com')
        eks_client = boto3.client('eks', endpoint_url=eks_endpoint)
        logger.info("Creating EKS cluster with payload: " + json.dumps(create_cluster_payload, default=str))
        response = eks_client.create_cluster(**create_cluster_payload)
        logger.info("EKS cluster created: " + json.dumps(response, default=str))

        eventData = {
            "PhysicalResourceId": response['cluster']['arn'],
            "ClusterName": response['cluster']['name'],
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
