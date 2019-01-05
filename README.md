# Big Brother

Big Brother is a rideshare price tracking service and is exposed as an HTTP endpoint hosted on AWS Lambda and consumed as a Slack slash command

## Getting Started

1. Local development and deployments are managed by AWS SAM. Install it using `pip install --user --upgrade aws-sam-cli`
2. AWS SAM uses Docker containers locally. Install Docker by visiting their [site](https://docs.docker.com/docker-for-mac/install/)
3. Configure your environment variables in a `local.env.json` file (see `example.env.json` for format)
4. Configure your desired addresses in `addresses.json`
5. Run the application locally with `make offline`

## Deployment

### Deployment Prep

1. Create a target S3 bucket in `us-east-1` for your package deployments. Change the `aws s3` upload script in the Makefile's `deploy` command to your given S3 bucket name
2. Configure your AWS account's `aws_access_key_id` and `aws_secret_access_key` as the `[personal]` profile in `~/.aws/credentials`

### Deploy

1. Add the values of your environment variables to `template.yml`
2. Deploy with `make deploy`

To tear down your production environment, run `make remove`

## Slack Integration

1. Grab your HTTP endpoint by visiting the AWS Lambda or AWS Gateway consoles
2. Create a new Slack app and navigate to the `Slash Commands` section where you'll create a new slash command using your new HTTP endpoint

You'll begin receiving responses in Slack using the `/command` that you defined

## Contributing

Contributions are welcome! Branch off of `master` and submit a Pull Request once complete using the following naming convention:

- `feature/name`
- `task/name`
- `bug/name`
