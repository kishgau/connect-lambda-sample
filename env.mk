VERSION=$(shell git describe --abbrev=0 --exact-match --tags)
BRANCH=$(shell git branch | grep \* | cut -d ' ' -f2)
DATE=$(shell date)
COMMIT=$(shell git rev-parse HEAD)
STACK-NAME=connect-lambda-sample
BUILD-PROFILE=sethkor-connect
DEPLOY-PROFILE=sethkor-connect
EXE-NAME=connect-lambda-sample
LAMBDA-NAME=ConnectLambdaSample
REGION=ap-southeast-2
SAM-BUCKET=sethkor-connect-sam
ACCOUNT-ID=166330533654
CONNECT-ARN=arn:aws:connect:ap-southeast-2:166330533654:instance/15a79462-a745-4bcc-9d7a-a94fae538b2e
