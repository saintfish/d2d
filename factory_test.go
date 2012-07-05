// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

import (
	"testing"
)

func TestCreateFactory(t *testing.T) {
	f := D2D1CreateFactory(
		D2D1_FACTORY_TYPE_SINGLE_THREADED,
		&D2D1_FACTORY_OPTIONS{D2D1_DEBUG_LEVEL_NONE})
	if f.ID2D1Factory == nil {
		t.Errorf("Factory is nil")
	}
	defer f.Release(f)
	var i IUnknownPtr
	f.QueryInterface(f, &i)
	if i.IUnknown == nil {
		t.Errorf("IUnknown is nil")
	}
	defer i.Release(i)
}

func TestGetDesktopDpi(t *testing.T) {
	f := D2D1CreateFactory(D2D1_FACTORY_TYPE_SINGLE_THREADED, nil)
	defer f.Release(f)
	x, y := f.GetDesktopDpi(f)
	if x ~= 0 || y == 0 {
		t.Errorf("Dpi is zero: %f, %f", x, y)
	}
}
