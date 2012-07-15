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
