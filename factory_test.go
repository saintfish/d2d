// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package d2d

import (
	"testing"
)

func TestCreateFactory(t *testing.T) {
	f, e := CreateFactory(
		FACTORY_TYPE_SINGLE_THREADED,
		&FACTORY_OPTIONS{DEBUG_LEVEL_NONE})
	if e != nil {
		t.Errorf("Error in CreateFactory: %v", e)
	}
	if f == nil {
		t.Errorf("Factory is nil")
	}
}
