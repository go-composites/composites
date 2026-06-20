// Package composites re-exports every go-composites building block under a
// single import, so callers can `import "github.com/go-composites/composites"`
// and reach the whole vocabulary — Array, Boolean, String, Number, Dictionary,
// Result, Error, Null — without importing each repo's /src package individually.
//
// It is a thin facade: types are aliases (so a composites.Result IS a
// result/src.Interface) and constructors are re-exported function values, so
// there is no wrapping cost and full interoperability with the underlying repos.
package composites

import (
	arraysrc "github.com/go-composites/array/src"
	booleansrc "github.com/go-composites/boolean/src"
	dictionarysrc "github.com/go-composites/dictionary/src"
	errorsrc "github.com/go-composites/error/src"
	nullsrc "github.com/go-composites/null/src"
	numbersrc "github.com/go-composites/number/src"
	resultsrc "github.com/go-composites/result/src"
	stringsrc "github.com/go-composites/string/src"
)

// Composite interface types.
type (
	Array      = arraysrc.Interface
	Boolean    = booleansrc.Interface
	Dictionary = dictionarysrc.Interface
	Error      = errorsrc.Interface
	Null       = nullsrc.Interface
	Number     = numbersrc.Interface
	Result     = resultsrc.Interface
	String     = stringsrc.Interface
)

// Functional-option types for the option-taking constructors.
type (
	DictionaryOption = dictionarysrc.Option
	NumberOption     = numbersrc.Option
	ResultOption     = resultsrc.Option
	StringOption     = stringsrc.Option
)

// Constructors and options, re-exported as function values.
var (
	NewArray      = arraysrc.New
	NewBoolean    = booleansrc.New
	True          = booleansrc.True
	False         = booleansrc.False
	NewString     = stringsrc.New
	NewError      = errorsrc.New
	NewNull       = nullsrc.New
	NewNumber     = numbersrc.New
	NewResult     = resultsrc.New
	NewDictionary = dictionarysrc.New

	WithGoString = stringsrc.WithGoString
	WithPayload  = resultsrc.WithPayload
	WithError    = resultsrc.WithError
	WithInt      = numbersrc.WithInt
	WithFloat    = numbersrc.WithFloat
	WithPairs    = dictionarysrc.WithPairs
)
