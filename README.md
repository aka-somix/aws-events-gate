# EventBridge Gate

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/aka-somix/eventbridge-gate)](https://github.com/aka-somix/eventbridge-gate/releases)

> Real-time AWS EventBridge event debugger for developers.

`egate` hooks a temporary sniffer onto any EventBridge bus and streams every event to your terminal in real-time. When you're done, it tears everything down cleanly — no orphaned AWS resources.

---

## Install

```bash
brew tap aka-somix/eventbridge-gate
brew install egate
```

Or build from source (requires Go 1.23+):

```bash
git clone https://github.com/aka-somix/eventbridge-gate.git
cd eventbridge-gate && make dev-build
```

---

## Quick Start

```bash
# Optional: pick an AWS profile for this session
egate profile set

# Attach a monitor to your bus
egate bus monitor set my-event-bus

# Stream events live — press q to stop
egate bus monitor tail my-event-bus

# Clean up all AWS resources
egate bus monitor unset my-event-bus
```

---

## Commands

| Command | Description |
|---|---|
| `egate bus list` | List all EventBridge buses |
| `egate bus monitor set <bus>` | Attach a sniffer to a bus |
| `egate bus monitor list` | Show buses with active monitors |
| `egate bus monitor tail <bus>` | Stream events live (press `q` to stop) |
| `egate bus monitor unset <bus>` | Remove the sniffer and its AWS resources |
| `egate profile set` | Interactively select an AWS CLI profile |

---

## Documentation

Full docs are available on **[GitHub Pages](https://aka-somix.github.io/eventbridge-gate)**:

- [User Manual](https://aka-somix.github.io/eventbridge-gate/manual) — prerequisites, permissions, all commands, troubleshooting
- [Contributing Guide](https://aka-somix.github.io/eventbridge-gate/contributing) — dev setup, architecture, PR guidelines, release process

---

## Contributing

Bug reports, feature requests, and pull requests are welcome. See the [Contributing Guide](https://aka-somix.github.io/eventbridge-gate/contributing) for how to get started.

---

## License

Apache 2.0 — see [LICENSE](LICENSE).
