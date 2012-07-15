// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

import (
	"fmt"
	"syscall"
	"unsafe"
)

// 2cd906a1-12e2-11dc-9fed-001143a055f9
var IID_ID2D1Geometry = GUID{0x2cd906a1, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1GeometryVtbl struct {
	ID2D1ResourceVtbl
	pGetBounds            uintptr
	pGetWidenedBounds     uintptr
	pStrokeContainsPoint  uintptr
	pFillContainsPoint    uintptr
	pCompareWithGeometry  uintptr
	pSimplify             uintptr
	pTessellate           uintptr
	pCombineWithGeometry  uintptr
	pOutline              uintptr
	pComputeArea          uintptr
	pComputeLength        uintptr
	pComputePointAtLength uintptr
	pWiden                uintptr
}

type ID2D1Geometry struct {
	*ID2D1GeometryVtbl
}

type ID2D1GeometryPtr struct {
	*ID2D1Geometry
}

func (this ID2D1GeometryPtr) GUID() *GUID {
	return &IID_ID2D1Geometry
}

func (this ID2D1GeometryPtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.ID2D1Geometry))
}

func (this *ID2D1GeometryPtr) SetRawPtr(raw uintptr) {
	this.ID2D1Geometry = (*ID2D1Geometry)(unsafe.Pointer(raw))
}

func (this *ID2D1GeometryVtbl) GetBounds(
	ptr ComObjectPtr,
	worldTransform *D2D1_MATRIX_3X2_F) (
	bounds D2D1_RECT_F) {
	var ret, _, _ = syscall.Syscall(
		this.pGetBounds,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(unsafe.Pointer(&bounds)))
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call GetBounds: %#x", ret))
	}
	return
}

func (this *ID2D1GeometryVtbl) GetWidenedBounds(
	ptr ComObjectPtr,
	strokeWidth float32,
	strokeStyle *ID2D1StrokeStyle,
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32) (
	bounds D2D1_RECT_F) {
	var ret, _, _ = syscall.Syscall6(
		this.pGetWidenedBounds,
		6,
		ptr.RawPtr(),
		uintptr(*(*uintptr)(unsafe.Pointer(&strokeWidth))),
		uintptr(unsafe.Pointer(strokeStyle)),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(*(*uintptr)(unsafe.Pointer(&flatteningTolerance))),
		uintptr(unsafe.Pointer(&bounds)))
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call GetWidenedBounds: %#x", ret))
	}
	return
}

func (this *ID2D1GeometryVtbl) StrokeContainsPoint(
	ptr ComObjectPtr,
	point D2D1_POINT_2F,
	strokeWidth float32,
	strokeStyle *ID2D1StrokeStyle,
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32) (contains bool) {
	var ret, _, _ = syscall.Syscall9(
		this.pStrokeContainsPoint,
		8,
		ptr.RawPtr(),
		uintptr(*(*uintptr)(unsafe.Pointer(&point.X))),
		uintptr(*(*uintptr)(unsafe.Pointer(&point.Y))),
		uintptr(*(*uintptr)(unsafe.Pointer(&strokeWidth))),
		uintptr(unsafe.Pointer(strokeStyle)),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(*(*uintptr)(unsafe.Pointer(&flatteningTolerance))),
		uintptr(unsafe.Pointer(&contains)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call StrokeContainsPoint: %#x", ret))
	}
	return
}

/*
func (this *ID2D1GeometryVtbl) FillContainsPoint(
	ptr ComObjectPtr,
	point D2D1_POINT_2F,
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32,
	contains *BOOL)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pFillContainsPoint,
		5,
		ptr.RawPtr(),
		uintptr(point),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(flatteningTolerance),
		uintptr(unsafe.Pointer(contains)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call FillContainsPoint: %#x", ret))
	}
	return
}

func (this *ID2D1GeometryVtbl) CompareWithGeometry(
	ptr ComObjectPtr,
	inputGeometry *ID2D1Geometry,
	inputGeometryTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32,
	relation *D2D1_GEOMETRY_RELATION)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCompareWithGeometry,
		5,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(inputGeometry)),
		uintptr(unsafe.Pointer(inputGeometryTransform)),
		uintptr(flatteningTolerance),
		uintptr(unsafe.Pointer(relation)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CompareWithGeometry: %#x", ret))
	}
	return
}

func (this *ID2D1GeometryVtbl) Simplify(
	ptr ComObjectPtr,
	simplificationOption D2D1_GEOMETRY_SIMPLIFICATION_OPTION,
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32,
	geometrySink *ID2D1SimplifiedGeometrySink)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pSimplify,
		5,
		ptr.RawPtr(),
		uintptr(simplificationOption),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(flatteningTolerance),
		uintptr(unsafe.Pointer(geometrySink)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call Simplify: %#x", ret))
	}
	return
}

func (this *ID2D1GeometryVtbl) Tessellate(
	ptr ComObjectPtr,
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32,
	tessellationSink *ID2D1TessellationSink)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pTessellate,
		4,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(flatteningTolerance),
		uintptr(unsafe.Pointer(tessellationSink)),
		0,
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call Tessellate: %#x", ret))
	}
	return
}

func (this *ID2D1GeometryVtbl) CombineWithGeometry(
	ptr ComObjectPtr,
	inputGeometry *ID2D1Geometry,
	combineMode D2D1_COMBINE_MODE,
	inputGeometryTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32,
	geometrySink *ID2D1SimplifiedGeometrySink)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCombineWithGeometry,
		6,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(inputGeometry)),
		uintptr(combineMode),
		uintptr(unsafe.Pointer(inputGeometryTransform)),
		uintptr(flatteningTolerance),
		uintptr(unsafe.Pointer(geometrySink)))
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CombineWithGeometry: %#x", ret))
	}
	return
}

func (this *ID2D1GeometryVtbl) Outline(
	ptr ComObjectPtr,
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32,
	geometrySink *ID2D1SimplifiedGeometrySink)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pOutline,
		4,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(flatteningTolerance),
		uintptr(unsafe.Pointer(geometrySink)),
		0,
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call Outline: %#x", ret))
	}
	return
}

func (this *ID2D1GeometryVtbl) ComputeArea(
	ptr ComObjectPtr,
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32,
	area *float32)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pComputeArea,
		4,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(flatteningTolerance),
		uintptr(unsafe.Pointer(area)),
		0,
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call ComputeArea: %#x", ret))
	}
	return
}

func (this *ID2D1GeometryVtbl) ComputeLength(
	ptr ComObjectPtr,
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32,
	length *float32)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pComputeLength,
		4,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(flatteningTolerance),
		uintptr(unsafe.Pointer(length)),
		0,
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call ComputeLength: %#x", ret))
	}
	return
}

func (this *ID2D1GeometryVtbl) ComputePointAtLength(
	ptr ComObjectPtr,
	length float32,
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32,
	point *D2D1_POINT_2F,
	unitTangentVector *D2D1_POINT_2F)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pComputePointAtLength,
		6,
		ptr.RawPtr(),
		uintptr(length),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(flatteningTolerance),
		uintptr(unsafe.Pointer(point)),
		uintptr(unsafe.Pointer(unitTangentVector)))
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call ComputePointAtLength: %#x", ret))
	}
	return
}

func (this *ID2D1GeometryVtbl) Widen(
	ptr ComObjectPtr,
	strokeWidth float32,
	strokeStyle *ID2D1StrokeStyle,
	worldTransform *D2D1_MATRIX_3X2_F,
	flatteningTolerance float32,
	geometrySink *ID2D1SimplifiedGeometrySink)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pWiden,
		6,
		ptr.RawPtr(),
		uintptr(strokeWidth),
		uintptr(unsafe.Pointer(strokeStyle)),
		uintptr(unsafe.Pointer(worldTransform)),
		uintptr(flatteningTolerance),
		uintptr(unsafe.Pointer(geometrySink)))
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call Widen: %#x", ret))
	}
	return
}

*/

// 2cd906a2-12e2-11dc-9fed-001143a055f9
var IID_ID2D1RectangleGeometry = GUID{0x2cd906a2, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1RectangleGeometryVtbl struct {
	ID2D1GeometryVtbl
	pGetRect uintptr
}

type ID2D1RectangleGeometry struct {
	*ID2D1RectangleGeometryVtbl
}

type ID2D1RectangleGeometryPtr struct {
	*ID2D1RectangleGeometry
}

func (this ID2D1RectangleGeometryPtr) GUID() *GUID {
	return &IID_ID2D1RectangleGeometry
}

func (this ID2D1RectangleGeometryPtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.ID2D1RectangleGeometry))
}

func (this *ID2D1RectangleGeometryPtr) SetRawPtr(raw uintptr) {
	this.ID2D1RectangleGeometry = (*ID2D1RectangleGeometry)(unsafe.Pointer(raw))
}

/*
func (this *ID2D1RectangleGeometryVtbl) GetRect(
	ptr ComObjectPtr,
	rect *D2D1_RECT_F) {
	var _, _, _ = syscall.Syscall(
		this.pGetRect,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(rect)),
		0)

	return
}

*/

// 2cd906a3-12e2-11dc-9fed-001143a055f9
var IID_ID2D1RoundedRectangleGeometry = GUID{0x2cd906a3, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1RoundedRectangleGeometryVtbl struct {
	ID2D1GeometryVtbl
	pGetRoundedRect uintptr
}

type ID2D1RoundedRectangleGeometry struct {
	*ID2D1RoundedRectangleGeometryVtbl
}

type ID2D1RoundedRectangleGeometryPtr struct {
	*ID2D1RoundedRectangleGeometry
}

func (this ID2D1RoundedRectangleGeometryPtr) GUID() *GUID {
	return &IID_ID2D1RoundedRectangleGeometry
}

func (this ID2D1RoundedRectangleGeometryPtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.ID2D1RoundedRectangleGeometry))
}

func (this *ID2D1RoundedRectangleGeometryPtr) SetRawPtr(raw uintptr) {
	this.ID2D1RoundedRectangleGeometry = (*ID2D1RoundedRectangleGeometry)(unsafe.Pointer(raw))
}

/*
func (this *ID2D1RoundedRectangleGeometryVtbl) GetRoundedRect(
	ptr ComObjectPtr,
	roundedRect *D2D1_ROUNDED_RECT) {
	var _, _, _ = syscall.Syscall(
		this.pGetRoundedRect,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(roundedRect)),
		0)

	return
}

*/

// 2cd906a4-12e2-11dc-9fed-001143a055f9
var IID_ID2D1EllipseGeometry = GUID{0x2cd906a4, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1EllipseGeometryVtbl struct {
	ID2D1GeometryVtbl
	pGetEllipse uintptr
}

type ID2D1EllipseGeometry struct {
	*ID2D1EllipseGeometryVtbl
}

type ID2D1EllipseGeometryPtr struct {
	*ID2D1EllipseGeometry
}

func (this ID2D1EllipseGeometryPtr) GUID() *GUID {
	return &IID_ID2D1EllipseGeometry
}

func (this ID2D1EllipseGeometryPtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.ID2D1EllipseGeometry))
}

func (this *ID2D1EllipseGeometryPtr) SetRawPtr(raw uintptr) {
	this.ID2D1EllipseGeometry = (*ID2D1EllipseGeometry)(unsafe.Pointer(raw))
}

/*
func (this *ID2D1EllipseGeometryVtbl) GetEllipse(
	ptr ComObjectPtr,
	ellipse *D2D1_ELLIPSE) {
	var _, _, _ = syscall.Syscall(
		this.pGetEllipse,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(ellipse)),
		0)

	return
}

*/

// 2cd906a6-12e2-11dc-9fed-001143a055f9
var IID_ID2D1GeometryGroup = GUID{ 0x2cd906a6, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9} }

type ID2D1GeometryGroupVtbl struct {
	ID2D1GeometryVtbl
	pGetFillMode uintptr
	pGetSourceGeometryCount uintptr
	pGetSourceGeometries uintptr
}

type ID2D1GeometryGroup struct {
	*ID2D1GeometryGroupVtbl
}

type ID2D1GeometryGroupPtr struct {
	*ID2D1GeometryGroup
}

func (this ID2D1GeometryGroupPtr) GUID() *GUID {
	return &IID_ID2D1GeometryGroup
}

func (this ID2D1GeometryGroupPtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.ID2D1GeometryGroup))
}

func (this *ID2D1GeometryGroupPtr) SetRawPtr(raw uintptr) {
	this.ID2D1GeometryGroup = (*ID2D1GeometryGroup)(unsafe.Pointer(raw))
}

/*
func (this *ID2D1GeometryGroupVtbl) GetFillMode(
	ptr ComObjectPtr)
	(D2D1_FILL_MODE) {
	var _, _, _ = syscall.Syscall(
		this.pGetFillMode,
		1,
		ptr.RawPtr(),
		0,
		0)
	
	return
}

func (this *ID2D1GeometryGroupVtbl) GetSourceGeometryCount(
	ptr ComObjectPtr)
	(uint32) {
	var _, _, _ = syscall.Syscall(
		this.pGetSourceGeometryCount,
		1,
		ptr.RawPtr(),
		0,
		0)
	
	return
}

func (this *ID2D1GeometryGroupVtbl) GetSourceGeometries(
	ptr ComObjectPtr,
	geometries **ID2D1Geometry,
	geometriesCount UINT) {
	var _, _, _ = syscall.Syscall(
		this.pGetSourceGeometries,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(geometries)),
		uintptr(geometriesCount))
	
	return
}

*/

// 2cd906bb-12e2-11dc-9fed-001143a055f9
var IID_ID2D1TransformedGeometry = GUID{ 0x2cd906bb, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9} }

type ID2D1TransformedGeometryVtbl struct {
	ID2D1GeometryVtbl
	pGetSourceGeometry uintptr
	pGetTransform uintptr
}

type ID2D1TransformedGeometry struct {
	*ID2D1TransformedGeometryVtbl
}

type ID2D1TransformedGeometryPtr struct {
	*ID2D1TransformedGeometry
}

func (this ID2D1TransformedGeometryPtr) GUID() *GUID {
	return &IID_ID2D1TransformedGeometry
}

func (this ID2D1TransformedGeometryPtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.ID2D1TransformedGeometry))
}

func (this *ID2D1TransformedGeometryPtr) SetRawPtr(raw uintptr) {
	this.ID2D1TransformedGeometry = (*ID2D1TransformedGeometry)(unsafe.Pointer(raw))
}

/*
func (this *ID2D1TransformedGeometryVtbl) GetSourceGeometry(
	ptr ComObjectPtr,
	sourceGeometry **ID2D1Geometry) {
	var _, _, _ = syscall.Syscall(
		this.pGetSourceGeometry,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(sourceGeometry)),
		0)
	
	return
}

func (this *ID2D1TransformedGeometryVtbl) GetTransform(
	ptr ComObjectPtr,
	transform *D2D1_MATRIX_3X2_F) {
	var _, _, _ = syscall.Syscall(
		this.pGetTransform,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(transform)),
		0)
	
	return
}

*/

// 2cd906a5-12e2-11dc-9fed-001143a055f9
var IID_ID2D1PathGeometry = GUID{ 0x2cd906a5, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9} }

type ID2D1PathGeometryVtbl struct {
	ID2D1GeometryVtbl
	pOpen uintptr
	pStream uintptr
	pGetSegmentCount uintptr
	pGetFigureCount uintptr
}

type ID2D1PathGeometry struct {
	*ID2D1PathGeometryVtbl
}

type ID2D1PathGeometryPtr struct {
	*ID2D1PathGeometry
}

func (this ID2D1PathGeometryPtr) GUID() *GUID {
	return &IID_ID2D1PathGeometry
}

func (this ID2D1PathGeometryPtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.ID2D1PathGeometry))
}

func (this *ID2D1PathGeometryPtr) SetRawPtr(raw uintptr) {
	this.ID2D1PathGeometry = (*ID2D1PathGeometry)(unsafe.Pointer(raw))
}

/*
func (this *ID2D1PathGeometryVtbl) Open(
	ptr ComObjectPtr,
	geometrySink **ID2D1GeometrySink)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall(
		this.pOpen,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(geometrySink)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call Open: %#x", ret))
	}
	return
}

func (this *ID2D1PathGeometryVtbl) Stream(
	ptr ComObjectPtr,
	geometrySink *ID2D1GeometrySink)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall(
		this.pStream,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(geometrySink)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call Stream: %#x", ret))
	}
	return
}

func (this *ID2D1PathGeometryVtbl) GetSegmentCount(
	ptr ComObjectPtr,
	count *uint32)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall(
		this.pGetSegmentCount,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(count)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call GetSegmentCount: %#x", ret))
	}
	return
}

func (this *ID2D1PathGeometryVtbl) GetFigureCount(
	ptr ComObjectPtr,
	count *uint32)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall(
		this.pGetFigureCount,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(count)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call GetFigureCount: %#x", ret))
	}
	return
}

*/
