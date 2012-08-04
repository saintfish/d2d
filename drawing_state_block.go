// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

import (
	"syscall"
	"unsafe"
)

// 28506e39-ebf6-46a1-bb47-fd85565ab957
var IID_ID2D1DrawingStateBlock = GUID{0x28506e39, 0xebf6, 0x46a1, [8]byte{0xbb, 0x47, 0xfd, 0x85, 0x56, 0x5a, 0xb9, 0x57}}

type ID2D1DrawingStateBlockVtbl struct {
	ID2D1ResourceVtbl
	pGetDescription         uintptr
	pSetDescription         uintptr
	pSetTextRenderingParams uintptr
	pGetTextRenderingParams uintptr
}

type ID2D1DrawingStateBlock struct {
	*ID2D1DrawingStateBlockVtbl
}

type ID2D1DrawingStateBlockPtr struct {
	*ID2D1DrawingStateBlock
}

func (this ID2D1DrawingStateBlockPtr) GUID() *GUID {
	return &IID_ID2D1DrawingStateBlock
}

func (this ID2D1DrawingStateBlockPtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.ID2D1DrawingStateBlock))
}

func (this *ID2D1DrawingStateBlockPtr) SetRawPtr(raw uintptr) {
	this.ID2D1DrawingStateBlock = (*ID2D1DrawingStateBlock)(unsafe.Pointer(raw))
}

func (this *ID2D1DrawingStateBlockVtbl) GetDescription(
	ptr ComObjectPtr) (stateDescription D2D1_DRAWING_STATE_DESCRIPTION) {
	var _, _, _ = syscall.Syscall(
		this.pGetDescription,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(&stateDescription)),
		0)

	return
}

func (this *ID2D1DrawingStateBlockVtbl) SetDescription(
	ptr ComObjectPtr,
	stateDescription *D2D1_DRAWING_STATE_DESCRIPTION) {
	var _, _, _ = syscall.Syscall(
		this.pSetDescription,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(stateDescription)),
		0)

	return
}

func (this *ID2D1DrawingStateBlockVtbl) SetTextRenderingParams(
	ptr ComObjectPtr,
	// *IDWriteRenderingParams
	textRenderingParams unsafe.Pointer) {
	var _, _, _ = syscall.Syscall(
		this.pSetTextRenderingParams,
		2,
		ptr.RawPtr(),
		uintptr(textRenderingParams),
		0)

	return
}

func (this *ID2D1DrawingStateBlockVtbl) GetTextRenderingParams(
	ptr ComObjectPtr) (
	// *IDWriteRenderingParams
	textRenderingParams unsafe.Pointer) {
	var out uintptr
	var _, _, _ = syscall.Syscall(
		this.pGetTextRenderingParams,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(&out)),
		0)
	textRenderingParams = unsafe.Pointer(out)
	return
}
