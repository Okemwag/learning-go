// Package testingdemo provides small functions used to demonstrate Go testing.
package testingdemo

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

// Add is a tiny public API function that is easy to test.
func Add(a int, b int) int {
	return a + b
}

// NormalizeName trims space and uppercases a name.
func NormalizeName(input string) (string, error) {
	name := strings.TrimSpace(input)
	if name == "" {
		return "", fmt.Errorf("name cannot be empty")
	}
	return strings.ToUpper(name), nil
}

// GreetingPrefix reads a prefix from the environment.
func GreetingPrefix() string {
	prefix := os.Getenv("GREETING_PREFIX")
	if prefix == "" {
		return "Hello"
	}
	return prefix
}

// BuildGreeting combines environment-derived and explicit input.
func BuildGreeting(name string) string {
	return GreetingPrefix() + ", " + name
}

// Notifier is a tiny interface used to demonstrate stubs in tests.
type Notifier interface {
	Notify(message string) error
}

// WelcomeUser sends a greeting through a dependency.
func WelcomeUser(name string, notifier Notifier) error {
	return notifier.Notify(BuildGreeting(name))
}

// FetchStatusCode is used with httptest in unit and integration examples.
func FetchStatusCode(client *http.Client, url string) (int, error) {
	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

// Counter is safe for concurrent use because it protects state with a mutex.
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc increments the counter safely.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current count safely.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
