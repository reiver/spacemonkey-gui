package vbox

import (
	"github.com/reiver/spacemonkey-gui/app/window/vbox/border"
	"github.com/reiver/spacemonkey-gui/app/window/vbox/richtext"

	fyne          "fyne.io/fyne/v2"
	fynecontainer "fyne.io/fyne/v2/container"
)

var (
	Container *fyne.Container
)

func init() {
	Container = fynecontainer.NewVBox(border.Container, richtext.RichText)
	if nil == Container {
		panic("nil fyne vbox container")
	}
}
