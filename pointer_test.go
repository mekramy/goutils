package goutils_test

import (
	"testing"

	"github.com/mekramy/goutils"
)

func TestPointerOf(t *testing.T) {
	value := 42
	ptr := goutils.PointerOf(value)
	if ptr == nil || *ptr != value {
		t.Errorf("PointerOf() = %v, want %v", ptr, value)
	}
}

func TestSafeValue(t *testing.T) {
	var value *int
	if got := goutils.SafeValue(value); got != 0 {
		t.Errorf("SafeValue() = %v, want %v", got, 0)
	}

	val := 42
	value = &val
	if got := goutils.SafeValue(value); got != val {
		t.Errorf("SafeValue() = %v, want %v", got, val)
	}
}

func TestValueOf(t *testing.T) {
	var value *int
	fallback := 42
	if got := goutils.ValueOf(value, fallback); got != fallback {
		t.Errorf("ValueOf() = %v, want %v", got, fallback)
	}

	val := 24
	value = &val
	if got := goutils.ValueOf(value, fallback); got != val {
		t.Errorf("ValueOf() = %v, want %v", got, val)
	}
}

func TestAlter(t *testing.T) {
	var value *int
	fallback := 42
	if got := goutils.Alter(value, fallback); got != fallback {
		t.Errorf("Alter() = %v, want %v", got, fallback)
	}

	val := 0
	value = &val
	if got := goutils.Alter(value, fallback); got != fallback {
		t.Errorf("Alter() = %v, want %v", got, fallback)
	}

	val = 24
	value = &val
	if got := goutils.Alter(value, fallback); got != val {
		t.Errorf("Alter() = %v, want %v", got, val)
	}
}

func TestNullableOf(t *testing.T) {
	var value *int
	if got := goutils.NullableOf(value); got != nil {
		t.Errorf("NullableOf() = %v, want %v", got, nil)
	}

	val := 0
	value = &val
	if got := goutils.NullableOf(value); got != nil {
		t.Errorf("NullableOf() = %v, want %v", got, nil)
	}

	val = 24
	value = &val
	if got := goutils.NullableOf(value); got != value {
		t.Errorf("NullableOf() = %v, want %v", got, value)
	}
}

func TestIsEmpty(t *testing.T) {
	var value *int
	if got := goutils.IsEmpty(value); !got {
		t.Errorf("IsEmpty() = %v, want %v", got, true)
	}

	val := 0
	value = &val
	if got := goutils.IsEmpty(value); !got {
		t.Errorf("IsEmpty() = %v, want %v", got, true)
	}

	val = 24
	value = &val
	if got := goutils.IsEmpty(value); got {
		t.Errorf("IsEmpty() = %v, want %v", got, false)
	}
}

func TestIsSame(t *testing.T) {
	var a, b *int
	if got := goutils.IsSame(a, b); !got {
		t.Errorf("IsSame() = %v, want %v", got, true)
	}

	valA := 42
	a = &valA
	if got := goutils.IsSame(a, b); got {
		t.Errorf("IsSame() = %v, want %v", got, false)
	}

	valB := 42
	b = &valB
	if got := goutils.IsSame(a, b); !got {
		t.Errorf("IsSame() = %v, want %v", got, true)
	}

	valB = 24
	if got := goutils.IsSame(a, b); got {
		t.Errorf("IsSame() = %v, want %v", got, false)
	}
}
