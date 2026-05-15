---
layout: default
title: Contributing
nav_order: 3
---

# Contributing to EventBridge Gate
{: .no_toc }

<details open markdown="block">
  <summary>Contents</summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

---

## Development Setup

### Prerequisites

- Go 1.23+
- AWS CLI v2 configured with valid credentials
- [golangci-lint](https://golangci-lint.run/usage/install/) for linting

### Clone and Build

```bash
git clone https://github.com/aka-somix/eventbridge-gate.git
cd eventbridge-gate

# Build the binary (outputs: ./egate)
make dev-build

# Run tests
go test ./...

# Lint
golangci-lint run
```

### GoReleaser (release builds)

```bash
# Requires goreleaser installed
make release
```

---

## Architecture Overview

```
eventbridge-gate/
├── main.go                  # Entry point — calls cmd.Execute()
├── cmd/                     # Cobra CLI command definitions (no AWS logic here)
│   ├── root.go             # Root command; wires bus + profile subcommands
│   ├── bus/
│   │   ├── list.go         # egate bus list
│   │   └── monitor/
│   │       ├── set.go      # egate bus monitor set <bus>
│   │       ├── unset.go    # egate bus monitor unset <bus>
│   │       ├── list.go     # egate bus monitor list
│   │       └── tail.go     # egate bus monitor tail <bus>
│   └── profile/
│       └── set.go          # egate profile set (interactive promptui)
├── pkg/services/            # All AWS API calls and business logic
│   ├── config.go           # MonitorConfig — resource naming constants
│   ├── eventBus.go         # EventBusService (list buses)
│   └── monitor.go          # MonitorService (create/destroy/list/tail)
└── internal/
    ├── store/              # In-memory ProfileStore singleton
    └── aws/                # AWS CLI exec wrapper (currently unused)
```

### Key Design Patterns

- **Cobra commands** are package-level variables, initialized at package init time.
- **Services** (`pkg/services`) hold all AWS SDK v2 client state and are constructed once per command run.
- **ProfileStore** is a singleton shared across all commands in a session (in-memory only).
- **Resource naming** is centralized in `MonitorConfig` (see `pkg/services/config.go`).

---

## Adding a New Command

Follow this pattern for any new subcommand:

### 1. Create the command file

Add a new file in the appropriate `cmd/` subdirectory:

```go
// cmd/bus/describe.go
package bus

import (
    "fmt"
    "github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
    Use:   "describe <busname>",
    Short: "Describe an EventBridge bus",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        busName := args[0]
        // Call service here
        fmt.Println(busName)
    },
}
```

### 2. Register with the parent command

In `cmd/bus/root.go`, add to the `init()` function:

```go
func init() {
    BusCmd.AddCommand(describeCmd)
}
```

### 3. Add business logic to `pkg/services`

If the command needs AWS API calls, add a method to the relevant service in `pkg/services/`. Keep all AWS SDK usage there — commands stay thin.

### 4. Update package docs

Add a one-line summary of the new command to `cmd/bus/doc.go`.

---

## Pull Request Guidelines

- **One feature or fix per PR** — keeps review focused and history clean
- **Run tests and lint before opening a PR:**
  ```bash
  go test ./...
  golangci-lint run
  ```
- **Use conventional commits** (the repo's GoReleaser changelog is auto-generated from them):
  - `feat:` new feature
  - `fix:` bug fix
  - `docs:` documentation only
  - `refactor:` no behavior change
  - `chore:` maintenance (deps, build, CI)
- **No breaking changes without discussion** — open an issue first for anything that changes CLI flags or command behavior

---

## Release Process

Releases are fully automated via [GoReleaser](https://goreleaser.com/) and published to the Homebrew tap `aka-somix/homebrew-eventbridge-gate`.

### Steps

1. Bump the version in `VERSION`
2. Tag the commit:
   ```bash
   git tag -a v0.x.y -m "Release v0.x.y"
   git push origin v0.x.y
   ```
3. GoReleaser runs in CI and publishes:
   - GitHub Release with binaries (Linux, macOS, Windows)
   - Homebrew formula update in the tap repo

### Local release test (dry-run)

```bash
goreleaser release --snapshot --clean
```

---

## Reporting Issues

Open an issue on [GitHub](https://github.com/aka-somix/eventbridge-gate/issues) with:
- `egate` version (`egate --version`)
- AWS region and service involved
- Command run and full error output
