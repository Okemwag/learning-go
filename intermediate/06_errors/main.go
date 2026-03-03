package main

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
)

// Sentinel errors are package-level reusable error values.
var (
	ErrEmptyName = errors.New("empty name")
	ErrTooShort  = errors.New("name too short")
)

// ValidationError is a custom error type.
// Errors are values, so custom types are often useful when callers need structured details.
type ValidationError struct {
	Field string
	Err   error
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s: %v", v.Field, v.Err)
}

func (v ValidationError) Unwrap() error {
	return v.Err
}

func main() {
	// Use strings for simple examples:
	// fmt.Errorf or errors.New are fine for basic, local errors.
	if err := simpleStringError(""); err != nil {
		fmt.Println("simple string error:", err)
	}

	// Sentinel errors let callers compare against a known meaning.
	if err := validateName("go"); err != nil {
		fmt.Println("sentinel/custom error:", err)
		fmt.Println("is ErrTooShort:", errors.Is(err, ErrTooShort))

		var validationErr ValidationError
		if errors.As(err, &validationErr) {
			fmt.Println("as ValidationError field:", validationErr.Field)
		}
	}

	// Wrapping errors adds context while preserving the original cause.
	if err := loadUser(""); err != nil {
		fmt.Println("wrapped error:", err)
		fmt.Println("wrapped contains ErrEmptyName:", errors.Is(err, ErrEmptyName))
	}

	// Wrapping multiple errors combines several failures.
	if err := validateRequest("", "xy"); err != nil {
		fmt.Println("joined error:", err)
		fmt.Println("joined contains ErrEmptyName:", errors.Is(err, ErrEmptyName))
		fmt.Println("joined contains ErrTooShort:", errors.Is(err, ErrTooShort))
	}

	// Wrapping errors with defer is useful for cleanup failures or close errors.
	if err := processWithDeferredWrap(); err != nil {
		fmt.Println("deferred wrap:", err)
	}

	// panic and recover:
	// panic is for truly exceptional situations, not routine errors.
	if err := runSafely(); err != nil {
		fmt.Println("recovered panic as error:", err)
	}

	// Getting a stack trace from an error:
	// one practical approach is to capture debug.Stack() when creating the error.
	if err := stackTraceExample(); err != nil {
		fmt.Println("stack trace error message:", err)

		var stackErr StackError
		if errors.As(err, &stackErr) {
			fmt.Println("stack trace captured:")
			fmt.Println(string(stackErr.Stack))
		}
	}
}

func simpleStringError(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name is required")
	}
	return nil
}

func validateName(name string) error {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return ValidationError{Field: "name", Err: ErrEmptyName}
	}
	if len(trimmed) < 3 {
		return ValidationError{Field: "name", Err: ErrTooShort}
	}
	return nil
}

func loadUser(name string) error {
	if err := validateName(name); err != nil {
		return fmt.Errorf("load user: %w", err)
	}
	return nil
}

func validateRequest(name string, code string) error {
	var errs []error

	if err := validateName(name); err != nil {
		errs = append(errs, err)
	}
	if len(strings.TrimSpace(code)) < 3 {
		errs = append(errs, fmt.Errorf("code: %w", ErrTooShort))
	}

	if len(errs) == 0 {
		return nil
	}

	return errors.Join(errs...)
}

func processWithDeferredWrap() (err error) {
	// Named return is justified here because defer needs to wrap the returned error.
	defer func() {
		if err != nil {
			err = fmt.Errorf("processWithDeferredWrap: %w", err)
		}
	}()

	err = errors.New("write failed")
	return err
}

func runSafely() (err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			err = fmt.Errorf("panic recovered: %v", recovered)
		}
	}()

	causePanic()
	return nil
}

func causePanic() {
	panic("unexpected nil dependency")
}

type StackError struct {
	Message string
	Stack   []byte
}

func (s StackError) Error() string {
	return s.Message
}

func stackTraceExample() error {
	base := errors.New("database timeout")
	return StackError{
		Message: fmt.Sprintf("query user: %v", base),
		Stack:   debug.Stack(),
	}
}
