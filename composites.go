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
	bignumbersrc "github.com/go-composites/bignumber/src"
	booleansrc "github.com/go-composites/boolean/src"
	complexsrc "github.com/go-composites/complex/src"
	dictionarysrc "github.com/go-composites/dictionary/src"
	enumeratorsrc "github.com/go-composites/enumerator/src"
	errorsrc "github.com/go-composites/error/src"
	nullsrc "github.com/go-composites/null/src"
	numbersrc "github.com/go-composites/number/src"
	orderedsetsrc "github.com/go-composites/orderedset/src"
	pairsrc "github.com/go-composites/pair/src"
	procsrc "github.com/go-composites/proc/src"
	rangesrc "github.com/go-composites/range/src"
	rationalsrc "github.com/go-composites/rational/src"
	resultsrc "github.com/go-composites/result/src"
	setsrc "github.com/go-composites/set/src"
	stringsrc "github.com/go-composites/string/src"
	symbolsrc "github.com/go-composites/symbol/src"
	timesrc "github.com/go-composites/time/src"
	durationsrc "github.com/go-composites/time/src/duration"
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
	Set        = setsrc.Interface
	Range      = rangesrc.Interface
	Symbol     = symbolsrc.Interface
	Time       = timesrc.Interface
	Duration   = durationsrc.Interface
	Pair       = pairsrc.Interface
	Proc       = procsrc.Interface
	Enumerator = enumeratorsrc.Interface
	BigNumber  = bignumbersrc.Interface
	OrderedSet = orderedsetsrc.Interface
	Rational   = rationalsrc.Interface
	Complex    = complexsrc.Interface
)

// Functional-option types for the option-taking constructors.
type (
	DictionaryOption = dictionarysrc.Option
	NumberOption     = numbersrc.Option
	ResultOption     = resultsrc.Option
	StringOption     = stringsrc.Option
	RangeOption      = rangesrc.Option
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

	NewSet              = setsrc.New
	NewRange            = rangesrc.New
	NewSymbol           = symbolsrc.New
	NewPair             = pairsrc.New
	NewProc             = procsrc.New
	NewEnumerator       = enumeratorsrc.New
	Generate            = enumeratorsrc.Generate
	NewOrderedSet       = orderedsetsrc.New
	BigNumberFromInt64  = bignumbersrc.FromInt64
	BigNumberParse      = bignumbersrc.FromString
	RationalFromInts    = rationalsrc.FromInts
	RationalParse       = rationalsrc.FromString
	NewComplex          = complexsrc.New
	ComplexFromReal     = complexsrc.FromReal
	TimeFromUnix        = timesrc.FromUnix
	TimeParse           = timesrc.Parse
	DurationFromSeconds = durationsrc.FromSeconds
	DurationParse       = durationsrc.Parse

	WithGoString = stringsrc.WithGoString
	WithPayload  = resultsrc.WithPayload
	WithError    = resultsrc.WithError
	WithInt      = numbersrc.WithInt
	WithFloat    = numbersrc.WithFloat
	WithPairs    = dictionarysrc.WithPairs
	WithStep     = rangesrc.WithStep
	Exclusive    = rangesrc.Exclusive
)
