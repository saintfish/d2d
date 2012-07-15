// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

import (
	"fmt"
	"syscall"
	"unsafe"
)

// 06152247-6f50-465a-9245-118bfd3b6007
var IID_ID2D1Factory = GUID{0x06152247, 0x6f50, 0x465a, [8]byte{0x92, 0x45, 0x11, 0x8b, 0xfd, 0x3b, 0x60, 0x07}}

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

func (this *ID2D1FactoryVtbl) ReloadSystemMetrics(
	ptr ComObjectPtr) {
	var ret, _, _ = syscall.Syscall(
		this.pReloadSystemMetrics,
		1,
		ptr.RawPtr(),
		0,
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call ReloadSystemMetrics: %#x", ret))
	}
	return
}

func (this *ID2D1FactoryVtbl) GetDesktopDpi(
	ptr ComObjectPtr) (dpiX, dpiY float32) {
	var _, _, _ = syscall.Syscall(
		this.pGetDesktopDpi,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(&dpiX)),
		uintptr(unsafe.Pointer(&dpiY)))
	return
}

func (this *ID2D1FactoryVtbl) CreateRectangleGeometry(
	ptr ComObjectPtr,
	rectangle *D2D1_RECT_F) (
	rectangleGeometry ID2D1RectangleGeometryPtr) {
	var out uintptr
	var ret, _, _ = syscall.Syscall(
		this.pCreateRectangleGeometry,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(rectangle)),
		uintptr(unsafe.Pointer(&out)))
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateRectangleGeometry: %#x", ret))
	}
	(&rectangleGeometry).SetRawPtr(out)
	return
}

func (this *ID2D1FactoryVtbl) CreateRoundedRectangleGeometry(
	ptr ComObjectPtr,
	roundedRectangle *D2D1_ROUNDED_RECT) (
	roundedRectangleGeometry ID2D1RoundedRectangleGeometryPtr) {
	var out uintptr
	var ret, _, _ = syscall.Syscall(
		this.pCreateRoundedRectangleGeometry,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(roundedRectangle)),
		uintptr(unsafe.Pointer(&out)))
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateRoundedRectangleGeometry: %#x", ret))
	}
	(&roundedRectangleGeometry).SetRawPtr(out)
	return
}

func (this *ID2D1FactoryVtbl) CreateEllipseGeometry(
	ptr ComObjectPtr,
	ellipse *D2D1_ELLIPSE) (
	ellipseGeometry ID2D1EllipseGeometryPtr) {
	var out uintptr
	var ret, _, _ = syscall.Syscall(
		this.pCreateEllipseGeometry,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(ellipse)),
		uintptr(unsafe.Pointer(out)))
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateEllipseGeometry: %#x", ret))
	}
	(&ellipseGeometry).SetRawPtr(out)
	return
}

/*
func (this *ID2D1FactoryVtbl) CreateGeometryGroup(
	ptr ComObjectPtr,
	fillMode D2D1_FILL_MODE,
	geometries **ID2D1Geometry,
	geometriesCount UINT,
	geometryGroup **ID2D1GeometryGroup)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCreateGeometryGroup,
		5,
		ptr.RawPtr(),
		uintptr(fillMode),
		uintptr(unsafe.Pointer(geometries)),
		uintptr(geometriesCount),
		uintptr(unsafe.Pointer(geometryGroup)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateGeometryGroup: %#x", ret))
	}
	return
}

func (this *ID2D1FactoryVtbl) CreateTransformedGeometry(
	ptr ComObjectPtr,
	sourceGeometry *ID2D1Geometry,
	transform *D2D1_MATRIX_3X2_F,
	transformedGeometry **ID2D1TransformedGeometry)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCreateTransformedGeometry,
		4,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(sourceGeometry)),
		uintptr(unsafe.Pointer(transform)),
		uintptr(unsafe.Pointer(transformedGeometry)),
		0,
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateTransformedGeometry: %#x", ret))
	}
	return
}

func (this *ID2D1FactoryVtbl) CreatePathGeometry(
	ptr ComObjectPtr,
	pathGeometry **ID2D1PathGeometry)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall(
		this.pCreatePathGeometry,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(pathGeometry)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreatePathGeometry: %#x", ret))
	}
	return
}

func (this *ID2D1FactoryVtbl) CreateStrokeStyle(
	ptr ComObjectPtr,
	strokeStyleProperties *D2D1_STROKE_STYLE_PROPERTIES,
	dashes *float32,
	dashesCount UINT,
	strokeStyle **ID2D1StrokeStyle)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCreateStrokeStyle,
		5,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(strokeStyleProperties)),
		uintptr(unsafe.Pointer(dashes)),
		uintptr(dashesCount),
		uintptr(unsafe.Pointer(strokeStyle)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateStrokeStyle: %#x", ret))
	}
	return
}

func (this *ID2D1FactoryVtbl) CreateDrawingStateBlock(
	ptr ComObjectPtr,
	drawingStateDescription *D2D1_DRAWING_STATE_DESCRIPTION,
	textRenderingParams *IDWriteRenderingParams,
	drawingStateBlock **ID2D1DrawingStateBlock)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCreateDrawingStateBlock,
		4,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(drawingStateDescription)),
		uintptr(unsafe.Pointer(textRenderingParams)),
		uintptr(unsafe.Pointer(drawingStateBlock)),
		0,
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateDrawingStateBlock: %#x", ret))
	}
	return
}

func (this *ID2D1FactoryVtbl) CreateWicBitmapRenderTarget(
	ptr ComObjectPtr,
	target *IWICBitmap,
	renderTargetProperties *D2D1_RENDER_TARGET_PROPERTIES,
	renderTarget **ID2D1RenderTarget)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCreateWicBitmapRenderTarget,
		4,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(target)),
		uintptr(unsafe.Pointer(renderTargetProperties)),
		uintptr(unsafe.Pointer(renderTarget)),
		0,
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateWicBitmapRenderTarget: %#x", ret))
	}
	return
}

func (this *ID2D1FactoryVtbl) CreateHwndRenderTarget(
	ptr ComObjectPtr,
	renderTargetProperties *D2D1_RENDER_TARGET_PROPERTIES,
	hwndRenderTargetProperties *D2D1_HWND_RENDER_TARGET_PROPERTIES,
	hwndRenderTarget **ID2D1HwndRenderTarget)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCreateHwndRenderTarget,
		4,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(renderTargetProperties)),
		uintptr(unsafe.Pointer(hwndRenderTargetProperties)),
		uintptr(unsafe.Pointer(hwndRenderTarget)),
		0,
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateHwndRenderTarget: %#x", ret))
	}
	return
}

func (this *ID2D1FactoryVtbl) CreateDxgiSurfaceRenderTarget(
	ptr ComObjectPtr,
	dxgiSurface *IDXGISurface,
	renderTargetProperties *D2D1_RENDER_TARGET_PROPERTIES,
	renderTarget **ID2D1RenderTarget)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCreateDxgiSurfaceRenderTarget,
		4,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(dxgiSurface)),
		uintptr(unsafe.Pointer(renderTargetProperties)),
		uintptr(unsafe.Pointer(renderTarget)),
		0,
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateDxgiSurfaceRenderTarget: %#x", ret))
	}
	return
}

func (this *ID2D1FactoryVtbl) CreateDCRenderTarget(
	ptr ComObjectPtr,
	renderTargetProperties *D2D1_RENDER_TARGET_PROPERTIES,
	dcRenderTarget **ID2D1DCRenderTarget)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall(
		this.pCreateDCRenderTarget,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(renderTargetProperties)),
		uintptr(unsafe.Pointer(dcRenderTarget)))
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateDCRenderTarget: %#x", ret))
	}
	return
}

*/
