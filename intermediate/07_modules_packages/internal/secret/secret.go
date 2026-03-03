// Package secret lives under internal, so it can only be imported by code
// within the parent tree: intermediate/07_modules_packages and its children.
package secret

// Hint returns a simple internal-only message.
func Hint() string {
	return "internal packages hide implementation details from outside callers"
}
