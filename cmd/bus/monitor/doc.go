// Package monitor defines the `bus monitor` subcommand group for egate.
// It exposes commands to create (set), destroy (unset), list, and stream
// events (tail) from CloudWatch-backed EventBridge sniffers. All AWS
// operations are delegated to pkg/services.MonitorService.
package monitor