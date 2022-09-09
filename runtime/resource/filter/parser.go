package filter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
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
	ops := values[0]
	if ops == "==" {
		ops = "="
	}
	if len(values) > 1 {
		panic(fmt.Sprintf("vals: '%q'", strings.Join(values, ", ")))
	}

	for _, valid := range validOperators {
		if ops == valid {
			*op = CompareOperator(ops)
			return nil
		}
	}

	return fmt.Errorf("invalid operator: '%s', (expected: %s )", *op, strings.Join(validOperators, " | "))
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

type Condition struct {
	Not           *Condition        `  "NOT" @@`
	SubExpression *Expression       `| "(" @@ ")"`
	Operand       *ConditionOperand `|  @@`
}

type ConditionOperand struct {
	FieldPath    string       `@FieldPath`
	ConditionRHS ConditionRHS `@@ `
}

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
	filterLexer = lexer.Must(ebnf.New(`
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
`))

	filterParser = participle.MustBuild(
		&Expression{},
		participle.Lexer(filterLexer),
		participle.Unquote("String"),
		participle.Elide("Whitespace", "Comment"),
		participle.Map(func(token lexer.Token) (lexer.Token, error) {
			token.Value = strings.Replace(token.Value, `\.`, ".", -1)
			return token, nil
		}, "FieldPath"),
		participle.CaseInsensitive("Ident", "FieldPath"),
	)
)

func Parse(data []byte) (*Expression, error) {
	filter := &Expression{}
	if err := filterParser.Parse(bytes.NewReader(data), filter); err != nil {
		return nil, fmt.Errorf("error when parsing filter expression: \"%q\": %s", string(data), err)
	}
	return filter, nil
}
