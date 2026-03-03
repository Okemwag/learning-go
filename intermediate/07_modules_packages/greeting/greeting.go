// Package greeting demonstrates a small exported package inside a module.
package greeting

import "fmt"

// Message is an exported type because it starts with an uppercase letter.
// Exported identifiers are visible to other packages.
type Message struct {
	Name string
}

// Build creates a formatted greeting.
func Build(name string) Message {
	return Message{Name: name}
}

// Text is an exported method.
func (m Message) Text() string {
	return fmt.Sprintf("hello, %s", m.Name)
}

// whisper is unexported because it starts with a lowercase letter.
// Only code inside the greeting package can use it.
func whisper(name string) string {
	return "psst, " + name
}

// DebugText shows internal code can call unexported helpers.
func DebugText(name string) string {
	return whisper(name)
}
