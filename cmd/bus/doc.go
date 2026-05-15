// Package bus defines the `bus` subcommand group for egate.
// It exposes commands for listing EventBridge event buses and managing
// monitor subcommands (set, unset, list, tail). Business logic is delegated
// to pkg/services.EventBusService.
package bus