package do

import (
	"github.com/reiver/spacemonkey-gui/app/window/vbox/border/entry"

	fynedatabinding "fyne.io/fyne/v2/data/binding"
	fynetheme       "fyne.io/fyne/v2/theme"
	fynewidget      "fyne.io/fyne/v2/widget"

	"fmt"
)

var (
	Button *fynewidget.Button
	Value   fynedatabinding.String
)

func init() {
	Value = fynedatabinding.NewString()
	if nil == Value{
		panic("nil fyne string data binding")
	}


	Button = fynewidget.NewButtonWithIcon("", fynetheme.ConfirmIcon(), tapped)
	if nil == Button {
		panic("nil fyne button widget")
	}
}

func tapped() {

	var entryValue string
	{
		var err error

		entryValue, err = entry.Value.Get()
		if nil != err {
			
			return
		}
	}
			
	fmt.Println("entry-value:", entryValue)

	{
		err := Value.Set(entryValue)
		if nil != err {
			
			return
		}
	}
}
