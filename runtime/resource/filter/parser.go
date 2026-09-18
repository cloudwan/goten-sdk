package filter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"

	"github.com/cloudwan/goten-sdk/runtime/resource/filter/ebnf"
)

// bits stolen from: https://github.com/alecthomas/participle/blob/v0.1.0/_examples/sql/main.go

type Boolean bool

func (b *Boolean) Capture(values []string) error {
	*b = strings.ToUpper(values[0]) == "TRUE"
	return nil
}

type CompareOperator string

type CompositeOperator string

var (
	Eq  = CompareOperator("=")
	Neq = CompareOperator("!=")
	Gt  = CompareOperator(">")
	Gte = CompareOperator(">=")
	Lt  = CompareOperator("<")
	Lte = CompareOperator("<=")

	OR  = CompositeOperator("OR")
	AND = CompositeOperator("AND")
)

var validOperators = []string{
	string(Eq), string(Neq), string(Gt), string(Gte), string(Lt), string(Lte),
}

func (op *CompareOperator) Capture(values []string) error {
	// Reached while parsing a client-supplied filter, so a surprise here has to
	// be reported rather than panicked on - and the empty case has to be checked
	// before values is indexed.
	if len(values) != 1 {
		return fmt.Errorf("expected a single comparison operator, got %d: '%s'",
			len(values), strings.Join(values, " "))
	}
	ops := values[0]
	if ops == "==" {
		ops = "="
	}

	for _, valid := range validOperators {
		if ops == valid {
			*op = CompareOperator(ops)
			return nil
		}
	}

	// Reports what was captured, not *op: that still holds whatever the operator
	// held before this call, which for a fresh parse is nothing at all.
	return fmt.Errorf("invalid operator: '%s', (expected: %s )", values[0], strings.Join(validOperators, " | "))
}

func (op *CompareOperator) MatchCompareResult(cmp int) bool {
	switch *op {
	case Eq:
		return cmp == 0
	case Neq:
		return cmp != 0
	case Gt:
		return cmp < 0
	case Gte:
		return cmp <= 0
	case Lt:
		return cmp > 0
	case Lte:
		return cmp >= 0
	default:
		panic("Invalid compare operator")
	}
}

type Expression struct {
	And []AndCondition `@@ { ("OR" | "|") @@ }`
}

type AndCondition struct {
	Or []Condition `@@ { [("AND" | "+" )] @@ }`
}

// Condition holds the grammar's two recursive productions. checkNestingDepth
// counts them on the raw input, before the parser recurses, so any change to
// which productions recurse - or to the brackets and keywords that enter them -
// has to be mirrored there. The same applies to ConditionRHS and Map below.
type Condition struct {
	Not           *Condition        `  "NOT" @@`
	SubExpression *Expression       `| "(" @@ ")"`
	Operand       *ConditionOperand `|  @@`
}

type ConditionOperand struct {
	FieldPath    string       `@FieldPath`
	ConditionRHS ConditionRHS `@@ `
}

// ConditionRHS carries the keywords endsCondition and opensValue match on, and
// the non-prefix "NOT" that isPrefixNot has to exclude. Keep checkNestingDepth
// in step with it.
type ConditionRHS struct {
	Compare  *Compare  `  @@`
	Is       *Is       `| "IS" @@`
	In       Array     `| "IN" ("(" @@ { "," @@ } ")" | "[" @@ { "," @@ } "]")`
	NotIn    Array     `| "NOT" "IN" ("(" @@ { "," @@ } ")" | "[" @@ { "," @@ } "]")`
	Contains *Contains `| ("CONTAINS" | "HAS") @@`
	Like     *string   `| "LIKE" @String`
}

func (crhs ConditionRHS) JSONValue() ([]byte, error) {
	var value interface{}
	if crhs.Compare != nil {
		value = crhs.Compare.Value
	} else if crhs.In != nil {
		value = crhs.In
	} else if crhs.NotIn != nil {
		value = crhs.NotIn
	} else if crhs.Like != nil {
		value = *crhs.Like
	} else if crhs.Contains != nil {
		if crhs.Contains.All != nil {
			value = crhs.Contains.All
		} else if crhs.Contains.Any != nil {
			value = crhs.Contains.Any
		} else if crhs.Contains.Value != nil {
			value = crhs.Contains.Value
		}
	}
	return json.Marshal(value)
}

type Compare struct {
	Operator CompareOperator `@CompareOperator`
	Value    Value           `@@`
}

type Contains struct {
	Any   Array  `  "ANY" ("(" @@ { "," @@ } ")" | "[" @@ { "," @@ } "]")`
	All   Array  `| "ALL" ("(" @@ { "," @@ } ")" | "[" @@ { "," @@ } "]")`
	Value *Value `| ["VALUE"] @@`
}

func (c *Contains) GetArray() Array {
	if c.Any != nil {
		return c.Any
	} else if c.All != nil {
		return c.All
	} else {
		panic(fmt.Errorf("conditon %v has no array", c))
	}
}

type Is struct {
	Not  bool `[ @"NOT" ]`
	Null bool `(  @"NULL"`
	NaN  bool ` | @"NAN" )`
}

type Array []Value

type Value struct {
	Null    bool     `(  @"NULL"`
	Boolean *Boolean ` | @("TRUE" | "FALSE")`
	Number  *string  ` | @Number`
	String  *string  ` | @String`
	Map     *Map     ` | @@ )`
}

// Map recurses through Value, which is why "{" counts towards the nesting depth
// in checkNestingDepth alongside "(" and "[".
type Map struct {
	Entries []*MapEntry `"{" ( @@ ( ( "," )? @@ )* )? "}"`
}

type MapEntry struct {
	Key   *Value `@@`
	Value *Value `":"? @@`
}

func (v Value) MarshalJSON() ([]byte, error) {
	if v.Null {
		return json.Marshal(nil)
	} else if v.Boolean != nil {
		return json.Marshal(v.Boolean)
	} else if v.Number != nil {
		return []byte(*v.Number), nil
	} else if v.String != nil {
		return json.Marshal(v.String)
	} else if v.Map != nil {
		jsonEntries := make([]string, 0, len(v.Map.Entries))
		for _, entry := range v.Map.Entries {
			jsonKey, err := entry.Key.MarshalJSON()
			if err != nil {
				return nil, errors.New("error marshaling map key: " + err.Error())
			}
			jsonValue, err := entry.Value.MarshalJSON()
			if err != nil {
				return nil, errors.New("error marshaling map value: " + err.Error())
			}
			jsonEntries = append(jsonEntries, fmt.Sprintf("%s:%s", string(jsonKey), string(jsonValue)))
		}
		value := fmt.Sprintf("{%s}", strings.Join(jsonEntries, ","))
		return []byte(value), nil
	} else {
		return nil, errors.New("primitive value contains none of Null, Boolean, Number nor String")
	}
}

var (
	ebnfDef = ebnf.New(`
Comment = "--" { "\u0000"…"\uffff"-"\n" } .
FieldPath = Ident { dot Ident } .
Ident = (alpha | "_") { "_" | "-" | "/" | alpha | digit } { "_" | "/" | alpha | digit } .
String = "\"" { "\u0000"…"\uffff"-"\""-"\\" | "\\" any } "\"" .
CompareOperator = opchar { opchar } .
Number = [ "-" | "+" ] ("." | digit) {"." | digit} [("E" | "e") [ "-" | "+" ] digit { digit }] .
Punct = "!"…"/" | ":"…"@" | "["…` + "\"`\"" + ` | "{"…"~" .
Whitespace = " " | "\t" | "\n" | "\r" .
alpha = "a"…"z" | "A"…"Z" .
digit = "0"…"9" .
opchar = ("<" | ">" | "=" | "!") .
dot = ("\\." | ".") .
any = "\u0000"…"\uffff" .
`)

	filterParser = participle.MustBuild[Expression](
		participle.Lexer(ebnfDef),
		participle.Unquote("String"),
		participle.Elide("Whitespace", "Comment"),
		participle.Map(func(token lexer.Token) (lexer.Token, error) {
			token.Value = strings.Replace(token.Value, `\.`, ".", -1)
			return token, nil
		}, "FieldPath"),
		participle.CaseInsensitive("Ident", "FieldPath"),
	)
)

// MaxNestingDepth bounds how deeply a filter expression may nest. The grammar
// is recursive in two ways - "(" Expression ")" for sub-expressions and "{" for
// map values, and the prefix "NOT" Condition, which recurses without any
// bracket to show for it - and the generated parser is recursive descent, so
// nesting in the input maps directly onto stack frames. Filters arrive from API
// clients, so without this the request alone decides how deep the server
// recurses.
//
// Backends that lower a parsed filter cap depth again, on the condition tree
// rather than on the input. A bracket counted once here can cost two levels of
// condition tree, so those caps have to be derived from this constant rather
// than set alongside it; see the mongo backend's maxFilterConditionDepth.
const MaxNestingDepth = 32

// MaxLength bounds a filter expression in bytes. The depth guard below and the
// per-backend clause budgets both act on a filter that has already been parsed,
// so neither bounds what parsing itself costs: a flat filter is not deep and
// lowers to few clauses, yet "a=1 AND " repeated fills the condition tree with
// one node per repetition before any of those limits is consulted. Only the
// transport's message size stands in the way otherwise, and that is measured in
// megabytes.
//
// The limit is set well above any filter a client would write by hand or a
// builder would assemble - a page of scope references is a few kilobytes - so
// that it only ever catches abuse.
const MaxLength = 128 * 1024

// errNestedTooDeep is what checkNestingDepth rejects with. It is built from a
// constant, so it is built once rather than on every filter parsed.
var errNestedTooDeep = fmt.Errorf(
	"error when parsing filter expression: nested deeper than %d levels, which is not supported",
	MaxNestingDepth)

// Token types the nesting-depth scan needs to tell apart. They are resolved
// once, and a missing one is a programming error rather than something to fall
// back on: an absent symbol would compare equal to no token at all - types are
// assigned downwards from lexer.EOF, so the zero value is never a real type -
// and the guard would silently stop firing.
var (
	tokWhitespace      = mustSymbol("Whitespace")
	tokComment         = mustSymbol("Comment")
	tokString          = mustSymbol("String")
	tokNumber          = mustSymbol("Number")
	tokCompareOperator = mustSymbol("CompareOperator")
)

func mustSymbol(name string) lexer.TokenType {
	tokenType, ok := ebnfDef.Symbols()[name]
	if !ok {
		panic(fmt.Sprintf("filter grammar defines no %q token, "+
			"which the nesting-depth guard depends on", name))
	}
	return tokenType
}

// checkNestingDepth rejects filters that nest deeper than MaxNestingDepth. It
// deliberately runs before parsing: once the parser has recursed, the stack is
// already spent, so a check on the parsed tree would be too late. Lexing is a
// flat loop over the input - it does not recurse on nesting - so it is safe to
// run first, and counting lexemes rather than bytes is what keeps this in step
// with the parser: brackets inside a string literal or a comment are then
// neither counted nor able to confuse the scan about where the next literal
// begins.
//
// Brackets are matched by value: only Punct can carry one, so the value is
// unambiguous. Keywords are not, and are handled in endsCondition.
func checkNestingDepth(data []byte) error {
	if !canNest(data) {
		return nil
	}
	lex, err := ebnfDef.Lex("", bytes.NewReader(data))
	if err != nil {
		return nil // not lexable at all - the parser reports it
	}
	// next yields the tokens the parser sees; the elided ones carry no grammar.
	// A lexical error ends the scan: it belongs to the parser, which stops at the
	// same place, and everything before it has been counted here already.
	next := func() lexer.Token {
		for {
			token, err := lex.Next()
			if err != nil {
				return lexer.Token{Type: lexer.EOF}
			}
			if token.Type == tokWhitespace || token.Type == tokComment {
				continue
			}
			return token
		}
	}

	// brackets is the depth accounted for by the brackets still open, each
	// carrying the run of NOTs that preceded it; notRun is the run of NOTs still
	// outstanding. Both are frames the parser still holds, so the depth to check
	// against the limit is their sum.
	brackets, notRun := 0, 0
	// inherited remembers, per open bracket, the NOT run that bracket was opened
	// under, so closing it unwinds exactly the frames it opened.
	var inherited []int

	eof := lexer.Token{Type: lexer.EOF}
	prev2, prev, token, lookahead := eof, eof, next(), next()
	for !token.EOF() {
		switch {
		case token.Value == "(" || token.Value == "[" || token.Value == "{":
			brackets += notRun + 1
			inherited = append(inherited, notRun)
			notRun = 0
			if brackets > MaxNestingDepth {
				return errNestedTooDeep
			}
		case token.Value == ")" || token.Value == "]" || token.Value == "}":
			// A closer with no matching open frame is ignored, so unbalanced
			// closers cannot buy headroom for a deeper run of openers later in the
			// same input. That guard is also what keeps brackets non-negative:
			// inherited is a strict stack, so each pop subtracts exactly what its
			// matching push added, leaving brackets the sum over open frames of
			// inherited[i]+1. No separate floor is needed.
			if len(inherited) > 0 {
				brackets -= inherited[len(inherited)-1] + 1
				inherited = inherited[:len(inherited)-1]
			}
			notRun = 0
		case isPrefixNot(prev, token, lookahead):
			notRun++
			if brackets+notRun > MaxNestingDepth {
				return errNestedTooDeep
			}
		case endsCondition(prev2, prev, token):
			notRun = 0
		}
		prev2, prev, token, lookahead = prev, token, lookahead, next()
	}
	return nil
}

// canNest reports whether data could nest at all, so that the common filter -
// which cannot, and is parsed on a per-request path - is spared a second pass
// over the input. Every recursive production in the grammar is entered either
// by an opening bracket or by the prefix NOT, so an input carrying neither is
// necessarily flat. This has to be revisited if the grammar gains a recursive
// production reachable without one of those.
//
// Being over-eager is safe: a bracket in a string literal or a comment costs
// the full scan, which then declines to count it. Being under-eager is not,
// hence matching NOT rather than just its letters - "AND" carries an "N".
func canNest(data []byte) bool {
	if bytes.ContainsAny(data, "([{") {
		return true
	}
	// Case-insensitive search for "NOT" without copying the input. OR-ing 0x20
	// lower-cases ASCII letters, and no other byte maps onto one of these.
	for i := 0; i+2 < len(data); i++ {
		if data[i]|0x20 == 'n' && data[i+1]|0x20 == 'o' && data[i+2]|0x20 == 't' {
			return true
		}
	}
	return false
}

// isPrefixNot reports whether a NOT opens a nested condition. Only the prefix
// form recurses - Condition.Not captures another Condition - while the NOT of
// "<field> NOT IN (...)" and of "<field> IS NOT NULL" belongs to a condition's
// right-hand side and costs no stack.
func isPrefixNot(prev, token, lookahead lexer.Token) bool {
	return strings.EqualFold(token.Value, "NOT") &&
		!strings.EqualFold(lookahead.Value, "IN") &&
		!strings.EqualFold(prev.Value, "IS")
}

// endsCondition reports whether a token completes a condition's right-hand
// side, which is where the NOT frames prefixing that condition unwind. Until
// then those frames are live, so the run has to be carried across everything
// between - the field path, the operator, and any brackets the value opens.
// Every right-hand side ends either at a literal value, at the NULL or NAN of
// an IS, or at the bracket closing a list or a map; the last of those is a
// closer, which the scan unwinds anyway.
//
// A string or a number is unambiguously a value: neither can be a field path.
// The bare keywords are not, because NULL, NAN, TRUE and FALSE lex as FieldPath
// like any other identifier - the grammar matches them by value - so a field of
// the same name is indistinguishable from the literal by token alone. Those are
// therefore only taken as values where the grammar admits one.
func endsCondition(prev2, prev, token lexer.Token) bool {
	if token.Type == tokString || token.Type == tokNumber {
		return true
	}
	switch strings.ToUpper(token.Value) {
	case "NULL", "NAN", "TRUE", "FALSE":
		return opensValue(prev2, prev)
	}
	return false
}

// opensValue reports whether the grammar admits a literal value after prev.
// Brackets and separators need no entry: the value they introduce belongs to a
// list or a map, and opening that bracket already unwound the run.
func opensValue(prev2, prev lexer.Token) bool {
	if prev.Type == tokCompareOperator {
		return true
	}
	switch strings.ToUpper(prev.Value) {
	case "IS", "LIKE", "CONTAINS", "HAS", "VALUE", "ANY", "ALL", "IN":
		return true
	case "NOT":
		// Only the NOT of "IS NOT NULL" introduces a value. A prefix NOT
		// introduces a condition, and treating what follows it as a value would
		// unwind a run of them while the parser still holds every frame.
		return strings.EqualFold(prev2.Value, "IS")
	}
	return false
}

func Parse(data []byte) (*Expression, error) {
	filter := &Expression{}
	if len(data) == 2 && data[0] == '(' && data[1] == ')' {
		filter.And = append(filter.And, AndCondition{})
	} else {
		if len(data) > MaxLength {
			return nil, fmt.Errorf(
				"error when parsing filter expression: %d bytes long, which is over the %d byte limit",
				len(data), MaxLength)
		}
		if err := checkNestingDepth(data); err != nil {
			return nil, err
		}
		var err error
		if filter, err = filterParser.Parse("", bytes.NewReader(data)); err != nil {
			return nil, fmt.Errorf("error when parsing filter expression: %q: %s", string(data), err)
		}
	}
	return filter, nil
}
