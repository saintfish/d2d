// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

const (
	tt_typedef TokenType = iota
	tt_struct
	tt_interface
	tt_enum
	tt_punct
	tt_id
	tt_number
)

var cTypeToGoType = map[string]string{
	"FLOAT":  "float32",
	"UINT32": "uint32",
}

func CreateCHeaderLexer() RegexpLexer {
	r := NewRegexpLexer()
	r.RegisterToken(TokenTypeIgnored, "//.*?\\n")
	r.RegisterToken(TokenTypeIgnored, "#.*?\\n")
	r.RegisterToken(TokenTypeIgnored, "\\s+")
	r.RegisterToken(TokenTypeIgnored, "\\s+")
	r.RegisterToken(tt_typedef, "typedef")
	r.RegisterToken(tt_struct, "struct")
	r.RegisterToken(tt_interface, "interface")
	r.RegisterToken(tt_enum, "enum")
	r.RegisterToken(tt_punct, "[{};=()*]")
	r.RegisterToken(tt_id, "[A-Za-z_][A-Za-z0-9_]*")
	r.RegisterToken(tt_number, "([0-9]+)|(0x[0-9a-fA-F]+)")
	return r
}

func camelName(name string) string {
	if name[0] == '_' {
		return "A" + name[1:]
	}
	if s := name[0]; s >= 'a' && s <= 'z' {
		s = s - 'a' + 'A'
		return string(s) + name[1:]
	}
	log.Panicf("Unexpected name: %s", name)
	return ""
}

type D2DHeaderParser struct {
	t TokenType
	i int
	c []byte
	l RegexpLexer
}

func NewD2DHeaderParser() *D2DHeaderParser {
	p := &D2DHeaderParser{}
	p.l = CreateCHeaderLexer()
	return p
}

type Struct struct {
	Name  string
	Field []TypeName
}

type TypeName struct {
	Type      string
	Name      string
	IsPointer bool
}

var structTempl = template.Must(template.New("struct").Parse(`
type {{.Name}} struct {
{{range .Field}}	{{.Name}} {{if .IsPointer}}*{{end}}{{.Type}}
{{end}}}`))

func (s *Struct) String() string {
	b := &bytes.Buffer{}
	structTempl.Execute(b, s)
	return b.String()
}

func (p *D2DHeaderParser) ParseStruct(content []byte, name string) {
	for p.t, p.i, p.c = p.l.Lex(content); p.t != TokenTypeEOF; p.next() {
		if !p.isToken(tt_typedef, "") {
			continue
		}
		p.next()
		if !p.isToken(tt_struct, "") {
			continue
		}
		p.next()
		if !p.isToken(tt_id, name) {
			continue
		}
		var s Struct
		s.Name = name
		p.next()
		if !p.isToken(tt_punct, "{") {
			continue
		}
		p.next()
		for !p.isToken(tt_punct, "}") {
			if p.isToken(tt_id, "__field_ecount_opt") {
				p.next()
				p.expect(tt_punct, "(")
				p.next()
				p.expect(tt_number, "1")
				p.next()
				p.expect(tt_punct, ")")
				p.next()
			}
			p.expect(tt_id, "")
			var tn TypeName
			tn.Type = string(p.c)
			if goType, ok := cTypeToGoType[tn.Type]; ok {
				tn.Type = goType
			}
			p.next()
			if p.isToken(tt_punct, "*") {
				tn.IsPointer = true
				p.next()
			} else {
				tn.IsPointer = false
			}
			p.expect(tt_id, "")
			tn.Name = camelName(string(p.c))
			p.next()
			p.expect(tt_punct, ";")
			p.next()
			s.Field = append(s.Field, tn)
		}
		p.next()
		p.expect(tt_id, name)
		p.next()
		p.expect(tt_punct, ";")
		fmt.Println(&s)
		return
	}
}

func (p *D2DHeaderParser) next() {
	// log.Printf("%s", p.c)
	p.t, p.i, p.c = p.l.GetNextToken()
}

func (p *D2DHeaderParser) isToken(tt TokenType, literal string) bool {
	if p.t == tt && (literal == "" || literal == string(p.c)) {
		return true
	}
	return false
}

func (p *D2DHeaderParser) expect(tt TokenType, literal string) {
	if !p.isToken(tt, literal) {
		log.Panicf("Unexpected token %v,%s, expect %v:%s", p.t, p.c, tt, literal)
	}
}
