// Package aws provides a thin wrapper around the AWS CLI executable.
// AwsCommand constructs an exec.Cmd for a given AWS CLI sub-command, automatically
// appending the --profile flag when a profile is active in internal/store.
// Note: this package is currently unused — all AWS interactions in egate go
// through the AWS SDK v2 directly via pkg/services.
package aws