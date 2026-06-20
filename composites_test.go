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
	if composites.NewSet(1, 2, 2).Len() != 2 {
		t.Fatal("NewSet")
	}
	rng := composites.NewRange(1, 4, composites.WithStep(1), composites.Exclusive())
	if rng.HasError() || rng.Payload().(composites.Range).Len() != 3 {
		t.Fatal("NewRange/WithStep/Exclusive")
	}
	if composites.NewSymbol("foo") != composites.NewSymbol("foo") {
		t.Fatal("NewSymbol interning")
	}
	if composites.TimeFromUnix(0).ToUnix() != 0 {
		t.Fatal("TimeFromUnix")
	}
	if composites.TimeParse("2006-01-02", "x").HasError() != true {
		t.Fatal("TimeParse error path")
	}
	if composites.DurationFromSeconds(60).ToSeconds() != 60 {
		t.Fatal("DurationFromSeconds")
	}
	if composites.DurationParse("bad").HasError() != true {
		t.Fatal("DurationParse error path")
	}
	if composites.NewPair(1, 2).Second() != 2 {
		t.Fatal("NewPair")
	}
	called := composites.NewProc(func(args ...interface{}) composites.Result {
		return composites.NewResult(composites.WithPayload(args[0]))
	}).Call(99)
	if called.HasError() || called.Payload() != 99 {
		t.Fatal("NewProc/Call")
	}
	if composites.NewEnumerator(1, 2, 3).ToArray().Len() != 3 {
		t.Fatal("NewEnumerator")
	}
	lazy := composites.Generate(1, func(x interface{}) interface{} { return x.(int) + 1 }).Take(2).ToArray()
	if lazy.Len() != 2 {
		t.Fatal("Generate/Take")
	}
}
