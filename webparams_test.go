package webparams

import (
	"testing"
)

func TestExtractInt64(t *testing.T) {

	intval, err := ExtractInt64("12345")

	if err != nil {
		t.Fatalf("Error ExtractInt64 did not convert to int64 12345 : %s", err)
	}

	if intval != 12345 {
		t.Fatalf("Error ExtractInt64 did not return 12345 it returned : %v", intval)
	}

	intval, err = ExtractInt64("")

	if err == nil {
		t.Fatalf("Error ExtractInt64 should return error on empty string but it did not return an error")
	}

	intval, err = ExtractInt64("3.14159260")

	if err == nil {
		t.Fatalf("Error ExtractInt64 should return error on string looking like a floating point number but it did not return an error")
	}

	intval, err = ExtractInt64("abc123")

	if err == nil {
		t.Fatalf("Error ExtractInt64 should return error on string with mixed alpha numeric values but it did not return an error")
	}
}

func TestExtractFloat64(t *testing.T) {

	floatval, err := ExtractFloat64("-0.00")

	if err != nil {
		t.Fatalf("Error ExtractFloat64 did not convert to float64 -0.00 : %s", err)
	}

	if floatval != -0.00 {
		t.Fatalf("Error ExtractFloat64 did not return -0.00 it returned : %v", floatval)
	}

	floatval, err = ExtractFloat64("")

	if err == nil {
		t.Fatalf("Error ExtractFloat64 should return error on empty string but it did not return an error")
	}

	floatval, err = ExtractFloat64("3.14159260")

	if err != nil {
		t.Fatalf("Error ExtractFloat64 should not return error for pi, it returned: %s", err)
	}

	floatval, err = ExtractFloat64("abc123")

	if err == nil {
		t.Fatalf("Error ExtractFloat64 should return error on string with mixed alpha numeric values but it did not return an error")
	}
}
