package utils

import (
	"reflect"
	"testing"
)

func AssertEqual[T comparable](t *testing.T, value T, expected T) {
	t.Helper()

	if value != expected {
		t.Fatalf("Not equal:\nexpected: %#v\nactual  : %#v", expected, value)
	}
}

func AssertNotEqual[T comparable](t *testing.T, value T, expected T) {
	t.Helper()

	if value == expected {
		t.Fatalf("Expected to not equal %#v", expected)
	}
}

func AssertTrue(t *testing.T, value bool) {
	t.Helper()

	if !value {
		t.Fatal("should be true")
	}
}

func AssertFalse(t *testing.T, value bool) {
	t.Helper()

	if !value {
		t.Fatal("should be false")
	}
}

func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Slice, reflect.UnsafePointer:
		return value.IsNil()
	}

	return false
}

func AssertNil(t *testing.T, value interface{}) {
	t.Helper()

	if !isNil(value) {
		t.Fatalf("Expected nil, but got %#v", value)
	}
}

func AssertNotNil(t *testing.T, value interface{}) {
	t.Helper()

	if isNil(value) {
		t.Fatal("Expected value to not be nil.")
	}
}
