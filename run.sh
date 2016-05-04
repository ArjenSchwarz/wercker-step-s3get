#!/usr/bin/env bash

export AWS_ACCESS_KEY_ID=$WERCKER_S3GET_ACCESS_KEY
export AWS_SECRET_ACCESS_KEY=$WERCKER_S3GET_SECRET_KEY

if [[ -z $WERCKER_S3GET_REGION ]]; then
  WERCKER_S3GET_REGION="us-east-1"
fi

if [[ -z $WERCKER_S3GET_FILENAME ]]; then
  WERCKER_S3GET_FILENAME="$WERCKER_SOURCE_DIR/$WERCKER_S3GET_KEY"
fi

S3GET="${WERCKER_STEP_ROOT}/s3get --bucket ${WERCKER_S3GET_BUCKET} --region ${WERCKER_S3GET_REGION} --key ${WERCKER_S3GET_KEY} --filename ${WERCKER_S3GET_FILENAME}"
debug "$S3GET"
update_output=$($S3GET)

if [[ $? -ne 0 ]];then
    echo "${update_output}"
    fail 'File retrieval failed';
else
    echo "${update_output}"
    success 'File retrieval succeeded';
fi
