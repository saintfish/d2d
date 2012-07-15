// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
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
	tt_guid
	tt_method_hresult
	tt_method
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
	r.RegisterToken(tt_method, "STDMETHOD_")
	r.RegisterToken(tt_method_hresult, "STDMETHOD")
	r.RegisterToken(tt_guid, "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}")
	r.RegisterToken(tt_punct, "[{};=()*,:\"]")
	r.RegisterToken(tt_id, "[A-Za-z_][A-Za-z0-9_]*")
	r.RegisterToken(tt_number, "(0x[0-9a-fA-F]+)|([0-9]+)")
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
	return name
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

func (s Struct) String() string {
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
		fmt.Println(s)
		return
	}
}

type Enum struct {
	Name  string
	Value []NameValue
}

type NameValue struct {
	Name  string
	Value string
}

var enumTempl = template.Must(template.New("enum").Parse(`
{{$n := .Name}}type {{$n}} uint32
const (
{{range .Value}}	{{.Name}} {{$n}} = {{.Value}}
{{end}})`))

func (e Enum) String() string {
	b := &bytes.Buffer{}
	enumTempl.Execute(b, e)
	return b.String()
}

func (p *D2DHeaderParser) ParseEnum(content []byte, name string) {
	for p.t, p.i, p.c = p.l.Lex(content); p.t != TokenTypeEOF; p.next() {
		if !p.isToken(tt_typedef, "") {
			continue
		}
		p.next()
		if !p.isToken(tt_enum, "") {
			continue
		}
		p.next()
		if !p.isToken(tt_id, name) {
			continue
		}
		var e Enum
		e.Name = name
		p.next()
		if !p.isToken(tt_punct, "{") {
			continue
		}
		p.next()
		for {
			p.expect(tt_id, "")
			var nv NameValue
			nv.Name = string(p.c)
			p.next()
			p.expect(tt_punct, "=")
			p.next()
			p.expect(tt_number, "")
			nv.Value = string(p.c)
			p.next()
			if !strings.HasSuffix(nv.Name, "FORCE_DWORD") {
				e.Value = append(e.Value, nv)
			}
			if p.isToken(tt_punct, "}") {
				break
			}
			p.expect(tt_punct, ",")
			p.next()
		}
		p.next()
		p.expect(tt_id, name)
		p.next()
		p.expect(tt_punct, ";")
		fmt.Println(e)
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

type Interface struct {
	Name          string
	Guid          string
	ConvertedGuid string
	Parent        string
	Methods       []InterfaceMethod
}

type InterfaceMethod struct {
	Name       string
	ResultType Type
	Params     []Param
}

func (im InterfaceMethod) ParamsSize() int {
	return len(im.Params) + 1
}

func (im InterfaceMethod) SyscallFunc() string {
	s := im.ParamsSize()
	if s <= 3 {
		return "Syscall"
	}
	return fmt.Sprintf("Syscall%d", s + im.ParamsPaddingSize())
}

func (im InterfaceMethod) ParamsPadding() []string {
	r := make([]string, im.ParamsPaddingSize())
	for i := range r {
		r[i] = "0"
	}
	return r
}

func (im InterfaceMethod) ParamsPaddingSize() int {
	s := im.ParamsSize()
	if s % 3 == 0 {
		return 0
	}
	return 3 - (s % 3)
}

type Type struct {
	Name             string
	IsPointer        bool
	IsPointerPointer bool
}

func (t Type) TypeEq(ty string) bool {
	return t.Name == ty
}

type Param struct {
	Name       string
	Type       Type
	IsOptional bool
	IsDeref    bool
	IsArray    bool
	IsIn       bool
	IsOut      bool
	Oops       string
}

var interfaceTempl = template.Must(template.New("interface").Parse(`
// {{.Guid}}
var IID_{{.Name}} = GUID{ {{.ConvertedGuid}} }

type {{.Name}}Vtbl struct {
	{{.Parent}}Vtbl
{{range .Methods}}	p{{.Name}} uintptr
{{end}}}

type {{.Name}} struct {
	*{{.Name}}Vtbl
}

type {{.Name}}Ptr struct {
	*{{.Name}}
}

func (this {{.Name}}Ptr) GUID() *GUID {
	return &IID_{{.Name}}
}

func (this {{.Name}}Ptr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.{{.Name}}))
}

func (this *{{.Name}}Ptr) SetRawPtr(raw uintptr) {
	this.{{.Name}} = (*{{.Name}})(unsafe.Pointer(raw))
}
{{$n := .Name}}{{range .Methods}}
func (this *{{$n}}Vtbl) {{.Name}}(
	ptr ComObjectPtr{{range .Params}},
	{{.Name}} {{if .Type.IsPointerPointer}}**{{else}}{{if .Type.IsPointer}}*{{end}}{{end}}{{.Type.Name}}{{end}}){{if .ResultType.TypeEq "void" | not}}
	({{if .ResultType.IsPointerPointer}}**{{else}}{{if .ResultType.IsPointer}}*{{end}}{{end}}{{.ResultType.Name}}){{end}} {
	var {{if .ResultType.TypeEq "HRESULT"}}ret{{else}}_{{end}}, _, _ = syscall.{{.SyscallFunc}}(
		this.p{{.Name}},
		{{.ParamsSize}},
		ptr.RawPtr(){{range .Params}},
		uintptr({{if .Type.IsPointer}}unsafe.Pointer({{.Name}}){{else}}{{.Name}}{{end}}){{end}}{{range .ParamsPadding}},
		{{.}}{{end}})
	{{if .ResultType.TypeEq "HRESULT"}}if ret != S_OK {
		panic(fmt.Sprintf("Fail to call {{.Name}}: %#x", ret))
	}{{end}}
	return
}
{{end}}
`))

func (ii Interface) String() string {
	b := &bytes.Buffer{}
	interfaceTempl.Execute(b, ii)
	return b.String()
}

func convertGuid(guid string) string {
	return fmt.Sprintf(
		"0x%s, 0x%s, 0x%s, [8]byte{0x%s, 0x%s, 0x%s, 0x%s, 0x%s, 0x%s, 0x%s, 0x%s}",
		guid[0:8], guid[9:13], guid[14:18], guid[19:21], guid[21:23],
		guid[24:26], guid[26:28], guid[28:30], guid[30:32], guid[32:34], guid[34:36])
}

func (p *D2DHeaderParser) consumeStatement() {
	var nestedBlock int
	var isEnd = false
	for !isEnd {
		if nestedBlock > 0 {
			if p.isToken(tt_punct, "}") {
				nestedBlock--
				if nestedBlock == 0 {
					isEnd = true
				}
			}
		} else {
			if p.isToken(tt_punct, ";") {
				isEnd = true
			}
		}
		if p.isToken(tt_punct, "{") {
			nestedBlock++
		}
		p.next()
	}
}

func (p *D2DHeaderParser) parseType() (t Type) {
	if p.isToken(tt_id, "CONST") {
		p.next()
	}
	p.expect(tt_id, "")
	t.Name = string(p.c)
	if goType, ok := cTypeToGoType[t.Name]; ok {
		t.Name = goType
	}
	p.next()
	if p.isToken(tt_punct, "*") {
		t.IsPointer = true
		p.next()
		if p.isToken(tt_punct, "*") {
			t.IsPointerPointer = true
			p.next()
		}
	}
	return
}

func (p *D2DHeaderParser) ParseInterface(content []byte, name string) {
	for p.t, p.i, p.c = p.l.Lex(content); p.t != TokenTypeEOF; p.next() {
		if !p.isToken(tt_interface, "") {
			continue
		}
		p.next()
		if !p.isToken(tt_id, "DECLSPEC_UUID") {
			continue
		}
		p.next()
		p.expect(tt_punct, "(")
		p.next()
		p.expect(tt_punct, "\"")
		p.next()
		p.expect(tt_guid, "")
		var ii Interface
		ii.Guid = string(p.c)
		ii.ConvertedGuid = convertGuid(string(p.c))
		p.next()
		p.expect(tt_punct, "\"")
		p.next()
		p.expect(tt_punct, ")")
		p.next()
		p.expect(tt_id, "DECLSPEC_NOVTABLE")
		p.next()
		if !p.isToken(tt_id, name) {
			continue
		}
		ii.Name = string(p.c)
		p.next()
		p.expect(tt_punct, ":")
		p.next()
		p.expect(tt_id, "public")
		p.next()
		p.expect(tt_id, "")
		ii.Parent = string(p.c)
		p.next()
		p.expect(tt_punct, "{")
		p.next()
		for {
			if p.isToken(tt_punct, "}") {
				break
			}
			var method InterfaceMethod
			if p.isToken(tt_method_hresult, "") {
				p.next()
				p.expect(tt_punct, "(")
				p.next()
				p.expect(tt_id, "")
				method.Name = string(p.c)
				method.ResultType.Name = "HRESULT"
				p.next()
				p.expect(tt_punct, ")")
				p.next()
			} else if p.isToken(tt_method, "") {
				p.next()
				p.expect(tt_punct, "(")
				p.next()
				method.ResultType = p.parseType()
				p.expect(tt_punct, ",")
				p.next()
				p.expect(tt_id, "")
				method.Name = string(p.c)
				p.next()
				p.expect(tt_punct, ")")
				p.next()
			} else {
				p.consumeStatement()
				continue
			}
			p.expect(tt_punct, "(")
			p.next()
			for {
				if p.isToken(tt_punct, ")") {
					break
				}
				var attribute = string(p.c)
				var param Param
				if strings.HasPrefix(attribute, "__") {
					switch attribute {
					case "__in":
						param.IsIn = true
					case "__in_opt":
						param.IsIn = true
						param.IsOptional = true
					case "__in_ecount":
						param.IsIn = true
						param.IsArray = true
						p.next()
						p.expect(tt_punct, "(")
						p.next()
						p.expect(tt_id, "")
						p.next()
						p.expect(tt_punct, ")")
					case "__in_ecount_opt":
						param.IsIn = true
						param.IsArray = true
						param.IsOptional = true
						p.next()
						p.expect(tt_punct, "(")
						p.next()
						p.expect(tt_id, "")
						p.next()
						p.expect(tt_punct, ")")
					case "__out":
						param.IsOut = true
					case "__out_opt":
						param.IsOut = true
						param.IsOptional = true
					case "__deref_out":
						param.IsOut = true
						param.IsDeref = true
					case "__deref_out_opt":
						param.IsOut = true
						param.IsDeref = true
						param.IsOptional = true
					case "__out_ecount":
						param.IsOut = true
						param.IsArray = true
						p.next()
						p.expect(tt_punct, "(")
						p.next()
						p.expect(tt_id, "")
						p.next()
						p.expect(tt_punct, ")")
					case "__inout":
						param.IsIn = true
						param.IsOut = true
					case "__range":
						param.IsIn = true
						for !p.isToken(tt_punct, ")") {
							p.next()
						}
					default:
						param.Oops = attribute
					}
					p.next()
				} else {
					param.IsIn = true
				}
				param.Type = p.parseType()
				p.expect(tt_id, "")
				param.Name = string(p.c)
				p.next()
				if p.isToken(tt_punct, "=") {
					for !p.isToken(tt_punct, ",") && !p.isToken(tt_punct, ")") {
						p.next()
					}
				}
				if p.isToken(tt_punct, ",") {
					p.next()
				}
				method.Params = append(method.Params, param)
			}
			p.expect(tt_punct, ")")
			p.next()
			for p.isToken(tt_id, "PURE") || p.isToken(tt_id, "CONST") {
				p.next()
			}
			p.expect(tt_punct, ";")
			p.next()
			ii.Methods = append(ii.Methods, method)
		}
		p.next()
		p.expect(tt_punct, ";")
		fmt.Println(ii)
		return
	}
}
