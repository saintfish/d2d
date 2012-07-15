// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func usage() {
	fmt.Println("gen <struct|enum> <name>")
}

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		usage()
		return
	}
	what := flag.Arg(0)
	name := flag.Arg(1)

	p := NewD2DHeaderParser()
	f, e := os.Open("D2D1_preprocessed.txt")
	if e != nil {
		panic(e)
	}
	b, e := ioutil.ReadAll(f)
	if e != nil {
		panic(e)
	}
	switch what {
	case "struct":
		p.ParseStruct(b, name)
	case "enum":
		p.ParseEnum(b, name)
	case "interface":
		p.ParseInterface(b, name)
	}
}
