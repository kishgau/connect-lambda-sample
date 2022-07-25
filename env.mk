VERSION=$(shell git describe --abbrev=0 --exact-match --tags)
BRANCH=$(shell git branch | grep \* | cut -d ' ' -f2)
DATE=$(shell date)
COMMIT=$(shell git rev-parse HEAD)
STACK-NAME=connect-lambda-sample
BUILD-PROFILE=$(AWS_PROFILE)
DEPLOY-PROFILE=$(AWS_PROFILE)
EXE-NAME=connect-lambda-sample
LAMBDA-NAME=ConnectLambdaSample
REGION=ap-southeast-2
SAM-BUCKET=$(AWS_PROFILE)-aws-connect-sam
ACCOUNT-ID=272435851616
CONNECT-ARN=arn:aws:connect:ap-southeast-2:272435851616:instance/036262ea-dc94-4585-9959-0258fd8f4b8d
