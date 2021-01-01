#!/bin/sh

set -e

AWS_REGION="ap-northeast-1"
PROFILE="merpochi-dev"

SSM_ACTIVATION=$(aws ssm create-activation \
                  --iam-role "service-role/AmazonEC2RunCommandRoleForManagedInstances" \
                  --registration-limit 1 \
                  --region $AWS_REGION \
                  --default-instance-name "merpochi-fargate-container" \
                  --profile $PROFILE)

export SSM_ACTIVATION_ID=$(echo $SSM_ACTIVATION | jq -r .ActivationId)
export SSM_ACTIVATION_CODE=$(echo $SSM_ACTIVATION | jq -r .ActivationCode)

aws ssm put-parameter --name "merpochi-activation-id" --value $SSM_ACTIVATION_ID --type "SecureString" --profile $PROFILE
aws ssm put-parameter --name "merpochi-activation-code" --value $SSM_ACTIVATION_CODE --type "SecureString" --profile $PROFILE