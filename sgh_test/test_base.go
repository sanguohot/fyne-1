package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/sgh_test/sgh_theme"
	"fyne.io/fyne/v2/widget"
	"os"
	"path/filepath"
	"time"
)

func main() {
	dir, _ := os.Getwd()
	os.Setenv("FYNE_FONT", filepath.Join(dir, "sgh_test", "assets", "static", "Consolas-with-Yahei Regular Nerd Font.ttf"))
	os.Setenv("FYNE_SCALE", "0.9")
	a := fyneApp.New()
	a.Settings().SetTheme(sghTheme.NewSghTheme())
	//w := a.Driver().(desktop.Driver).CreateSplashWindow()
	w := a.Driver().(desktop.DriverSgh).CreateWindowSgh("1223", false, true, true)
	hello := widget.NewLabel("你好啊同学!")
	w.SetContent(hello)
	filePath := filepath.Join(dir, "sgh_test", "assets", "icon", "reader-svgrepo-com.svg")
	resource, err := fyne.LoadResourceFromPath(filePath)
	if err != nil {
		panic(err)
	}
	w.SetIcon(resource)
	w.SetContent(container.NewVBox(
		hello,
		//widget.NewButton("你好啊!", func() {
		//	hello.SetText("Welcome :)11213234234")
		//}),
	))
	go w.Show()
	go func() {
		time.Sleep(time.Second * 1)
		gw := w.GetGlfwWindowSgh()
		if gw != nil {
			x, y := gw.GetPos()
			fmt.Printf("x=%d, y=%d\r\n", x, y)
		}
	}()

	a.Run()
}
