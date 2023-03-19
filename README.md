# maya

Maya is web systems control bot that can manage AWS EC2 instance lifecycle through Discord chat.

Presently, it can list, get status, start, and stop EC2 instances.

## Setup
### Infrastructure Requirements
Maya runs as a Go binary on an EC2 instance with the following requirements:

1. Go 1.20+ is installed.
2. An IAM instance role with `ec2:StartInstances`, `ec2:StopInstances`, and `ec2:DescribeInstance*`.

### Install
```
go build
./maya -token $DISCORD_TOKEN
```

## Usage

- `/describe-instances` - Get a list of all instances manageable by Maya.
- `/describe-instance-status [instance-id]` Get an instance's reachability and system health status.
- `/start-instance [instance-id]` - Start a specified instance.
- `/stop-instance [instance-id]` - Stop a specified instance.
