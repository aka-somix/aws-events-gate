// Package services contains all business logic and AWS API interactions for egate.
//
// It exposes two primary services:
//   - EventBusService: lists available EventBridge event buses.
//   - MonitorService: manages the full monitor lifecycle — creating a CloudWatch
//     Log Group + EventBridge rule that captures all events on a bus, listing
//     active monitors, streaming live events via CloudWatch Live Tail, and
//     destroying all associated AWS resources when done.
//
// All AWS calls use the SDK v2 (github.com/aws/aws-sdk-go-v2). Resource naming
// constants (log group prefix, rule name, policy name, etc.) are defined in
// config.go via MonitorConfig.
package services