# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Build binary (outputs: ./egate)
make dev-build

# Release build (goreleaser)
make release

# Run tests
go test ./...

# Single package test
go test ./pkg/services/...

# Lint
golangci-lint run
```

## Architecture

**EventBridge Gate** (`egate`) — CLI debugger for AWS EventBridge. Creates temporary CloudWatch-backed sniffers on event buses to capture and stream events in real-time.

### Package layout

```
cmd/           Cobra CLI commands
  bus/         Bus subcommands (list, monitor set/unset/list/tail)
  profile/     AWS profile selector
pkg/services/  Business logic — all AWS operations
internal/
  store/       In-memory profile state (singleton)
  aws/         AWS execution helpers
```

### CLI tree

```
egate
├── bus list                    List EventBridge buses
├── bus monitor set <bus>       Create monitor (CloudWatch log group + EB rule)
├── bus monitor unset <bus>     Destroy monitor resources
├── bus monitor list            List buses with active monitors
├── bus monitor tail <bus>      Stream live events from monitor
└── profile set                 Interactive AWS profile selector (promptui)
```

### How a monitor works

`bus monitor set` creates three AWS resources:
1. CloudWatch Log Group: `/eventbridge-gate/watch/<busname>`
2. Log Resource Policy: `allow-logging-from-eventbridge`
3. EventBridge Rule with wildcard pattern → target: CloudWatch Log Group

`bus monitor tail` opens a CloudWatch Live Tail stream on that log group, printing events to stdout until user presses `q`.

`bus monitor unset` tears down all three resources.

Resource name constants live in `pkg/services/config.go`.

### Key service: `MonitorService` (`pkg/services/monitor.go`)

Methods: `Create`, `Destroy`, `List`, `Tail` — all AWS operations use `context.TODO()`.

### AWS profile

`profile set` stores the selected profile in `internal/store/store.go` (singleton, in-memory only — does not persist across invocations).

## Binary

Binary name: `egate`. Released via GoReleaser to Homebrew tap `aka-somix/homebrew-eventbridge-gate`.
