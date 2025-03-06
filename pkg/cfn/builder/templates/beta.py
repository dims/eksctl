import json
import boto3
import os
import logging
from botocore.config import Config
import cfnresponse  # AWS-provided helper for sending responses

# Set up logging
logger = logging.getLogger()
logger.setLevel(logging.INFO)


def validate_input(event):
    required_fields = ['ClusterName', 'RoleArn', 'ResourcesVpcConfig']
    for field in required_fields:
        if field not in event['ResourceProperties']:
            raise ValueError(f"Missing required field: {field}")


def delete_cluster(cluster_name):
    eks_endpoint = os.environ.get('AWS_ENDPOINT_URL_EKS', 'https://api.beta.us-west-2.wesley.amazonaws.com')
    eks_client = boto3.client('eks', endpoint_url=eks_endpoint)
    logger.info(f"Deleting EKS cluster: {cluster_name}")
    eks_client.delete_cluster(name=cluster_name)
    logger.info(f"EKS cluster deleted: {cluster_name}")


def handler(event, context):
    try:
        logger.info("Received event: " + json.dumps(event, default=str))

        # Validate that the invocation is from CloudFormation
        if 'RequestType' not in event:
            raise ValueError("Invalid invocation source. This Lambda function can only be invoked by CloudFormation.")

        # Handle Delete event
        if event['RequestType'] == 'Delete':
            cluster_name = event['ResourceProperties']['ClusterName']
            delete_cluster(cluster_name)
            cfnresponse.send(event, context, cfnresponse.SUCCESS, {"Message": "Resource deleted"})
            return {
                'PhysicalResourceId': event['PhysicalResourceId']
            }

        # Validate input
        validate_input(event)

        # Prepare the create_cluster request payload
        properties = event['ResourceProperties']
        create_cluster_payload = {
            'name': properties['ClusterName'],
            'roleArn': properties['RoleArn'],
            'resourcesVpcConfig': {
                'subnetIds': properties['ResourcesVpcConfig']['SubnetIds'],
                'securityGroupIds': properties['ResourcesVpcConfig']['SecurityGroupIds']
            }
        }

        # Add optional properties if they exist in the request
        if 'Version' in properties:
            create_cluster_payload['version'] = properties['Version']

        # Create the EKS cluster
        eks_endpoint = os.environ.get('AWS_ENDPOINT_URL_EKS', 'https://api.beta.us-west-2.wesley.amazonaws.com')
        eks_client = boto3.client('eks', endpoint_url=eks_endpoint)
        logger.info("Creating EKS cluster with payload: " + json.dumps(create_cluster_payload, default=str))
        response = eks_client.create_cluster(**create_cluster_payload)
        logger.info("EKS cluster created: " + json.dumps(response, default=str))

        cfnresponse.send(event, context, cfnresponse.SUCCESS, {"PhysicalResourceId": response['cluster']['arn']})
        # Return the cluster ARN as the PhysicalResourceId
        return {
            'PhysicalResourceId': response['cluster']['arn'],
            'Data': json.dumps(response, default=str)
        }
    except Exception as e:
        logger.error("Error: " + str(e))
        cfnresponse.send(event, context, cfnresponse.FAILED, {"Message": str(e)})
