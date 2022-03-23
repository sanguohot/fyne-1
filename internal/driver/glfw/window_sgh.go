package glfw

import (
	"fyne.io/fyne/v2"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func (d *gLDriver) CreateWindowSgh(title string, decorate, transparent, centered bool) fyne.Window {
	win := d.createWindowSgh(title, decorate, transparent)
	if centered {
		win.SetPadded(false)
		win.CenterOnScreen()
	}
	return win
}

func (d *gLDriver) createWindowSgh(title string, decorate, transparent bool) fyne.Window {
	var ret *window
	if title == "" {
		title = defaultTitle
	}
	runOnMain(func() {
		d.initGLFW()

		ret = &window{title: title, decorate: decorate, transparent: transparent, driver: d}
		// This queue is destroyed when the window is closed.
		ret.InitEventQueue()
		go ret.RunEventQueue()

		ret.canvas = newCanvas()
		ret.canvas.context = ret
		ret.SetIcon(ret.icon)
		d.addWindow(ret)
	})
	return ret
}

func (w *window) GetGlfwWindowSgh() *glfw.Window {
	return w.view()
}

func (w *window) GetGlfwMonitorSgh() *glfw.Monitor {
	return w.getMonitorForWindow()
}

func (w *window) CreateGlfwWindowSgh() {
	w.createLock.Do(w.create)
}
