// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

import (
	"unsafe"
)

// 2cd90694-12e2-11dc-9fed-001143a055f9
var IID_ID2D1RenderTarget = GUID{0x2cd90694, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1RenderTargetVtbl struct {
	ID2D1ResourceVtbl
	pCreateBitmap                 uintptr
	pCreateBitmapFromWicBitmap    uintptr
	pCreateSharedBitmap           uintptr
	pCreateBitmapBrush            uintptr
	pCreateSolidColorBrush        uintptr
	pCreateGradientStopCollection uintptr
	pCreateLinearGradientBrush    uintptr
	pCreateRadialGradientBrush    uintptr
	pCreateCompatibleRenderTarget uintptr
	pCreateLayer                  uintptr
	pCreateMesh                   uintptr
	pDrawLine                     uintptr
	pDrawRectangle                uintptr
	pFillRectangle                uintptr
	pDrawRoundedRectangle         uintptr
	pFillRoundedRectangle         uintptr
	pDrawEllipse                  uintptr
	pFillEllipse                  uintptr
	pDrawGeometry                 uintptr
	pFillGeometry                 uintptr
	pFillMesh                     uintptr
	pFillOpacityMask              uintptr
	pDrawBitmap                   uintptr
	pDrawText                     uintptr
	pDrawTextLayout               uintptr
	pDrawGlyphRun                 uintptr
	pSetTransform                 uintptr
	pGetTransform                 uintptr
	pSetAntialiasMode             uintptr
	pGetAntialiasMode             uintptr
	pSetTextAntialiasMode         uintptr
	pGetTextAntialiasMode         uintptr
	pSetTextRenderingParams       uintptr
	pGetTextRenderingParams       uintptr
	pSetTags                      uintptr
	pGetTags                      uintptr
	pPushLayer                    uintptr
	pPopLayer                     uintptr
	pFlush                        uintptr
	pSaveDrawingState             uintptr
	pRestoreDrawingState          uintptr
	pPushAxisAlignedClip          uintptr
	pPopAxisAlignedClip           uintptr
	pClear                        uintptr
	pBeginDraw                    uintptr
	pEndDraw                      uintptr
	pGetPixelFormat               uintptr
	pSetDpi                       uintptr
	pGetDpi                       uintptr
	pGetSize                      uintptr
	pGetPixelSize                 uintptr
	pGetMaximumBitmapSize         uintptr
	pIsSupported                  uintptr
}

type ID2D1RenderTarget struct {
	*ID2D1RenderTargetVtbl
}

type ID2D1RenderTargetPtr struct {
	*ID2D1RenderTarget
}

func (this ID2D1RenderTargetPtr) GUID() *GUID {
	return &IID_ID2D1RenderTarget
}

func (this ID2D1RenderTargetPtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.ID2D1RenderTarget))
}

func (this *ID2D1RenderTargetPtr) SetRawPtr(raw uintptr) {
	this.ID2D1RenderTarget = (*ID2D1RenderTarget)(unsafe.Pointer(raw))
}

/*
func (this *ID2D1RenderTargetVtbl) CreateBitmap(
	ptr ComObjectPtr,
	size D2D1_SIZE_U,
	srcData *void,
	pitch uint32,
	bitmapProperties *D2D1_BITMAP_PROPERTIES,
	bitmap **ID2D1Bitmap)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCreateBitmap,
		6,
		ptr.RawPtr(),
		uintptr(size),
		uintptr(unsafe.Pointer(srcData)),
		uintptr(pitch),
		uintptr(unsafe.Pointer(bitmapProperties)),
		uintptr(unsafe.Pointer(bitmap)))
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateBitmap: %#x", ret))
	}
	return
}

func (this *ID2D1RenderTargetVtbl) CreateBitmapFromWicBitmap(
	ptr ComObjectPtr,
	wicBitmapSource *IWICBitmapSource,
	bitmapProperties *D2D1_BITMAP_PROPERTIES,
	bitmap **ID2D1Bitmap)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCreateBitmapFromWicBitmap,
		4,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(wicBitmapSource)),
		uintptr(unsafe.Pointer(bitmapProperties)),
		uintptr(unsafe.Pointer(bitmap)),
		0,
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateBitmapFromWicBitmap: %#x", ret))
	}
	return
}

func (this *ID2D1RenderTargetVtbl) CreateSharedBitmap(
	ptr ComObjectPtr,
	riid REFIID,
	data *void,
	bitmapProperties *D2D1_BITMAP_PROPERTIES,
	bitmap **ID2D1Bitmap)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCreateSharedBitmap,
		5,
		ptr.RawPtr(),
		uintptr(riid),
		uintptr(unsafe.Pointer(data)),
		uintptr(unsafe.Pointer(bitmapProperties)),
		uintptr(unsafe.Pointer(bitmap)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateSharedBitmap: %#x", ret))
	}
	return
}

func (this *ID2D1RenderTargetVtbl) CreateBitmapBrush(
	ptr ComObjectPtr,
	bitmap *ID2D1Bitmap,
	bitmapBrushProperties *D2D1_BITMAP_BRUSH_PROPERTIES,
	brushProperties *D2D1_BRUSH_PROPERTIES,
	bitmapBrush **ID2D1BitmapBrush)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCreateBitmapBrush,
		5,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(bitmapBrushProperties)),
		uintptr(unsafe.Pointer(brushProperties)),
		uintptr(unsafe.Pointer(bitmapBrush)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateBitmapBrush: %#x", ret))
	}
	return
}

func (this *ID2D1RenderTargetVtbl) CreateSolidColorBrush(
	ptr ComObjectPtr,
	color *D2D1_COLOR_F,
	brushProperties *D2D1_BRUSH_PROPERTIES,
	solidColorBrush **ID2D1SolidColorBrush)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCreateSolidColorBrush,
		4,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(color)),
		uintptr(unsafe.Pointer(brushProperties)),
		uintptr(unsafe.Pointer(solidColorBrush)),
		0,
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateSolidColorBrush: %#x", ret))
	}
	return
}

func (this *ID2D1RenderTargetVtbl) CreateGradientStopCollection(
	ptr ComObjectPtr,
	gradientStops *D2D1_GRADIENT_STOP,
	gradientStopsCount UINT,
	colorInterpolationGamma D2D1_GAMMA,
	extendMode D2D1_EXTEND_MODE,
	gradientStopCollection **ID2D1GradientStopCollection)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCreateGradientStopCollection,
		6,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(gradientStops)),
		uintptr(gradientStopsCount),
		uintptr(colorInterpolationGamma),
		uintptr(extendMode),
		uintptr(unsafe.Pointer(gradientStopCollection)))
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateGradientStopCollection: %#x", ret))
	}
	return
}

func (this *ID2D1RenderTargetVtbl) CreateLinearGradientBrush(
	ptr ComObjectPtr,
	linearGradientBrushProperties *D2D1_LINEAR_GRADIENT_BRUSH_PROPERTIES,
	brushProperties *D2D1_BRUSH_PROPERTIES,
	gradientStopCollection *ID2D1GradientStopCollection,
	linearGradientBrush **ID2D1LinearGradientBrush)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCreateLinearGradientBrush,
		5,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(linearGradientBrushProperties)),
		uintptr(unsafe.Pointer(brushProperties)),
		uintptr(unsafe.Pointer(gradientStopCollection)),
		uintptr(unsafe.Pointer(linearGradientBrush)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateLinearGradientBrush: %#x", ret))
	}
	return
}

func (this *ID2D1RenderTargetVtbl) CreateRadialGradientBrush(
	ptr ComObjectPtr,
	radialGradientBrushProperties *D2D1_RADIAL_GRADIENT_BRUSH_PROPERTIES,
	brushProperties *D2D1_BRUSH_PROPERTIES,
	gradientStopCollection *ID2D1GradientStopCollection,
	radialGradientBrush **ID2D1RadialGradientBrush)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCreateRadialGradientBrush,
		5,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(radialGradientBrushProperties)),
		uintptr(unsafe.Pointer(brushProperties)),
		uintptr(unsafe.Pointer(gradientStopCollection)),
		uintptr(unsafe.Pointer(radialGradientBrush)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateRadialGradientBrush: %#x", ret))
	}
	return
}

func (this *ID2D1RenderTargetVtbl) CreateCompatibleRenderTarget(
	ptr ComObjectPtr,
	desiredSize *D2D1_SIZE_F,
	desiredPixelSize *D2D1_SIZE_U,
	desiredFormat *D2D1_PIXEL_FORMAT,
	options D2D1_COMPATIBLE_RENDER_TARGET_OPTIONS,
	bitmapRenderTarget **ID2D1BitmapRenderTarget)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall6(
		this.pCreateCompatibleRenderTarget,
		6,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(desiredSize)),
		uintptr(unsafe.Pointer(desiredPixelSize)),
		uintptr(unsafe.Pointer(desiredFormat)),
		uintptr(options),
		uintptr(unsafe.Pointer(bitmapRenderTarget)))
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateCompatibleRenderTarget: %#x", ret))
	}
	return
}

func (this *ID2D1RenderTargetVtbl) CreateLayer(
	ptr ComObjectPtr,
	size *D2D1_SIZE_F,
	layer **ID2D1Layer)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall(
		this.pCreateLayer,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(size)),
		uintptr(unsafe.Pointer(layer)))
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateLayer: %#x", ret))
	}
	return
}

func (this *ID2D1RenderTargetVtbl) CreateMesh(
	ptr ComObjectPtr,
	mesh **ID2D1Mesh)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall(
		this.pCreateMesh,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(mesh)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call CreateMesh: %#x", ret))
	}
	return
}

func (this *ID2D1RenderTargetVtbl) DrawLine(
	ptr ComObjectPtr,
	point0 D2D1_POINT_2F,
	point1 D2D1_POINT_2F,
	brush *ID2D1Brush,
	strokeWidth float32,
	strokeStyle *ID2D1StrokeStyle) {
	var _, _, _ = syscall.Syscall6(
		this.pDrawLine,
		6,
		ptr.RawPtr(),
		uintptr(point0),
		uintptr(point1),
		uintptr(unsafe.Pointer(brush)),
		uintptr(strokeWidth),
		uintptr(unsafe.Pointer(strokeStyle)))

	return
}

func (this *ID2D1RenderTargetVtbl) DrawRectangle(
	ptr ComObjectPtr,
	rect *D2D1_RECT_F,
	brush *ID2D1Brush,
	strokeWidth float32,
	strokeStyle *ID2D1StrokeStyle) {
	var _, _, _ = syscall.Syscall6(
		this.pDrawRectangle,
		5,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(strokeWidth),
		uintptr(unsafe.Pointer(strokeStyle)),
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) FillRectangle(
	ptr ComObjectPtr,
	rect *D2D1_RECT_F,
	brush *ID2D1Brush) {
	var _, _, _ = syscall.Syscall(
		this.pFillRectangle,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(brush)))

	return
}

func (this *ID2D1RenderTargetVtbl) DrawRoundedRectangle(
	ptr ComObjectPtr,
	roundedRect *D2D1_ROUNDED_RECT,
	brush *ID2D1Brush,
	strokeWidth float32,
	strokeStyle *ID2D1StrokeStyle) {
	var _, _, _ = syscall.Syscall6(
		this.pDrawRoundedRectangle,
		5,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(roundedRect)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(strokeWidth),
		uintptr(unsafe.Pointer(strokeStyle)),
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) FillRoundedRectangle(
	ptr ComObjectPtr,
	roundedRect *D2D1_ROUNDED_RECT,
	brush *ID2D1Brush) {
	var _, _, _ = syscall.Syscall(
		this.pFillRoundedRectangle,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(roundedRect)),
		uintptr(unsafe.Pointer(brush)))

	return
}

func (this *ID2D1RenderTargetVtbl) DrawEllipse(
	ptr ComObjectPtr,
	ellipse *D2D1_ELLIPSE,
	brush *ID2D1Brush,
	strokeWidth float32,
	strokeStyle *ID2D1StrokeStyle) {
	var _, _, _ = syscall.Syscall6(
		this.pDrawEllipse,
		5,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(ellipse)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(strokeWidth),
		uintptr(unsafe.Pointer(strokeStyle)),
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) FillEllipse(
	ptr ComObjectPtr,
	ellipse *D2D1_ELLIPSE,
	brush *ID2D1Brush) {
	var _, _, _ = syscall.Syscall(
		this.pFillEllipse,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(ellipse)),
		uintptr(unsafe.Pointer(brush)))

	return
}

func (this *ID2D1RenderTargetVtbl) DrawGeometry(
	ptr ComObjectPtr,
	geometry *ID2D1Geometry,
	brush *ID2D1Brush,
	strokeWidth float32,
	strokeStyle *ID2D1StrokeStyle) {
	var _, _, _ = syscall.Syscall6(
		this.pDrawGeometry,
		5,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(geometry)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(strokeWidth),
		uintptr(unsafe.Pointer(strokeStyle)),
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) FillGeometry(
	ptr ComObjectPtr,
	geometry *ID2D1Geometry,
	brush *ID2D1Brush,
	opacityBrush *ID2D1Brush) {
	var _, _, _ = syscall.Syscall6(
		this.pFillGeometry,
		4,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(geometry)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(unsafe.Pointer(opacityBrush)),
		0,
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) FillMesh(
	ptr ComObjectPtr,
	mesh *ID2D1Mesh,
	brush *ID2D1Brush) {
	var _, _, _ = syscall.Syscall(
		this.pFillMesh,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(mesh)),
		uintptr(unsafe.Pointer(brush)))

	return
}

func (this *ID2D1RenderTargetVtbl) FillOpacityMask(
	ptr ComObjectPtr,
	opacityMask *ID2D1Bitmap,
	brush *ID2D1Brush,
	content D2D1_OPACITY_MASK_CONTENT,
	destinationRectangle *D2D1_RECT_F,
	sourceRectangle *D2D1_RECT_F) {
	var _, _, _ = syscall.Syscall6(
		this.pFillOpacityMask,
		6,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(opacityMask)),
		uintptr(unsafe.Pointer(brush)),
		uintptr(content),
		uintptr(unsafe.Pointer(destinationRectangle)),
		uintptr(unsafe.Pointer(sourceRectangle)))

	return
}

func (this *ID2D1RenderTargetVtbl) DrawBitmap(
	ptr ComObjectPtr,
	bitmap *ID2D1Bitmap,
	destinationRectangle *D2D1_RECT_F,
	opacity float32,
	interpolationMode D2D1_BITMAP_INTERPOLATION_MODE,
	sourceRectangle *D2D1_RECT_F) {
	var _, _, _ = syscall.Syscall6(
		this.pDrawBitmap,
		6,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(bitmap)),
		uintptr(unsafe.Pointer(destinationRectangle)),
		uintptr(opacity),
		uintptr(interpolationMode),
		uintptr(unsafe.Pointer(sourceRectangle)))

	return
}

func (this *ID2D1RenderTargetVtbl) DrawText(
	ptr ComObjectPtr,
	string *WCHAR,
	stringLength UINT,
	textFormat *IDWriteTextFormat,
	layoutRect *D2D1_RECT_F,
	defaultForegroundBrush *ID2D1Brush,
	options D2D1_DRAW_TEXT_OPTIONS,
	measuringMode DWRITE_MEASURING_MODE) {
	var _, _, _ = syscall.Syscall9(
		this.pDrawText,
		8,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(string)),
		uintptr(stringLength),
		uintptr(unsafe.Pointer(textFormat)),
		uintptr(unsafe.Pointer(layoutRect)),
		uintptr(unsafe.Pointer(defaultForegroundBrush)),
		uintptr(options),
		uintptr(measuringMode),
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) DrawTextLayout(
	ptr ComObjectPtr,
	origin D2D1_POINT_2F,
	textLayout *IDWriteTextLayout,
	defaultForegroundBrush *ID2D1Brush,
	options D2D1_DRAW_TEXT_OPTIONS) {
	var _, _, _ = syscall.Syscall6(
		this.pDrawTextLayout,
		5,
		ptr.RawPtr(),
		uintptr(origin),
		uintptr(unsafe.Pointer(textLayout)),
		uintptr(unsafe.Pointer(defaultForegroundBrush)),
		uintptr(options),
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) DrawGlyphRun(
	ptr ComObjectPtr,
	baselineOrigin D2D1_POINT_2F,
	glyphRun *DWRITE_GLYPH_RUN,
	foregroundBrush *ID2D1Brush,
	measuringMode DWRITE_MEASURING_MODE) {
	var _, _, _ = syscall.Syscall6(
		this.pDrawGlyphRun,
		5,
		ptr.RawPtr(),
		uintptr(baselineOrigin),
		uintptr(unsafe.Pointer(glyphRun)),
		uintptr(unsafe.Pointer(foregroundBrush)),
		uintptr(measuringMode),
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) SetTransform(
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

func (this *ID2D1RenderTargetVtbl) GetTransform(
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

func (this *ID2D1RenderTargetVtbl) SetAntialiasMode(
	ptr ComObjectPtr,
	antialiasMode D2D1_ANTIALIAS_MODE) {
	var _, _, _ = syscall.Syscall(
		this.pSetAntialiasMode,
		2,
		ptr.RawPtr(),
		uintptr(antialiasMode),
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) GetAntialiasMode(
	ptr ComObjectPtr)
	(D2D1_ANTIALIAS_MODE) {
	var _, _, _ = syscall.Syscall(
		this.pGetAntialiasMode,
		1,
		ptr.RawPtr(),
		0,
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) SetTextAntialiasMode(
	ptr ComObjectPtr,
	textAntialiasMode D2D1_TEXT_ANTIALIAS_MODE) {
	var _, _, _ = syscall.Syscall(
		this.pSetTextAntialiasMode,
		2,
		ptr.RawPtr(),
		uintptr(textAntialiasMode),
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) GetTextAntialiasMode(
	ptr ComObjectPtr)
	(D2D1_TEXT_ANTIALIAS_MODE) {
	var _, _, _ = syscall.Syscall(
		this.pGetTextAntialiasMode,
		1,
		ptr.RawPtr(),
		0,
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) SetTextRenderingParams(
	ptr ComObjectPtr,
	textRenderingParams *IDWriteRenderingParams) {
	var _, _, _ = syscall.Syscall(
		this.pSetTextRenderingParams,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(textRenderingParams)),
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) GetTextRenderingParams(
	ptr ComObjectPtr,
	textRenderingParams **IDWriteRenderingParams) {
	var _, _, _ = syscall.Syscall(
		this.pGetTextRenderingParams,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(textRenderingParams)),
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) SetTags(
	ptr ComObjectPtr,
	tag1 D2D1_TAG,
	tag2 D2D1_TAG) {
	var _, _, _ = syscall.Syscall(
		this.pSetTags,
		3,
		ptr.RawPtr(),
		uintptr(tag1),
		uintptr(tag2))

	return
}

func (this *ID2D1RenderTargetVtbl) GetTags(
	ptr ComObjectPtr,
	tag1 *D2D1_TAG,
	tag2 *D2D1_TAG) {
	var _, _, _ = syscall.Syscall(
		this.pGetTags,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(tag1)),
		uintptr(unsafe.Pointer(tag2)))

	return
}

func (this *ID2D1RenderTargetVtbl) PushLayer(
	ptr ComObjectPtr,
	layerParameters *D2D1_LAYER_PARAMETERS,
	layer *ID2D1Layer) {
	var _, _, _ = syscall.Syscall(
		this.pPushLayer,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(layerParameters)),
		uintptr(unsafe.Pointer(layer)))

	return
}

func (this *ID2D1RenderTargetVtbl) PopLayer(
	ptr ComObjectPtr) {
	var _, _, _ = syscall.Syscall(
		this.pPopLayer,
		1,
		ptr.RawPtr(),
		0,
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) Flush(
	ptr ComObjectPtr,
	tag1 *D2D1_TAG,
	tag2 *D2D1_TAG)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall(
		this.pFlush,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(tag1)),
		uintptr(unsafe.Pointer(tag2)))
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call Flush: %#x", ret))
	}
	return
}

func (this *ID2D1RenderTargetVtbl) SaveDrawingState(
	ptr ComObjectPtr,
	drawingStateBlock *ID2D1DrawingStateBlock) {
	var _, _, _ = syscall.Syscall(
		this.pSaveDrawingState,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(drawingStateBlock)),
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) RestoreDrawingState(
	ptr ComObjectPtr,
	drawingStateBlock *ID2D1DrawingStateBlock) {
	var _, _, _ = syscall.Syscall(
		this.pRestoreDrawingState,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(drawingStateBlock)),
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) PushAxisAlignedClip(
	ptr ComObjectPtr,
	clipRect *D2D1_RECT_F,
	antialiasMode D2D1_ANTIALIAS_MODE) {
	var _, _, _ = syscall.Syscall(
		this.pPushAxisAlignedClip,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(clipRect)),
		uintptr(antialiasMode))

	return
}

func (this *ID2D1RenderTargetVtbl) PopAxisAlignedClip(
	ptr ComObjectPtr) {
	var _, _, _ = syscall.Syscall(
		this.pPopAxisAlignedClip,
		1,
		ptr.RawPtr(),
		0,
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) Clear(
	ptr ComObjectPtr,
	clearColor *D2D1_COLOR_F) {
	var _, _, _ = syscall.Syscall(
		this.pClear,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(clearColor)),
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) BeginDraw(
	ptr ComObjectPtr) {
	var _, _, _ = syscall.Syscall(
		this.pBeginDraw,
		1,
		ptr.RawPtr(),
		0,
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) EndDraw(
	ptr ComObjectPtr,
	tag1 *D2D1_TAG,
	tag2 *D2D1_TAG)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall(
		this.pEndDraw,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(tag1)),
		uintptr(unsafe.Pointer(tag2)))
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call EndDraw: %#x", ret))
	}
	return
}

func (this *ID2D1RenderTargetVtbl) GetPixelFormat(
	ptr ComObjectPtr)
	(D2D1_PIXEL_FORMAT) {
	var _, _, _ = syscall.Syscall(
		this.pGetPixelFormat,
		1,
		ptr.RawPtr(),
		0,
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) SetDpi(
	ptr ComObjectPtr,
	dpiX float32,
	dpiY float32) {
	var _, _, _ = syscall.Syscall(
		this.pSetDpi,
		3,
		ptr.RawPtr(),
		uintptr(dpiX),
		uintptr(dpiY))

	return
}

func (this *ID2D1RenderTargetVtbl) GetDpi(
	ptr ComObjectPtr,
	dpiX *float32,
	dpiY *float32) {
	var _, _, _ = syscall.Syscall(
		this.pGetDpi,
		3,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(dpiX)),
		uintptr(unsafe.Pointer(dpiY)))

	return
}

func (this *ID2D1RenderTargetVtbl) GetSize(
	ptr ComObjectPtr)
	(D2D1_SIZE_F) {
	var _, _, _ = syscall.Syscall(
		this.pGetSize,
		1,
		ptr.RawPtr(),
		0,
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) GetPixelSize(
	ptr ComObjectPtr)
	(D2D1_SIZE_U) {
	var _, _, _ = syscall.Syscall(
		this.pGetPixelSize,
		1,
		ptr.RawPtr(),
		0,
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) GetMaximumBitmapSize(
	ptr ComObjectPtr)
	(uint32) {
	var _, _, _ = syscall.Syscall(
		this.pGetMaximumBitmapSize,
		1,
		ptr.RawPtr(),
		0,
		0)

	return
}

func (this *ID2D1RenderTargetVtbl) IsSupported(
	ptr ComObjectPtr,
	renderTargetProperties *D2D1_RENDER_TARGET_PROPERTIES)
	(BOOL) {
	var _, _, _ = syscall.Syscall(
		this.pIsSupported,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(renderTargetProperties)),
		0)

	return
}

*/

// 2cd90698-12e2-11dc-9fed-001143a055f9
var IID_ID2D1HwndRenderTarget = GUID{0x2cd90698, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

type ID2D1HwndRenderTargetVtbl struct {
	ID2D1RenderTargetVtbl
	pCheckWindowState uintptr
	pResize           uintptr
	pGetHwnd          uintptr
}

type ID2D1HwndRenderTarget struct {
	*ID2D1HwndRenderTargetVtbl
}

type ID2D1HwndRenderTargetPtr struct {
	*ID2D1HwndRenderTarget
}

func (this ID2D1HwndRenderTargetPtr) GUID() *GUID {
	return &IID_ID2D1HwndRenderTarget
}

func (this ID2D1HwndRenderTargetPtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.ID2D1HwndRenderTarget))
}

func (this *ID2D1HwndRenderTargetPtr) SetRawPtr(raw uintptr) {
	this.ID2D1HwndRenderTarget = (*ID2D1HwndRenderTarget)(unsafe.Pointer(raw))
}

/*
func (this *ID2D1HwndRenderTargetVtbl) CheckWindowState(
	ptr ComObjectPtr)
	(D2D1_WINDOW_STATE) {
	var _, _, _ = syscall.Syscall(
		this.pCheckWindowState,
		1,
		ptr.RawPtr(),
		0,
		0)

	return
}

func (this *ID2D1HwndRenderTargetVtbl) Resize(
	ptr ComObjectPtr,
	pixelSize *D2D1_SIZE_U)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall(
		this.pResize,
		2,
		ptr.RawPtr(),
		uintptr(unsafe.Pointer(pixelSize)),
		0)
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call Resize: %#x", ret))
	}
	return
}

func (this *ID2D1HwndRenderTargetVtbl) GetHwnd(
	ptr ComObjectPtr)
	(HWND) {
	var _, _, _ = syscall.Syscall(
		this.pGetHwnd,
		1,
		ptr.RawPtr(),
		0,
		0)

	return
}

*/

// 1c51bc64-de61-46fd-9899-63a5d8f03950
var IID_ID2D1DCRenderTarget = GUID{0x1c51bc64, 0xde61, 0x46fd, [8]byte{0x98, 0x99, 0x63, 0xa5, 0xd8, 0xf0, 0x39, 0x50}}

type ID2D1DCRenderTargetVtbl struct {
	ID2D1RenderTargetVtbl
	pBindDC uintptr
}

type ID2D1DCRenderTarget struct {
	*ID2D1DCRenderTargetVtbl
}

type ID2D1DCRenderTargetPtr struct {
	*ID2D1DCRenderTarget
}

func (this ID2D1DCRenderTargetPtr) GUID() *GUID {
	return &IID_ID2D1DCRenderTarget
}

func (this ID2D1DCRenderTargetPtr) RawPtr() uintptr {
	return uintptr(unsafe.Pointer(this.ID2D1DCRenderTarget))
}

func (this *ID2D1DCRenderTargetPtr) SetRawPtr(raw uintptr) {
	this.ID2D1DCRenderTarget = (*ID2D1DCRenderTarget)(unsafe.Pointer(raw))
}

/*
func (this *ID2D1DCRenderTargetVtbl) BindDC(
	ptr ComObjectPtr,
	hDC HDC,
	pSubRect *RECT)
	(HRESULT) {
	var ret, _, _ = syscall.Syscall(
		this.pBindDC,
		3,
		ptr.RawPtr(),
		uintptr(hDC),
		uintptr(unsafe.Pointer(pSubRect)))
	if ret != S_OK {
		panic(fmt.Sprintf("Fail to call BindDC: %#x", ret))
	}
	return
}

*/
