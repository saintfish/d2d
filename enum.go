// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

type D2D1_FACTORY_TYPE uint32

const (
	D2D1_FACTORY_TYPE_SINGLE_THREADED D2D1_FACTORY_TYPE = 0
	D2D1_FACTORY_TYPE_MULTI_THREADED  D2D1_FACTORY_TYPE = 1
)

type D2D1_DEBUG_LEVEL uint32

const (
	D2D1_DEBUG_LEVEL_NONE        D2D1_DEBUG_LEVEL = 0
	D2D1_DEBUG_LEVEL_ERROR       D2D1_DEBUG_LEVEL = 1
	D2D1_DEBUG_LEVEL_WARNING     D2D1_DEBUG_LEVEL = 2
	D2D1_DEBUG_LEVEL_INFORMATION D2D1_DEBUG_LEVEL = 3
)

type D2D1_FILL_MODE uint32

const (
	D2D1_FILL_MODE_ALTERNATE D2D1_FILL_MODE = 0
	D2D1_FILL_MODE_WINDING   D2D1_FILL_MODE = 1
)

type D2D1_CAP_STYLE uint32

const (
	D2D1_CAP_STYLE_FLAT     D2D1_CAP_STYLE = 0
	D2D1_CAP_STYLE_SQUARE   D2D1_CAP_STYLE = 1
	D2D1_CAP_STYLE_ROUND    D2D1_CAP_STYLE = 2
	D2D1_CAP_STYLE_TRIANGLE D2D1_CAP_STYLE = 3
)

type D2D1_LINE_JOIN uint32

const (
	D2D1_LINE_JOIN_MITER          D2D1_LINE_JOIN = 0
	D2D1_LINE_JOIN_BEVEL          D2D1_LINE_JOIN = 1
	D2D1_LINE_JOIN_ROUND          D2D1_LINE_JOIN = 2
	D2D1_LINE_JOIN_MITER_OR_BEVEL D2D1_LINE_JOIN = 3
)

type D2D1_DASH_STYLE uint32

const (
	D2D1_DASH_STYLE_SOLID        D2D1_DASH_STYLE = 0
	D2D1_DASH_STYLE_DASH         D2D1_DASH_STYLE = 1
	D2D1_DASH_STYLE_DOT          D2D1_DASH_STYLE = 2
	D2D1_DASH_STYLE_DASH_DOT     D2D1_DASH_STYLE = 3
	D2D1_DASH_STYLE_DASH_DOT_DOT D2D1_DASH_STYLE = 4
	D2D1_DASH_STYLE_CUSTOM       D2D1_DASH_STYLE = 5
)

type D2D1_GEOMETRY_RELATION uint32

const (
	D2D1_GEOMETRY_RELATION_UNKNOWN      D2D1_GEOMETRY_RELATION = 0
	D2D1_GEOMETRY_RELATION_DISJOINT     D2D1_GEOMETRY_RELATION = 1
	D2D1_GEOMETRY_RELATION_IS_CONTAINED D2D1_GEOMETRY_RELATION = 2
	D2D1_GEOMETRY_RELATION_CONTAINS     D2D1_GEOMETRY_RELATION = 3
	D2D1_GEOMETRY_RELATION_OVERLAP      D2D1_GEOMETRY_RELATION = 4
)

type D2D1_GEOMETRY_SIMPLIFICATION_OPTION uint32

const (
	D2D1_GEOMETRY_SIMPLIFICATION_OPTION_CUBICS_AND_LINES D2D1_GEOMETRY_SIMPLIFICATION_OPTION = 0
	D2D1_GEOMETRY_SIMPLIFICATION_OPTION_LINES            D2D1_GEOMETRY_SIMPLIFICATION_OPTION = 1
)

type D2D1_COMBINE_MODE uint32

const (
	D2D1_COMBINE_MODE_UNION     D2D1_COMBINE_MODE = 0
	D2D1_COMBINE_MODE_INTERSECT D2D1_COMBINE_MODE = 1
	D2D1_COMBINE_MODE_XOR       D2D1_COMBINE_MODE = 2
	D2D1_COMBINE_MODE_EXCLUDE   D2D1_COMBINE_MODE = 3
)
