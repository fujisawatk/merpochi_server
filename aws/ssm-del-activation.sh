#!/bin/sh

set -e

PROFILE="merpochi-dev"
SSM_ACTIVATION=$(aws ssm describe-instance-information --profile $PROFILE)

export SSM_INSTANCE_ID=$(echo $SSM_ACTIVATION | jq -r '.InstanceInformationList[] | .InstanceId')

aws ssm deregister-managed-instance --instance-id $SSM_INSTANCE_ID --profile $PROFILE

aws ssm delete-parameter --name "merpochi-activation-id" --profile $PROFILE
aws ssm delete-parameter --name "merpochi-activation-code" --profile $PROFILE
