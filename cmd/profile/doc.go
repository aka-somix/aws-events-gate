// Package profile defines the `profile` subcommand group for egate.
// It provides an interactive AWS CLI profile selector (via promptui) that
// stores the chosen profile in the in-memory internal/store.ProfileStore for
// use within the current session. The selection does not persist across
// invocations.
package profile