---
layout: default
title: User Manual
nav_order: 2
---

# User Manual
{: .no_toc }

<details open markdown="block">
  <summary>Contents</summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

---

## Prerequisites

### AWS Credentials

`egate` uses the AWS SDK v2 default credential chain. It will look for credentials in this order:

1. Environment variables (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`)
2. AWS CLI profile (`~/.aws/credentials` and `~/.aws/config`)
3. IAM role (EC2 instance profile, ECS task role, etc.)

Run `aws sts get-caller-identity` to verify your credentials are working before using `egate`.

### Required AWS Permissions

Your IAM principal needs permission to:

| Action | Purpose |
|--------|---------|
| `events:ListEventBuses` | `bus list` command |
| `events:PutRule` | Create EventBridge rule |
| `events:PutTargets` | Add CloudWatch target to rule |
| `events:RemoveTargets` | Remove target during teardown |
| `events:DeleteRule` | Delete rule during teardown |
| `events:ListRuleNamesByTarget` | `monitor list` command |
| `logs:CreateLogGroup` | Create CloudWatch log group |
| `logs:PutRetentionPolicy` | Set 1-day log retention |
| `logs:DeleteLogGroup` | Delete log group during teardown |
| `logs:DescribeLogGroups` | Resolve log group ARN for Live Tail |
| `logs:StartLiveTail` | Stream events in real-time |
| `logs:PutResourcePolicy` | Allow EventBridge to write logs |
| `logs:DeleteResourcePolicy` | Remove log resource policy |

---

## Installation

### Homebrew (macOS / Linux)

```bash
brew tap aka-somix/eventbridge-gate
brew install egate
```

### Build from Source

Requires Go 1.23+.

```bash
git clone https://github.com/aka-somix/eventbridge-gate.git
cd eventbridge-gate
make dev-build
# Binary available at ./egate
sudo mv ./egate /usr/local/bin/egate   # optional: add to PATH
```

### Verify Installation

```bash
egate --version
```

---

## Profile Management

`egate profile set` provides an interactive selector for AWS CLI named profiles. Use this when working with multiple AWS accounts or roles.

```bash
egate profile set
```

You'll see a list of profiles from `~/.aws/credentials`:

```
? Select AWS Profile:
  ▸ default
    prod
    staging
    my-sandbox
```

{: .warning }
> Profile selection is **in-memory only** and does not persist across invocations. You must run `egate profile set` each time you open a new terminal session.

---

## Listing Event Buses

List all EventBridge event buses in the current AWS account and region:

```bash
egate bus list
```

Example output:

```
default
my-application-bus
orders-service-bus
```

---

## Monitor Lifecycle

### Attaching a Monitor (`set`)

```bash
egate bus monitor set <bus-name>
```

This creates three AWS resources:

| Resource | Name |
|----------|------|
| CloudWatch Log Group | `/eventbridge-gate/watch/<bus-name>` |
| Log Resource Policy | `allow-logging-from-eventbridge` |
| EventBridge Rule | `eventbridge-gate-rule` (on the target bus) |

The rule uses a wildcard event pattern `{"source": [{"wildcard": "*"}]}` to capture **all events** regardless of source.

**Example:**

```bash
egate bus monitor set my-application-bus
```

{: .note }
> Setup takes a few seconds as the resources are created in AWS. Wait for the confirmation message before running `tail`.

---

### Listing Active Monitors

```bash
egate bus monitor list
```

Shows all buses that currently have an `eventbridge-gate-rule` attached.

---

### Streaming Events (`tail`)

```bash
egate bus monitor tail <bus-name>
```

Opens a [CloudWatch Live Tail](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/live-tail.html) stream on the monitor's log group. Events appear as they flow through the bus in real-time.

**Press `q` to stop streaming.**

**Example output:**

```json
{
  "version": "0",
  "id": "abc123",
  "source": "aws.s3",
  "detail-type": "Object Created",
  "account": "123456789012",
  "region": "eu-west-1",
  "detail": {
    "bucket": { "name": "my-bucket" },
    "object": { "key": "uploads/file.csv" }
  }
}
```

{: .tip }
> Pipe the output through `jq` for formatted JSON: `egate bus monitor tail my-bus | jq .`

---

### Removing a Monitor (`unset`)

```bash
egate bus monitor unset <bus-name>
```

Tears down all three AWS resources created by `set`:

1. Removes the EventBridge rule target
2. Deletes the EventBridge rule
3. Deletes the CloudWatch Log Group (and all logs within it)

**Example:**

```bash
egate bus monitor unset my-application-bus
```

{: .warning }
> All captured logs are permanently deleted when you run `unset`. Export events from CloudWatch before tearing down if you need to keep them.

---

## AWS Resources and Costs

`egate` creates lightweight, short-lived resources designed to minimize cost:

| Resource | Retention | Approximate Cost |
|----------|-----------|-----------------|
| CloudWatch Log Group | 1 day (auto-purged) | < $0.01/month per active monitor |
| EventBridge Rule | Deleted by `unset` | $1.00 per million events matched |

For development and debugging use cases (small event volumes, short sessions), costs are effectively $0.

---

## Troubleshooting

### "AccessDenied" errors

Check that your IAM principal has all permissions listed in the [Prerequisites](#prerequisites) section. Common missing permissions: `logs:PutResourcePolicy` and `events:PutTargets`.

### `tail` shows no events

- Verify the monitor is active: `egate bus monitor list`
- Confirm events are actually being published to the bus (check the source service)
- CloudWatch Live Tail may have a few seconds of latency

### Monitor already exists

If `set` fails because the rule already exists, run `unset` first to clean up, then `set` again.

### Wrong AWS account/region

Set the correct profile with `egate profile set`, or export environment variables:

```bash
export AWS_PROFILE=my-profile
export AWS_REGION=eu-west-1
egate bus list
```
