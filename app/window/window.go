package window

import (
	"github.com/reiver/spacemonkey-gui/app"
	"github.com/reiver/spacemonkey-gui/app/window/vbox"

	fyne "fyne.io/fyne/v2"
)

const windowtitle = "SpaceMonkey"+" "+"⚡"

var (
	Window fyne.Window
)

func init() {
	Window = app.NewWindow(windowtitle)
	if nil == Window {
		panic("nil fyne window")
	}

	Window.SetContent(vbox.Container)
}

func ShowAndRun() {
	Window.ShowAndRun()
}
