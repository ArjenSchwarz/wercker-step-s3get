# Wercker step for retrieving a single file from S3

This step lets you retrieve a single file from a private S3 bucket. This can be useful in case you store configuration files or assets there that you require in your build.

# AWS Configuration

You will require access keys for an AWS IAM user with sufficient privileges to be set up. The user only needs read permissions for the relevant bucket.

# Parameters

## access_key

The AWS access key from the above mentioned user. Set this up as a global variable in your project.

## secret_key

The AWS secret key from the above mentioned user. Set this up as a global variable in your project.

## bucket

The name of the bucket containing the file you require.

## key

The path of the file in your S3 bucket.

## region

The AWS region where your Lambda function is to be found. Defaults to "us-east-1"

## filename

Optional value for providing the name you wish the file to have. If you don't provide this, the key will be used instead.

# Example

```yml
build:
    steps:
        - arjen/s3get:
            access_key: $AWS_ACCESS_KEY
            secret_key: $AWS_SECRET_KEY
            bucket: my-config-files
            key: path/to/my/config.yml
            filename: $WERCKER_SOURCE_DIR/config.yml
```
