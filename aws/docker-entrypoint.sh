#!/bin/sh

set -e

amazon-ssm-agent -register -code $SSM_ACTIVATE_CODE -id $SSM_ACTIVATE_ID -region "ap-northeast-1"
amazon-ssm-agent &

/go/bin/server