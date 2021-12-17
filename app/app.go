package app

import (
	fyne    "fyne.io/fyne/v2"
	fyneapp "fyne.io/fyne/v2/app"
)

var (
	App fyne.App
)

func init() {
	App = fyneapp.New()
	if nil == App {
		panic("nil fyne app")
	}
}

func NewWindow(title string) fyne.Window {
	return App.NewWindow(title)
}
