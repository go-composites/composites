package composites_test

import (
	"testing"

	"github.com/go-composites/composites"
)

// TestReExports verifies every re-exported constructor wires through to the
// underlying composite and never returns nil (the Null-Object invariant).
func TestReExports(t *testing.T) {
	if composites.NewArray() == nil {
		t.Fatal("NewArray")
	}
	if composites.NewBoolean(true) == nil || composites.True() == nil || composites.False() == nil {
		t.Fatal("Boolean constructors")
	}
	if composites.NewString(composites.WithGoString("x")).ToGoString() != "x" {
		t.Fatal("NewString/WithGoString")
	}
	if composites.NewError("e").Message() != "e" {
		t.Fatal("NewError")
	}
	if !composites.NewNull().IsNull() {
		t.Fatal("NewNull")
	}
	if composites.NewNumber(composites.WithInt(7)).ToGoInt() != 7 {
		t.Fatal("NewNumber/WithInt")
	}
	if composites.NewNumber(composites.WithFloat(1.5)).ToGoFloat() != 1.5 {
		t.Fatal("NewNumber/WithFloat")
	}
	r := composites.NewResult(composites.WithPayload(42))
	if r.HasError() || r.Payload() != 42 {
		t.Fatal("NewResult/WithPayload")
	}
	if !composites.NewResult(composites.WithError(composites.NewError("boom"))).HasError() {
		t.Fatal("NewResult/WithError")
	}
	d := composites.NewDictionary(composites.WithPairs(map[interface{}]interface{}{"a": 1}))
	if d.Get("a").Payload() != 1 {
		t.Fatal("NewDictionary/WithPairs")
	}
}
