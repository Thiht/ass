package ass

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

// Err asserts that `got` matches the expectation defined by `want`.
// The behavior depends on the type of `want`:
//
//   - `nil` or `""`: asserts `got` is `nil`
//   - `true`: asserts `got` is non-`nil`
//   - `false`: asserts `got` is `nil`
//   - `string`: asserts `got.Error()` contains the substring
//   - `error`: asserts `errors.Is(got, want)`
//   - `reflect.Type`: asserts `errors.As(got, want)` (use `[reflect.TypeOf]` to create the type)
func Err(tb testing.TB, got error, want any) {
	tb.Helper()

	switch w := want.(type) {
	case nil:
		if got != nil {
			tb.Fatalf("unexpected error: %v", got)
		}

	case bool:
		if w && got == nil {
			tb.Fatalf("expected error, got nil")
		}
		if !w && got != nil {
			tb.Fatalf("unexpected error: %v", got)
		}

	case string:
		if w == "" {
			if got != nil {
				tb.Fatalf("unexpected error: %v", got)
			}
			return
		}

		if !strings.Contains(got.Error(), w) {
			tb.Fatalf("expected error %q to contain %q", got.Error(), w)
		}

	case reflect.Type:
		target := reflect.New(w).Interface()
		if !errors.As(got, target) {
			tb.Fatalf("expected error %T to be %s", got, w)
		}

	case error:
		if !errors.Is(got, w) {
			tb.Fatalf("expected error %T(%v) to be %T(%v)", got, got, w, w)
		}

	default:
		tb.Fatalf("unsupported error assertion: %T", want)
	}
}
