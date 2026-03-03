package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt Timestamp `json:"created_at"`
}

// Timestamp shows custom JSON parsing and formatting.
type Timestamp struct {
	time.Time
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Format(time.RFC3339))
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	parsed, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		return err
	}

	t.Time = parsed
	return nil
}

func main() {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	demoIO(logger)
	demoTime(logger)
	demoJSON(logger)
	demoHTTP(logger)
}

func demoIO(logger *slog.Logger) {
	// io and friends:
	// Readers and Writers are foundational interfaces in the standard library.
	reader := strings.NewReader("hello through io")
	var dst bytes.Buffer

	written, err := io.Copy(&dst, reader)
	if err != nil {
		fmt.Println("io.Copy error:", err)
		return
	}

	fmt.Println("io.Copy bytes:", written)
	fmt.Println("copied text:", dst.String())
	logger.Info("io demo complete", "bytes", written)
}

func demoTime(logger *slog.Logger) {
	// time and monotonic time:
	// time.Now() includes a monotonic component used for safe duration math.
	start := time.Now()
	time.Sleep(5 * time.Millisecond)
	elapsed := time.Since(start)
	fmt.Println("elapsed with monotonic clock:", elapsed > 0)

	// Timers and timeouts:
	timer := time.NewTimer(10 * time.Millisecond)
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Println("timer fired")
	case <-time.After(20 * time.Millisecond):
		fmt.Println("outer timeout fired first")
	}

	// Context-based timeouts are often cleaner for real APIs.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	<-ctx.Done()
	fmt.Println("context timeout:", ctx.Err())
	logger.Info("time demo complete")
}

func demoJSON(logger *slog.Logger) {
	// Struct tags add metadata that encoding/json uses.
	original := User{
		ID:        1,
		Name:      "Gopher",
		CreatedAt: Timestamp{Time: time.Date(2026, 3, 3, 12, 0, 0, 0, time.UTC)},
	}

	// Marshaling turns Go values into JSON.
	data, err := json.Marshal(original)
	if err != nil {
		fmt.Println("marshal error:", err)
		return
	}
	fmt.Println("marshaled JSON:", string(data))

	// Unmarshaling turns JSON into Go values.
	var parsed User
	if err := json.Unmarshal(data, &parsed); err != nil {
		fmt.Println("unmarshal error:", err)
		return
	}
	fmt.Println("unmarshaled user:", parsed.Name, parsed.CreatedAt.Format(time.RFC3339))

	// JSON, readers, and writers:
	// Encoder and Decoder work directly on streams.
	stream := strings.NewReader(`{"id":2,"name":"Stream","created_at":"2026-03-03T13:00:00Z"}`)
	decoder := json.NewDecoder(stream)

	var streamed User
	if err := decoder.Decode(&streamed); err != nil {
		fmt.Println("decode stream error:", err)
		return
	}
	fmt.Println("decoded from reader:", streamed.Name)

	var out bytes.Buffer
	encoder := json.NewEncoder(&out)
	if err := encoder.Encode(streamed); err != nil {
		fmt.Println("encode stream error:", err)
		return
	}
	fmt.Println("encoded to writer:", strings.TrimSpace(out.String()))
	logger.Info("json demo complete", "user_id", parsed.ID)
}

func demoHTTP(logger *slog.Logger) {
	// The server:
	// use a local handler and recorder so the example works without opening sockets.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ResponseController gives lower-level response controls.
		controller := http.NewResponseController(w)
		_ = controller.SetWriteDeadline(time.Now().Add(100 * time.Millisecond))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if _, err := io.WriteString(w, `{"ok":true}`); err == nil {
			_ = controller.Flush()
		}
	})

	// The client:
	client := http.Client{
		Timeout: 200 * time.Millisecond,
		Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, req)
			return recorder.Result(), nil
		}),
	}

	req, err := http.NewRequest(http.MethodGet, "http://example.local", nil)
	if err != nil {
		fmt.Println("http request build error:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http client error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error:", err)
		return
	}

	fmt.Println("http status:", resp.Status)
	fmt.Println("http body:", string(body))
	logger.Info("http demo complete", "status", resp.StatusCode)
}
