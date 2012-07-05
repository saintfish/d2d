// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	modd2d1               = syscall.NewLazyDLL("d2d1.dll")
	procD2D1CreateFactory = modd2d1.NewProc("D2D1CreateFactory")
)

type D2D1_FACTORY_TYPE uint32

const (
	// The resulting factory and derived resources may only be invoked serially.
	// Reference counts on resources are interlocked, however, resource and render
	// target state is not protected from multi-threaded access.
	D2D1_FACTORY_TYPE_SINGLE_THREADED D2D1_FACTORY_TYPE = 0
	// The resulting factory may be invoked from multiple threads. Returned resources
	// use interlocked reference counting and their state is protected.
	D2D1_FACTORY_TYPE_MULTI_THREADED D2D1_FACTORY_TYPE = 1
)

// Indicates the debug level to be outputed by the debug layer.
type D2D1_DEBUG_LEVEL uint32

const (
	D2D1_DEBUG_LEVEL_NONE        D2D1_DEBUG_LEVEL = 0
	D2D1_DEBUG_LEVEL_ERROR       D2D1_DEBUG_LEVEL = 1
	D2D1_DEBUG_LEVEL_WARNING     D2D1_DEBUG_LEVEL = 2
	D2D1_DEBUG_LEVEL_INFORMATION D2D1_DEBUG_LEVEL = 3
)

type D2D1_FACTORY_OPTIONS struct {
	DebugLevel D2D1_DEBUG_LEVEL
}

var (
	// Interface ID of ID2D1Factory "06152247-6f50-465a-9245-118bfd3b6007"
	IID_ID2D1Factory = GUID{0x06152247, 0x6f50, 0x465a, [8]byte{0x92, 0x45, 0x11, 0x8b, 0xfd, 0x3b, 0x60, 0x07}}
)

type ID2D1FactoryVtbl struct {
	IUnknownVtbl
	pReloadSystemMetrics            uintptr
	pGetDesktopDpi                  uintptr
	pCreateRectangleGeometry        uintptr
	pCreateRoundedRectangleGeometry uintptr
	pCreateEllipseGeometry          uintptr
	pCreateGeometryGroup            uintptr
	pCreateTransformedGeometry      uintptr
	pCreatePathGeometry             uintptr
	pCreateStrokeStyle              uintptr
	pCreateDrawingStateBlock        uintptr
	pCreateWicBitmapRenderTarget    uintptr
	pCreateHwndRenderTarget         uintptr
	pCreateDxgiSurfaceRenderTarget  uintptr
	pCreateDCRenderTarget           uintptr
}

func (this *ID2D1FactoryVtbl) GetDesktopDpi(ptr ComObjectPtr) (dpiX, dpiY float32) {
	_, _, _ = syscall.Syscall(
		this.pGetDesktopDpi,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(&dpiX)),
		uintptr(unsafe.Pointer(&dpiY)))
	return
}

type ID2D1Factory struct {
	*ID2D1FactoryVtbl
}

type ID2D1FactoryPtr struct {
	*ID2D1Factory
}

func (this ID2D1FactoryPtr) GUID() *GUID {
	return &IID_ID2D1Factory
}

func (this ID2D1FactoryPtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.ID2D1Factory))
}

func (this *ID2D1FactoryPtr) SetRawPtr(raw uintptr) {
	this.ID2D1Factory = (*ID2D1Factory)(unsafe.Pointer(raw))
}

// CreateFactory creates instance of factory
func D2D1CreateFactory(factoryType D2D1_FACTORY_TYPE, factoryOption *D2D1_FACTORY_OPTIONS) (f ID2D1FactoryPtr) {
	var tmp uintptr
	ret, _, _ := procD2D1CreateFactory.Call(
		uintptr(factoryType),
		uintptr(unsafe.Pointer(&IID_ID2D1Factory)),
		uintptr(unsafe.Pointer(factoryOption)),
		uintptr(unsafe.Pointer(&tmp)))
	if ret == S_OK {
		(&f).SetRawPtr(tmp)
	} else {
		panic(fmt.Sprintf("Fail to create factory: %#x", ret))
	}
	return
}
