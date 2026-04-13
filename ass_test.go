package ass_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/Thiht/ass"
)

type mockTB struct {
	testing.TB
	failed  bool
	message string
}

func (m *mockTB) Helper() {}
func (m *mockTB) Fatalf(format string, args ...any) {
	m.failed = true
	m.message = fmt.Sprintf(format, args...)
}

var errSentinel = errors.New("sentinel error")

type customErr struct{ msg string }

func (e *customErr) Error() string { return e.msg }

func TestErr(t *testing.T) {
	t.Run("got nil, expect nil: passes", func(t *testing.T) {
		ass.Err(t, nil, nil)
	})

	t.Run("got nil, expect error: fails", func(t *testing.T) {
		mb := &mockTB{}
		ass.Err(mb, errors.New("oops"), nil)
		if !mb.failed {
			t.Fatal("expected failure, got none")
		}
	})

	t.Run("got error, expect error (bool): passes", func(t *testing.T) {
		ass.Err(t, errors.New("oops"), true)
	})

	t.Run("got nil, expect error (bool): fails", func(t *testing.T) {
		mb := &mockTB{}
		ass.Err(mb, nil, true)
		if !mb.failed {
			t.Fatal("expected failure, got none")
		}
	})

	t.Run("got nil, expect no error (bool): passes", func(t *testing.T) {
		ass.Err(t, nil, false)
	})

	t.Run("got error, expect no error (bool): fails", func(t *testing.T) {
		mb := &mockTB{}
		ass.Err(mb, errors.New("oops"), false)
		if !mb.failed {
			t.Fatal("expected failure, got none")
		}
	})

	t.Run("got error, expect substring: passes", func(t *testing.T) {
		ass.Err(t, errors.New("something went wrong"), "went wrong")
	})

	t.Run("got error, expect substring: fails", func(t *testing.T) {
		mb := &mockTB{}
		ass.Err(mb, errors.New("something went wrong"), "network")
		if !mb.failed {
			t.Fatal("expected failure, got none")
		}
	})

	t.Run("got nil, expect empty string: passes", func(t *testing.T) {
		ass.Err(t, nil, "")
	})

	t.Run("got error, expect empty string: fails", func(t *testing.T) {
		mb := &mockTB{}
		ass.Err(mb, errors.New("oops"), "")
		if !mb.failed {
			t.Fatal("expected failure, got none")
		}
	})

	t.Run("got error, expect sentinel error (errors.Is): passes", func(t *testing.T) {
		wrapped := fmt.Errorf("wrapped: %w", errSentinel)
		ass.Err(t, wrapped, errSentinel)
	})

	t.Run("got error, expect sentinel error (errors.Is): fails", func(t *testing.T) {
		mb := &mockTB{}
		ass.Err(mb, errors.New("other"), errSentinel)
		if !mb.failed {
			t.Fatal("expected failure, got none")
		}
	})

	t.Run("got error, expect custom error type (errors.As): passes", func(t *testing.T) {
		ass.Err(t, &customErr{"boom"}, reflect.TypeOf(&customErr{}))
	})

	t.Run("got error, expect custom error type (errors.As): fails", func(t *testing.T) {
		mb := &mockTB{}
		ass.Err(mb, errors.New("plain"), reflect.TypeOf(&customErr{}))
		if !mb.failed {
			t.Fatal("expected failure, got none")
		}
	})

	t.Run("unsupported type: fails", func(t *testing.T) {
		mb := &mockTB{}
		ass.Err(mb, errors.New("oops"), 1)
		if !mb.failed {
			t.Fatal("expected failure, got none")
		}
	})
}
