package entry

import (
	fynedatabinding "fyne.io/fyne/v2/data/binding"
	fynewidget      "fyne.io/fyne/v2/widget"

	"fmt"
)

var (
	Entry *fynewidget.Entry
	Value  fynedatabinding.String
)

func init() {
	Value = fynedatabinding.NewString()
	if nil == Value {
		panic("nil fyne string data binding")
	}

	Entry = fynewidget.NewEntryWithData(Value)
	if nil == Entry {
		panic("nil fyne entry widget")
	}

       Value.AddListener(fynedatabinding.NewDataListener(dataChanged))
}

func dataChanged() {
        s, err := Value.Get()

        fmt.Printf("ENTRY STRING: %q\n", s)
        fmt.Printf("ENTRY ERROR:  %v\n", err)
}
