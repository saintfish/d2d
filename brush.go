// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

import (
	"syscall"
	"unsafe"
)

// 2cd906a8-12e2-11dc-9fed-001143a055f9
var IID_ID2D1Brush = GUID{ 0x2cd906a8, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9} }

type ID2D1BrushVtbl struct {
	ID2D1ResourceVtbl
	pSetOpacity uintptr
	pSetTransform uintptr
	pGetOpacity uintptr
	pGetTransform uintptr
}

type ID2D1Brush struct {
	*ID2D1BrushVtbl
}

type ID2D1BrushPtr struct {
	*ID2D1Brush
}

func (this ID2D1BrushPtr) GUID() *GUID {
	return &IID_ID2D1Brush
}

func (this ID2D1BrushPtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.ID2D1Brush))
}

func (this *ID2D1BrushPtr) SetRawPtr(raw uintptr) {
	this.ID2D1Brush = (*ID2D1Brush)(unsafe.Pointer(raw))
}

func (this *ID2D1BrushVtbl) SetOpacity(
	ptr ComObjectPtr,
	opacity float32) {
	var _, _, _ = syscall.Syscall(
		this.pSetOpacity,
		2,
		ptr.RawPtr(),
		*((*uintptr)(unsafe.Pointer(&opacity))),
		0)
	
	return
}

func (this *ID2D1BrushVtbl) SetTransform(
	ptr ComObjectPtr,
	transform *D2D1_MATRIX_3X2_F) {
	var _, _, _ = syscall.Syscall(
		this.pSetTransform,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(transform)),
		0)
	
	return
}

func (this *ID2D1BrushVtbl) GetOpacity(
	ptr ComObjectPtr) float32 {
	var ret, _, _ = syscall.Syscall(
		this.pGetOpacity,
		1,
		ptr.RawPtr(),
		0,
		0)
	
	return *((*float32)(unsafe.Pointer(&ret)))
}

func (this *ID2D1BrushVtbl) GetTransform(
	ptr ComObjectPtr) (transform D2D1_MATRIX_3X2_F) {
	var _, _, _ = syscall.Syscall(
		this.pGetTransform,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(&transform)),
		0)
	
	return
}



// 2cd906a9-12e2-11dc-9fed-001143a055f9
var IID_ID2D1SolidColorBrush = GUID{ 0x2cd906a9, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9} }

type ID2D1SolidColorBrushVtbl struct {
	ID2D1BrushVtbl
	pSetColor uintptr
	pGetColor uintptr
}

type ID2D1SolidColorBrush struct {
	*ID2D1SolidColorBrushVtbl
}

type ID2D1SolidColorBrushPtr struct {
	*ID2D1SolidColorBrush
}

func (this ID2D1SolidColorBrushPtr) GUID() *GUID {
	return &IID_ID2D1SolidColorBrush
}

func (this ID2D1SolidColorBrushPtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.ID2D1SolidColorBrush))
}

func (this *ID2D1SolidColorBrushPtr) SetRawPtr(raw uintptr) {
	this.ID2D1SolidColorBrush = (*ID2D1SolidColorBrush)(unsafe.Pointer(raw))
}

func (this *ID2D1SolidColorBrushVtbl) SetColor(
	ptr ComObjectPtr,
	color *D2D1_COLOR_F) {
	var _, _, _ = syscall.Syscall(
		this.pSetColor,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(color)),
		0)
	
	return
}

func (this *ID2D1SolidColorBrushVtbl) GetColor(
	ptr ComObjectPtr) (color D2D1_COLOR_F) {
	var _, _, _ = syscall.Syscall(
		this.pGetColor,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(&color)),
		0)
	
	return
}


