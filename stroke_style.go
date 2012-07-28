// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

import (
	"syscall"
	"unsafe"
)

// 2cd9069d-12e2-11dc-9fed-001143a055f9
var IID_ID2D1StrokeStyle = GUID{0x2cd9069d, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1StrokeStyleVtbl struct {
	ID2D1ResourceVtbl
	pGetStartCap    uintptr
	pGetEndCap      uintptr
	pGetDashCap     uintptr
	pGetMiterLimit  uintptr
	pGetLineJoin    uintptr
	pGetDashOffset  uintptr
	pGetDashStyle   uintptr
	pGetDashesCount uintptr
	pGetDashes      uintptr
}

type ID2D1StrokeStyle struct {
	*ID2D1StrokeStyleVtbl
}

type ID2D1StrokeStylePtr struct {
	*ID2D1StrokeStyle
}

func (this ID2D1StrokeStylePtr) GUID() *GUID {
	return &IID_ID2D1StrokeStyle
}

func (this ID2D1StrokeStylePtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.ID2D1StrokeStyle))
}

func (this *ID2D1StrokeStylePtr) SetRawPtr(raw uintptr) {
	this.ID2D1StrokeStyle = (*ID2D1StrokeStyle)(unsafe.Pointer(raw))
}

func (this *ID2D1StrokeStyleVtbl) GetStartCap(
	ptr ComObjectPtr) D2D1_CAP_STYLE {
	var ret, _, _ = syscall.Syscall(
		this.pGetStartCap,
		1,
		ptr.RawPtr(),
		0,
		0)
	return D2D1_CAP_STYLE(ret)
}

func (this *ID2D1StrokeStyleVtbl) GetEndCap(
	ptr ComObjectPtr) D2D1_CAP_STYLE {
	var ret, _, _ = syscall.Syscall(
		this.pGetEndCap,
		1,
		ptr.RawPtr(),
		0,
		0)
	return D2D1_CAP_STYLE(ret)
}

func (this *ID2D1StrokeStyleVtbl) GetDashCap(
	ptr ComObjectPtr) D2D1_CAP_STYLE {
	var ret, _, _ = syscall.Syscall(
		this.pGetDashCap,
		1,
		ptr.RawPtr(),
		0,
		0)
	return D2D1_CAP_STYLE(ret)
}

func (this *ID2D1StrokeStyleVtbl) GetMiterLimit(
	ptr ComObjectPtr) float32 {
	var ret, _, _ = syscall.Syscall(
		this.pGetMiterLimit,
		1,
		ptr.RawPtr(),
		0,
		0)
	return *(*float32)(unsafe.Pointer(&ret))
}

func (this *ID2D1StrokeStyleVtbl) GetLineJoin(
	ptr ComObjectPtr) D2D1_LINE_JOIN {
	var ret, _, _ = syscall.Syscall(
		this.pGetLineJoin,
		1,
		ptr.RawPtr(),
		0,
		0)
	return D2D1_LINE_JOIN(ret)
}

func (this *ID2D1StrokeStyleVtbl) GetDashOffset(
	ptr ComObjectPtr) float32 {
	var ret, _, _ = syscall.Syscall(
		this.pGetDashOffset,
		1,
		ptr.RawPtr(),
		0,
		0)
	return *(*float32)(unsafe.Pointer(&ret))
}

func (this *ID2D1StrokeStyleVtbl) GetDashStyle(
	ptr ComObjectPtr) D2D1_DASH_STYLE {
	var ret, _, _ = syscall.Syscall(
		this.pGetDashStyle,
		1,
		ptr.RawPtr(),
		0,
		0)
	return D2D1_DASH_STYLE(ret)
}

func (this *ID2D1StrokeStyleVtbl) GetDashesCount(
	ptr ComObjectPtr) uint32 {
	var ret, _, _ = syscall.Syscall(
		this.pGetDashesCount,
		1,
		ptr.RawPtr(),
		0,
		0)
	return uint32(ret)
}

func (this *ID2D1StrokeStyleVtbl) GetDashes(
	ptr ComObjectPtr,
	dashesCount uint32) (
	dashes []float32) {
	dashes = make([]float32, int(dashesCount))
	var _, _, _ = syscall.Syscall(
		this.pGetDashes,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(&(dashes[0]))),
		uintptr(dashesCount))
	return
}
