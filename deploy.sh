#!/usr/bin/env bash

echo "Compiling handlers..."

GOOS=linux go build -o main
zip deployment.zip main

echo "Built the go file."
echo "Starting deploy via CloudFormation"

aws cloudformation package --template-file template.yaml --output-template-file serverless-output.yaml \
  --s3-bucket <<<YOUR_BUCKET_NAME>>>
aws cloudformation deploy --template-file serverless-output.yaml --stack-name HelloGoWebApp --capabilities CAPABILITY_IAM

echo "Done."
