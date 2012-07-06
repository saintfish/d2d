// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

var (
	// "2cd90694-12e2-11dc-9fed-001143a055f9"
	IID_ID2D1RenderTarget = GUID{0x2cd90694, 0x12e2, 0x11dc, [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}
)

type ID2D1RenderTargetVtbl struct {
	ID2D1ResourceVtbl
	pCreateBitmap
	pCreateBitmapFromWicBitmap
	pCreateSharedBitmap
	pCreateBitmapBrush
	pCreateSolidColorBrush
	pCreateGradientStopCollection
	pCreateLinearGradientBrush
	pCreateRadialGradientBrush
	pCreateCompatibleRenderTarget
	pCreateLayer
	pCreateMesh
	pDrawLine
	pDrawRectangle
	pFillRectangle
	pDrawRoundedRectangle
	pFillRoundedRectangle
	pDrawEllipse
	pFillEllipse
	pDrawGeometry
	pFillGeometry
	pFillMesh
	pFillOpacityMask
	pDrawBitmap
	pDrawText
	pDrawTextLayout
	pDrawGlyphRun
	pSetTransform
	pGetTransform
	pSetAntialiasMode
	pGetAntialiasMode
	pSetTextAntialiasMode
	pGetTextAntialiasMode
	pSetTextRenderingParams
	pGetTextRenderingParams
	pSetTags
	pGetTags
	pPushLayer
	pPopLayer
	pFlush
	pSaveDrawingState
	pRestoreDrawingState
	pPushAxisAlignedClip
	pPopAxisAlignedClip
	pClear
	pBeginDraw
	pEndDraw
	pGetPixelFormat
	pSetDpi
	pGetDpi
	pGetSize
	pGetPixelSize
	pGetMaximumBitmapSize
	pIsSupported
}
