// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

type D2D1_FACTORY_OPTIONS struct {
	DebugLevel D2D1_DEBUG_LEVEL
}

type D2D_RECT_F struct {
	Left   float32
	Top    float32
	Right  float32
	Bottom float32
}

type D2D1_RECT_F D2D_RECT_F

type D2D1_ROUNDED_RECT struct {
	Rect    D2D1_RECT_F
	RadiusX float32
	RadiusY float32
}

type D2D1_ELLIPSE struct {
	Point   D2D1_POINT_2F
	RadiusX float32
	RadiusY float32
}

type D2D_POINT_2F struct {
	X float32
	Y float32
}

type D2D1_POINT_2F D2D_POINT_2F

type D2D_MATRIX_3X2_F struct {
	A11 float32
	A12 float32
	A21 float32
	A22 float32
	A31 float32
	A32 float32
}

type D2D1_MATRIX_3X2_F D2D_MATRIX_3X2_F

type D2D1_TRIANGLE struct {
	Point1 D2D1_POINT_2F
	Point2 D2D1_POINT_2F
	Point3 D2D1_POINT_2F
}
