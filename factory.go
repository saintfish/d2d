// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

import (
	"fmt"
	"github.com/AllenDang/w32"
	"syscall"
	"unsafe"
)

var (
	modd2d1               = syscall.NewLazyDLL("d2d1.dll")
	procD2D1CreateFactory = modd2d1.NewProc("D2D1CreateFactory")
)

type FACTORY_TYPE uint32

const (
	// The resulting factory and derived resources may only be invoked serially.
	// Reference counts on resources are interlocked, however, resource and render
	// target state is not protected from multi-threaded access.
	FACTORY_TYPE_SINGLE_THREADED FACTORY_TYPE = 0
	// The resulting factory may be invoked from multiple threads. Returned resources
	// use interlocked reference counting and their state is protected.
	FACTORY_TYPE_MULTI_THREADED FACTORY_TYPE = 1
)

// Indicates the debug level to be outputed by the debug layer.
type DEBUG_LEVEL uint32

const (
	DEBUG_LEVEL_NONE        DEBUG_LEVEL = 0
	DEBUG_LEVEL_ERROR       DEBUG_LEVEL = 1
	DEBUG_LEVEL_WARNING     DEBUG_LEVEL = 2
	DEBUG_LEVEL_INFORMATION DEBUG_LEVEL = 3
)

type FACTORY_OPTIONS struct {
	DebugLevel DEBUG_LEVEL
}

var (
	// Interface ID of ID2D1Factory "06152247-6f50-465a-9245-118bfd3b6007"
	IID_ID2D1Factory = &w32.GUID{0x06152247, 0x6f50, 0x465a, [8]byte{0x92, 0x45, 0x11, 0x8b, 0xfd, 0x3b, 0x60, 0x07}}
)

type Factory struct {
	factory uintptr
}

// CreateFactory creates instance of factory
func CreateFactory(factoryType FACTORY_TYPE, factoryOption *FACTORY_OPTIONS) (f *Factory, e error) {
	var tmp uintptr
	ret, _, _ := procD2D1CreateFactory.Call(
		uintptr(factoryType),
		uintptr(unsafe.Pointer(IID_ID2D1Factory)),
		uintptr(unsafe.Pointer(factoryOption)),
		uintptr(unsafe.Pointer(&tmp)))
	if ret == w32.S_OK {
		f = new(Factory)
		f.factory = tmp
	} else {
		e = fmt.Errorf("Fail to create factory: %#x", ret)
	}
	return
}
