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
