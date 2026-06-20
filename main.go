// Command (demo): one import reaches the whole go-composites vocabulary.
package main

import (
	"fmt"

	"github.com/go-composites/composites"
)

func main() {
	div := composites.NewNumber(composites.WithInt(10)).
		Div(composites.NewNumber(composites.WithInt(0)))
	fmt.Printf("10 / 0 → HasError=%t (%v)\n", div.HasError(), div.Error().Message())

	d := composites.NewDictionary().Set("k", "v")
	fmt.Printf("dict.Get(k) → %v\n", d.Get("k").Payload())

	a := composites.NewArray()
	a.Push(composites.True())
	fmt.Printf("array.First() → %v\n", a.First().Payload())

	r := composites.NewResult(composites.WithPayload("ok"))
	fmt.Printf("result → %v (err=%t)\n", r.Payload(), r.HasError())
}
