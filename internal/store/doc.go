// Package store provides a singleton in-memory store for the active AWS profile.
// ProfileStore is shared across all commands within a single egate invocation.
// The selected profile is not persisted to disk and is lost when the process exits.
package store