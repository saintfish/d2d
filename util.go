// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

import (
	"github.com/AllenDang/w32"
)

func RenderTargetProperties() *D2D1_RENDER_TARGET_PROPERTIES {
	return &D2D1_RENDER_TARGET_PROPERTIES {
		Type : D2D1_RENDER_TARGET_TYPE_DEFAULT,
		PixelFormat : *(PixelFormat()),
		DpiX : 0,
		DpiY : 0,
		Usage : D2D1_RENDER_TARGET_USAGE_NONE,
		MinLevel : D2D1_FEATURE_LEVEL_DEFAULT,
	}
}

func PixelFormat() *D2D1_PIXEL_FORMAT {
	return &D2D1_PIXEL_FORMAT {
		Format : DXGI_FORMAT_UNKNOWN,
		AlphaMode : D2D1_ALPHA_MODE_UNKNOWN,
	}
}

func HwndRenderTargetProperties(hwnd w32.HWND) *D2D1_HWND_RENDER_TARGET_PROPERTIES {
	return &D2D1_HWND_RENDER_TARGET_PROPERTIES {
		Hwnd : hwnd,
		PixelSize : D2D1_SIZE_U{0, 0},
		PresentOptions : D2D1_PRESENT_OPTIONS_NONE,
	}
}
