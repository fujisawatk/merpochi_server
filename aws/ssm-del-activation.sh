#!/bin/sh

set -e

PROFILE="merpochi-dev"

# SSMマネージドインスタンス削除
SSM_MANAGED_INSTANCE=$(aws ssm describe-instance-information --profile $PROFILE)

export SSM_INSTANCE_ID=$(echo $SSM_MANAGED_INSTANCE | jq -r '.InstanceInformationList[] | .InstanceId')

aws ssm deregister-managed-instance --instance-id $SSM_INSTANCE_ID --profile $PROFILE

# SSMハイブリッドアクティベーション削除
SSM_ACTIVATION=$(aws ssm describe-activations --profile $PROFILE)

export ACTIVATION_ID=$(echo $SSM_ACTIVATION | jq -r '.ActivationList[] | .ActivationId')

aws ssm delete-activation --activation-id $ACTIVATION_ID --profile $PROFILE

# SSMパラメータストアからアクティベーション関係のパラメータ削除
aws ssm delete-parameter --name "merpochi-activation-id" --profile $PROFILE
aws ssm delete-parameter --name "merpochi-activation-code" --profile $PROFILE
