---
title: Home
nav_order: 1
---

# EventBridge Gate

**Real-time AWS EventBridge event debugger for developers.**

Debugging event-driven systems on AWS is hard. When an EventBridge rule isn't firing, or you're unsure how an event is structured, you need to *see* what's actually flowing through your bus — right now, not after digging through CloudWatch queries.

`egate` hooks a temporary sniffer onto any EventBridge bus in your account and streams every event to your terminal in real-time. When you're done, it tears everything down cleanly.

---

## Key Features

- **Real-time event streaming** via CloudWatch Live Tail — see events as they happen
- **Zero-configuration sniffers** — one command creates the CloudWatch log group, resource policy, and EventBridge rule
- **Clean teardown** — `unset` destroys all AWS resources, leaving no orphans
- **Multi-bus support** — monitor multiple buses independently
- **AWS profile support** — interactive profile selector for multi-account workflows
- **Minimal cost** — 1-day log retention keeps CloudWatch costs near zero

---

## Install

### Homebrew (macOS / Linux)

```bash
brew tap aka-somix/eventbridge-gate
brew install egate
```

### Build from source

```bash
git clone https://github.com/aka-somix/eventbridge-gate.git
cd eventbridge-gate
make dev-build   # outputs ./egate
```

---

## Get Started in 60 Seconds

```bash
# 1. (Optional) Select an AWS profile for this session
egate profile set

# 2. Attach a monitor to your bus
egate bus monitor set my-event-bus

# 3. Stream events in real-time — press q to stop
egate bus monitor tail my-event-bus

# 4. Clean up when done
egate bus monitor unset my-event-bus
```

---

## Commands at a Glance

| Command | Description |
|---|---|
| `egate bus list` | List all EventBridge buses in the account |
| `egate bus monitor set <bus>` | Attach a sniffer to a bus |
| `egate bus monitor list` | Show buses with active monitors |
| `egate bus monitor tail <bus>` | Stream events live (press `q` to stop) |
| `egate bus monitor unset <bus>` | Remove the sniffer and all its AWS resources |
| `egate profile set` | Interactively select an AWS CLI profile |

---

{: .note }
> `egate` creates real AWS resources (CloudWatch Log Group + EventBridge Rule) that may incur minimal costs. See the [User Manual](manual) for details on AWS permissions and cost expectations.

[Full User Manual](manual){: .btn .btn-primary .mr-2 }
[Contributing](contributing){: .btn }
