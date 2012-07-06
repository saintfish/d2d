// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

import (
	"syscall"
	"unsafe"
)

var (
	// "2cd90691-12e2-11dc-9fed-001143a055f9"
	IID_ID2D1Resource = GUID{0x2cd90691, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}
)

type ID2D1ResourceVtbl struct {
	IUnknownVtbl
	pGetFactory uintptr
}

func (this *ID2D1ResourceVtbl) GetFactory(ptr ComObjectPtr) (f ID2D1FactoryPtr) {
	var tmp uintptr
	_, _, _ = syscall.Syscall(
		this.pGetFactory,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(&tmp)),
		0)
	(&f).SetRawPtr(tmp)
	return
}

type ID2D1Resource struct {
	*ID2D1ResourceVtbl
}

type ID2D1ResourcePtr struct {
	*ID2D1Resource
}

func (this ID2D1ResourcePtr) GUID() *GUID {
	return &IID_ID2D1Resource
}

func (this ID2D1ResourcePtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.ID2D1Resource))
}

func (this *ID2D1ResourcePtr) SetRawPtr(raw uintptr) {
	this.ID2D1Resource = (*ID2D1Resource)(unsafe.Pointer(raw))
}
