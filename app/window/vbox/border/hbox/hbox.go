package hbox

import (
	"github.com/reiver/spacemonkey-gui/app/window/vbox/border/hbox/prev"
	"github.com/reiver/spacemonkey-gui/app/window/vbox/border/hbox/next"

	fyne          "fyne.io/fyne/v2"
	fynecontainer "fyne.io/fyne/v2/container"
)

var (
	Container *fyne.Container
)

func init() {
	Container = fynecontainer.NewHBox(prev.Button, next.Button)
	if nil == Container {
		panic("nil fyne hbox container")
	}
}
