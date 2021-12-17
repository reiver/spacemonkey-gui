package border

import (
	"github.com/reiver/spacemonkey-gui/app/window/vbox/border/do"
	"github.com/reiver/spacemonkey-gui/app/window/vbox/border/entry"
	"github.com/reiver/spacemonkey-gui/app/window/vbox/border/hbox"

	fyne          "fyne.io/fyne/v2"
	fynecontainer "fyne.io/fyne/v2/container"
)

var (
	Container *fyne.Container
)

func init() {
	Container = fynecontainer.NewBorder(nil, nil, hbox.Container, do.Button, entry.Entry)
	if nil == Container {
		panic("nil fyne border container")
	}
}
