<p align="center"><img src="https://raw.githubusercontent.com/go-composites/brand/main/social/go-composites.png" alt="go-composites/composites" width="720"></p>

# composites

[![ci](https://github.com/go-composites/composites/actions/workflows/ci.yml/badge.svg)](https://github.com/go-composites/composites/actions/workflows/ci.yml)

One import for the whole [go-composites](https://github.com/go-composites)
vocabulary. Instead of importing each `…/src` package, import `composites` and
reach every building block — `Array`, `Boolean`, `String`, `Number`,
`Dictionary`, `Result`, `Error`, `Null` — with their constructors and options.

It is a zero-cost facade: the types are **aliases** (a `composites.Result` *is* a
`result/src.Interface`, fully interoperable), and the constructors are
re-exported function values.

## Install

```sh
go get github.com/go-composites/composites
```

## Use

```go
import "github.com/go-composites/composites"

n := composites.NewNumber(composites.WithInt(10))
r := n.Div(composites.NewNumber(composites.WithInt(0))) // Result with an error, not a panic

d := composites.NewDictionary().Set("k", "v")
v := d.Get("k") // Result; miss → Result carrying an error

a := composites.NewArray()
a.Push(composites.True())
```

| constructor | composite |
| --- | --- |
| `NewArray()` | ordered collection (`Map`/`Filter`/`Reduce`/…) |
| `NewBoolean(b)` / `True()` / `False()` | boolean |
| `NewString(WithGoString(s))` | string |
| `NewNumber(WithInt(i))` / `WithFloat(f)` | number (Result-returning arithmetic) |
| `NewDictionary(WithPairs(m))` | key→value map |
| `NewResult(WithPayload(x))` / `WithError(e)` | result |
| `NewError(msg)` | error |
| `NewNull()` | the Null object |

## License

BSD-3-Clause © the go-composites/composites authors.
