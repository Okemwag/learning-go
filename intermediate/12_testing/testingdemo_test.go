package testingdemo

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

// TestMain is one way to do package-level setup and teardown.
// Prefer test-local setup when possible, but TestMain is useful when package-wide
// lifecycle control is genuinely needed.
func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestAdd(t *testing.T) {
	// Reporting test failures:
	// t.Fatalf stops this test immediately and prints the failure message.
	got := Add(2, 3)
	if got != 5 {
		t.Fatalf("Add(2, 3) = %d, want 5", got)
	}
}

func TestNormalizeNameTable(t *testing.T) {
	// Running table tests is idiomatic in Go.
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{name: "trim and uppercase", input: "  gopher ", want: "GOPHER"},
		{name: "empty input", input: "   ", wantErr: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NormalizeName(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("NormalizeName(%q) error = nil, want non-nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("NormalizeName(%q) unexpected error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Fatalf("NormalizeName(%q) = %q, want %q", tt.input, got, tt.want)
			}

			// If this project adds github.com/google/go-cmp/cmp later,
			// cmp.Diff(want, got) is often clearer for complex values.
		})
	}
}

func TestBuildGreetingWithEnv(t *testing.T) {
	// Testing with environment variables:
	// t.Setenv automatically restores the old value after the test.
	t.Setenv("GREETING_PREFIX", "Welcome")

	got := BuildGreeting("Gopher")
	want := "Welcome, Gopher"
	if got != want {
		t.Fatalf("BuildGreeting() = %q, want %q", got, want)
	}
}

func TestReadSampleData(t *testing.T) {
	// Storing sample test data:
	// testdata/ is a special convention for files used by tests.
	path := filepath.Join("testdata", "sample.txt")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile(%q) error: %v", path, err)
	}

	got := strings.TrimSpace(string(data))
	if got != "gopher academy" {
		t.Fatalf("sample data = %q, want %q", got, "gopher academy")
	}
}

type stubNotifier struct {
	mu       sync.Mutex
	messages []string
}

func (s *stubNotifier) Notify(message string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messages = append(s.messages, message)
	return nil
}

func (s *stubNotifier) Messages() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]string, len(s.messages))
	copy(out, s.messages)
	return out
}

func TestWelcomeUserWithStub(t *testing.T) {
	// Using stubs in Go:
	// small interfaces make it easy to provide lightweight fakes in tests.
	t.Setenv("GREETING_PREFIX", "Hello")

	stub := &stubNotifier{}
	if err := WelcomeUser("Gopher", stub); err != nil {
		t.Fatalf("WelcomeUser() error: %v", err)
	}

	got := stub.Messages()
	if len(got) != 1 {
		t.Fatalf("stub message count = %d, want 1", len(got))
	}
	if got[0] != "Hello, Gopher" {
		t.Fatalf("stub first message = %q, want %q", got[0], "Hello, Gopher")
	}
}

func TestCounterConcurrent(t *testing.T) {
	// Running tests concurrently:
	// Use synchronization so the test remains deterministic.
	var (
		counter Counter
		wg      sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()

	if got := counter.Value(); got != 10 {
		t.Fatalf("counter.Value() = %d, want 10", got)
	}
}

func FuzzNormalizeName(f *testing.F) {
	// Fuzzing explores many generated inputs and keeps crashing inputs as seeds.
	f.Add("gopher")
	f.Add("  go  ")

	f.Fuzz(func(t *testing.T, input string) {
		got, err := NormalizeName(input)
		if err == nil && got == "" {
			t.Fatalf("NormalizeName(%q) returned empty string with nil error", input)
		}
	})
}

func BenchmarkAdd(b *testing.B) {
	// Benchmarks measure performance of a focused code path.
	for i := 0; i < b.N; i++ {
		_ = Add(10, 20)
	}
}

func BenchmarkNormalizeName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NormalizeName("  gopher  ")
	}
}
