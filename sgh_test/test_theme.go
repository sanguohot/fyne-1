package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	sghTheme "fyne.io/fyne/v2/sgh_test/sgh_theme"
	"fyne.io/fyne/v2/widget"
	"os"
	"path/filepath"
	"time"
)

var (
	content      = "阅读是一件非常愉快的事情，哈哈。早睡早起可以养生!o(╯□╰)o \r\nHappy Friday"
	defaultTitle = "Reader"
	app          fyne.App
)

type win struct {
	fyneWin    fyne.Window
	label      *widget.Label
	borderless bool
	x, y       int
}

func NewWin() *win {
	label := widget.NewLabel("请通过快捷键进行操作")
	return &win{
		label: label,
		x:     800,
		y:     500,
	}
}

func CreateWin(title string, borderless bool) fyne.Window {
	return createWinCore(title, borderless)
}

func createWinCore(title string, borderless bool) fyne.Window {
	w := app.Driver().(desktop.DriverSgh).CreateWindowSgh(title, !borderless, true, false)
	return w
}

func (w *win) draw() {
	nw := CreateWin(defaultTitle, w.borderless)
	nw.SetContent(w.label)
	gm := nw.GetGlfwMonitorSgh()
	if gm == nil {
		fmt.Printf("1 view window nil\r\n")
	} else {
		fmt.Printf("2 view window is not nil\r\n")
	}
	w.fyneWin = nw
	w.fyneWin.Show()
}

func (w *win) redraw() {
	fyneWin := w.fyneWin
	go func() {
		time.Sleep(time.Millisecond * 500)
		fyneWin.Close()
	}()
	w.draw()
}

func (w *win) Exit() {
	app.Quit()
	os.Exit(0)
}

func (w *win) SetBorderless(borderless bool) {
	w.borderless = borderless
	w.redraw()
}

func init() {
	dir, _ := os.Getwd()
	os.Setenv("FYNE_FONT", filepath.Join(dir, "sgh_test", "assets", "static", "Consolas-with-Yahei Regular Nerd Font.ttf"))
	os.Setenv("FYNE_SCALE", "0.8")
	app = fyneApp.New()
	app.Settings().SetTheme(sghTheme.NewSghTheme())
	filePath := filepath.Join(dir, "sgh_test", "assets", "icon", "reader-svgrepo-com.svg")
	resource, err := fyne.LoadResourceFromPath(filePath)
	if err != nil {
		panic(err)
	}
	app.SetIcon(resource)
}

func main() {
	time.Sleep(time.Second * 1)
	w := NewWin()
	w.borderless = true
	go func() {
		for {
			time.Sleep(time.Second * 3)
			w.borderless = !w.borderless
			fmt.Printf("now setting borderless %+v\r\n", w.borderless)
			w.redraw()
		}
	}()
	go func() {
		cnt := 0
		for {
			time.Sleep(time.Second * 1)
			w.label.SetText(fmt.Sprintf("%s %d", content, cnt))
			cnt++
		}
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond * 50)
			if w.fyneWin == nil {
				continue
			}
			gw := w.fyneWin.GetGlfwWindowSgh()
			if gw == nil {
				fmt.Printf("view window nil\r\n")
				continue
			}
			if gw != nil {
				x, y := gw.GetPos()
				if x != w.x && y != w.y {
					gw.SetPos(w.x, w.y)
				}
			}
		}
	}()
	// 主线程可以保证viewpoint不为nil，利于初始化位置
	w.draw()

	app.Run()
}
