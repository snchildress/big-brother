.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -o main

clean:
	rm -rf ./bin ./vendor Gopkg.lock

offline: clean build
	sam local start-api

deploy: clean build
	aws s3 mb s3://big-brother-prod \
		--region us-east-1 \
		--profile personal
	sam package \
		--template-file template.yml \
		--output-template-file serverless-output.yml \
		--s3-bucket big-brother-prod \
		--region us-east-1 \
		--profile personal
	sam deploy \
		--template-file \
		serverless-output.yml \
		--stack-name big-brother \
		--capabilities CAPABILITY_IAM \
		--region us-east-1 \
		--profile personal

remove:
	aws cloudformation delete-stack \
		--stack-name big-brother \
		--region us-east-1 \
		--profile personal
