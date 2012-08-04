// Copyright 2012 The d2d Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/AllenDang/w32"
	"github.com/saintfish/d2d"
	"math"
	"syscall"
	"unsafe"
)

type ReleasableComObjectPtr interface {
	d2d.ComObjectPtrPtr
	Release(d2d.ComObjectPtr) uint32
}

func safeRelease(p ReleasableComObjectPtr) {
	if p.RawPtr() != 0 {
		p.Release(p)
	}
	p.SetRawPtr(0)
}


type DemoApp struct {
	hwnd w32.HWND
	factory d2d.ID2D1FactoryPtr
	render_target d2d.ID2D1HwndRenderTargetPtr
}

func (app *DemoApp) Initialize() {
	app.CreateDeviceIndependentResources()
	hInstance := w32.GetModuleHandle("")
	icon := w32.LoadIcon(0, w32.MakeIntResource(w32.IDI_APPLICATION))
	wndProc := func (hwnd w32.HWND, msg uint, wParam, lParam uintptr) uintptr {
		return app.WndProc(hwnd, msg, wParam, lParam)
	}
	wndClass := w32.WNDCLASSEX{
		Size: uint(unsafe.Sizeof(w32.WNDCLASSEX{})),
		Style: w32.CS_HREDRAW | w32.CS_VREDRAW,
		WndProc: syscall.NewCallback(wndProc),
		ClsExtra: 0,
		WndExtra: 0,
		Instance: hInstance,
		Icon: icon,
		Cursor: w32.LoadCursor(0, w32.MakeIntResource(w32.IDC_ARROW)),
		Background: 0,
		MenuName: nil,
		ClassName: syscall.StringToUTF16Ptr("D2DDemoApp"),
		IconSm: icon,
	}
	w32.RegisterClassEx(&wndClass)

	dpiX, dpiY := app.factory.GetDesktopDpi(app.factory)

	app.hwnd = w32.CreateWindowEx(
		0,
		syscall.StringToUTF16Ptr("D2DDemoApp"),
		syscall.StringToUTF16Ptr("Hello Windows"),
		w32.WS_OVERLAPPEDWINDOW,
		w32.CW_USEDEFAULT,
		w32.CW_USEDEFAULT,
		int(math.Ceil(float64(640 * dpiX / 96))),
		int(math.Ceil(float64(640 * dpiY / 96))),
		0,
		0,
		hInstance,
		nil)
	w32.ShowWindow(app.hwnd, w32.SW_SHOW)
	w32.UpdateWindow(app.hwnd)
}

func (app *DemoApp) Dispose() {
	safeRelease(&(app.factory))
	safeRelease(&(app.render_target))
}

func (app *DemoApp) RunMessageLoop() {
	msg := w32.MSG{}
	for w32.GetMessage(&msg, 0, 0, 0) > 0 {
		w32.TranslateMessage(&msg)
		w32.DispatchMessage(&msg)
	}
}

func (app *DemoApp) CreateDeviceIndependentResources() {
	app.factory = d2d.D2D1CreateFactory(d2d.D2D1_FACTORY_TYPE_SINGLE_THREADED, nil)
}

func (app *DemoApp) CreateDeviceResources() {
	if app.render_target.RawPtr() == 0 {
		var rc = w32.GetClientRect(app.hwnd)
		var hwndRenderTargetProperties = d2d.HwndRenderTargetProperties(app.hwnd)
		hwndRenderTargetProperties.PixelSize = d2d.D2D1_SIZE_U{
			uint32(rc.Right - rc.Left),
			uint32(rc.Bottom - rc.Top),
		}
		app.render_target = app.factory.CreateHwndRenderTarget(
			app.factory,
			d2d.RenderTargetProperties(),
			hwndRenderTargetProperties)
	}
}

func (app *DemoApp) WndProc(hwnd w32.HWND, msg uint, wParam, lParam uintptr) uintptr {
	return w32.DefWindowProc(hwnd, msg, wParam, lParam)
}

func main() {
	w32.CoInitialize()
	defer w32.CoUninitialize()
	app := new(DemoApp)
	defer app.Dispose()
	app.Initialize()
	app.RunMessageLoop()
}
