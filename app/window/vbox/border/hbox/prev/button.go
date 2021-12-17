package prev

import (
	fynetheme  "fyne.io/fyne/v2/theme"
	fynewidget "fyne.io/fyne/v2/widget"

	"fmt"
)

var (
	Button *fynewidget.Button
)

func init() {
	Button = fynewidget.NewButtonWithIcon("", fynetheme.NavigateBackIcon(), tapped)
	if nil == Button {
		panic("nil fyne button widget")
	}
}

func tapped() {
	fmt.Println("PREV BUTTON TAPPED")
}
