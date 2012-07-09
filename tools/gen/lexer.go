// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"regexp"
)

type TokenType int32

const (
	TokenTypeError   TokenType = -1
	TokenTypeEOF     TokenType = -2
	TokenTypeIgnored TokenType = -3
)

type RegexpLexer interface {
	RegisterToken(t TokenType, re string)
	Lex([]byte) (t TokenType, index int, match []byte)
	GetNextToken() (t TokenType, index int, match []byte)
}

type regexpLexer struct {
	content        []byte
	tt             []TokenType
	re             []*regexp.Regexp
	literalPrefix  [][]byte
	prefixComplete []bool
	cur            int
}

func NewRegexpLexer() (l RegexpLexer) {
	r := &regexpLexer{}
	r.re = make([]*regexp.Regexp, 0, 10)
	r.literalPrefix = make([][]byte, 0, 10)
	r.prefixComplete = make([]bool, 0, 10)
	r.cur = 0
	return r
}

func (l *regexpLexer) RegisterToken(t TokenType, re string) {
	reNoPrefix := regexp.MustCompile(re)
	rePrefix := regexp.MustCompile("^" + re)
	l.tt = append(l.tt, t)
	l.re = append(l.re, rePrefix)
	prefix, complete := reNoPrefix.LiteralPrefix()
	l.literalPrefix = append(l.literalPrefix, []byte(prefix))
	l.prefixComplete = append(l.prefixComplete, complete)
}

func (l *regexpLexer) Lex(content []byte) (t TokenType, index int, match []byte) {
	l.content = content
	l.cur = 0
	return l.GetNextToken()
}

func (l *regexpLexer) GetNextToken() (t TokenType, index int, match []byte) {
	for t, index, match = l.getNextToken(); t == TokenTypeIgnored; t, index, match = l.getNextToken() {
	}
	return
}

func (l *regexpLexer) getNextToken() (t TokenType, index int, match []byte) {
	if len(l.re) == 0 {
		panic("GetNextToken called before any token being registered.")
	}
	if l.content == nil {
		panic("GetNextToken called before Lex.")
	}
	if l.cur >= len(l.content) {
		return TokenTypeEOF, len(l.content), nil
	}
	size := len(l.re)
	cur_buffer := l.content[l.cur:]
	for i := 0; i < size; i++ {
		if bytes.HasPrefix(cur_buffer, l.literalPrefix[i]) {
			if l.prefixComplete[i] {
				t = l.tt[i]
				index = l.cur
				match = cur_buffer[:len(l.literalPrefix[i])]
				l.cur += len(match)
				return
			}
			m := l.re[i].Find(cur_buffer)
			if m != nil {
				t = l.tt[i]
				index = l.cur
				match = m
				l.cur += len(match)
				return
			}
		}
	}
	t = TokenTypeError
	index = l.cur
	match = nil
	l.cur += 1
	return
}
