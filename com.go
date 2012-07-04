// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

import (
	"fmt"
	"syscall"
	"unsafe"
)

type GUID struct {
	Data1 uint
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

type ComObjectPtr interface {
	GUID() *GUID
	RawPtr() uintptr
}
type ComObjectPtrPtr interface {
	ComObjectPtr
	SetRawPtr(uintptr)
}

// HRESULT values
const (
	S_OK = 0
)

type IUnknownVtbl struct {
	pQueryInterface uintptr
	pAddRef         uintptr
	pRelease        uintptr
}

func (this *IUnknownVtbl) QueryInterface(srcPtr ComObjectPtr, destPtr ComObjectPtrPtr) {
	var dest uintptr
	ret, _, _ := syscall.Syscall(
		this.pQueryInterface,
		3,
		srcPtr.RawPtr(),
		uintptr(unsafe.Pointer(destPtr.GUID())),
		uintptr(unsafe.Pointer(&dest)))
	if ret == S_OK {
		destPtr.SetRawPtr(dest)
	} else {
		panic(fmt.Sprintf("Query interface error: %#x", ret))
	}
}

func (this *IUnknownVtbl) AddRef(ptr ComObjectPtr) uint32 {
	ret, _, _ := syscall.Syscall(this.pAddRef, 1, ptr.RawPtr(), 0, 0)
	return uint32(ret)
}

func (this *IUnknownVtbl) Release(ptr ComObjectPtr) uint32 {
	ret, _, _ := syscall.Syscall(this.pRelease, 1, ptr.RawPtr(), 0, 0)
	return uint32(ret)
}

type IUnknown struct {
	*IUnknownVtbl
}

type IUnknownPtr struct {
	*IUnknown
}

var IID_IUnknown = GUID{0x00000000, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

func (this IUnknownPtr) GUID() *GUID {
	return &IID_IUnknown
}

func (this IUnknownPtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.IUnknown))
}

func (this *IUnknownPtr) SetRawPtr(raw uintptr) {
	this.IUnknown = (*IUnknown)(unsafe.Pointer(raw))
}
