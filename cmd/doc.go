// Package cmd defines the root CLI command for egate using the Cobra framework.
// It wires together all subcommands (bus, profile) and delegates all business
// logic to pkg/services. No AWS interaction happens directly in this package.
package cmd