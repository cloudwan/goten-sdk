// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"io"
	"strconv"
	"text/scanner"
)

type parser struct {
	errors  errorList
	scanner scanner.Scanner
	pos     scanner.Position // token position
	tok     rune             // one token look-ahead
	lit     string           // token literal
}

func (p *parser) next() {
	p.tok = p.scanner.Scan()
	p.pos = p.scanner.Position
	p.lit = p.scanner.TokenText()
}

func (p *parser) error(pos scanner.Position, msg string) {
	p.errors = append(p.errors, newError(pos, msg))
}

func (p *parser) errorExpected(pos scanner.Position, msg string) {
	msg = `expected "` + msg + `"`
	if pos.Offset == p.pos.Offset {
		// the error happened at the current position;
		// make the error message more specific
		msg += ", found " + scanner.TokenString(p.tok)
		if p.tok < 0 {
			msg += " " + p.lit
		}
	}
	p.error(pos, msg)
}

func (p *parser) expect(tok rune) {
	pos := p.pos
	if p.tok != tok {
		p.errorExpected(pos, scanner.TokenString(tok))
	}
	p.next() // make progress in any case
}

func (p *parser) parseIdentifier() *Name {
	pos := p.pos
	name := p.lit
	p.expect(scanner.Ident)
	return &Name{pos, name}
}

func (p *parser) parseToken() *Token {
	pos := p.pos
	value := ""
	if p.tok == scanner.String {
		value, _ = strconv.Unquote(p.lit)
		// Unquote may fail with an error, but only if the scanner found
		// an illegal string in the first place. In this case the error
		// has already been reported.
		p.next()
	} else {
		p.expect(scanner.String)
	}
	return &Token{pos, value}
}

const ellipsis = '…' // U+2026, the horizontal ellipsis character

func (p *parser) parseRange() (x Expression) {
	tok := p.parseToken()
	x = tok
	if p.tok == ellipsis {
		p.next()
		end := p.parseToken()
		var exclude Expression
		if p.tok == '-' {
			p.next()
			exclude = p.parseRange()
		}
		x = &Range{tok, end, exclude}
	} else if p.tok == '-' {
		p.next()
		x = &Range{tok, nil, p.parseRange()}
	}
	return x
}

// ParseTerm returns nil if no term was found.
func (p *parser) parseTerm() (x Expression) {
	pos := p.pos

	switch p.tok {
	case scanner.Ident:
		x = p.parseIdentifier()

	case scanner.String:
		x = p.parseRange()

	case '(':
		p.next()
		x = &Group{pos, p.parseExpression()}
		p.expect(')')

	case '[':
		p.next()
		x = &Option{pos, p.parseExpression()}
		p.expect(']')

	case '{':
		p.next()
		x = &Repetition{pos, p.parseExpression()}
		p.expect('}')
	default:
	}

	return x
}

func (p *parser) parseSequence() Expression {
	var list Sequence

	for x := p.parseTerm(); x != nil; x = p.parseTerm() {
		list = append(list, x)
	}

	// no need for a sequence if list.Len() < 2
	switch len(list) {
	case 0:
		p.errorExpected(p.pos, "term")
		return &Bad{p.pos, "term expected"}
	case 1:
		return list[0]
	}

	return list
}

func (p *parser) parseExpression() Expression {
	var list Alternative

	for {
		list = append(list, p.parseSequence())
		if p.tok != '|' {
			break
		}
		p.next()
	}
	// len(list) > 0

	// no need for an Alternative node if list.Len() < 2
	if len(list) == 1 {
		return list[0]
	}

	return list
}

func (p *parser) parseProduction() *Production {
	name := p.parseIdentifier()
	p.expect('=')
	var expr Expression
	if p.tok != '.' {
		expr = p.parseExpression()
	}
	p.expect('.')
	return &Production{name, expr}
}

func (p *parser) parse(filename string, src io.Reader) Grammar {
	p.scanner.Init(src)
	p.scanner.Filename = filename
	p.next() // initializes pos, tok, lit

	grammar := Grammar{Index: map[string]*Production{}}
	for p.tok != scanner.EOF {
		prod := p.parseProduction()
		name := prod.Name.String
		if _, found := grammar.Index[name]; !found {
			grammar.Index[name] = prod
			grammar.Productions = append(grammar.Productions, &NamedProduction{
				Name:       name,
				Production: prod,
			})
		} else {
			p.error(prod.Pos(), name+" declared already")
		}
	}

	return grammar
}

// Parse parses a set of EBNF productions from source src.
// It returns a set of productions. Errors are reported
// for incorrect syntax and if a production is declared
// more than once; the filename is used only for error
// positions.
func Parse(filename string, src io.Reader) (Grammar, error) {
	var p parser
	grammar := p.parse(filename, src)
	return grammar, p.errors.Err()
}
