package richtext

import (
	fynewidget "fyne.io/fyne/v2/widget"
)

var (
	RichText *fynewidget.RichText
)

func init() {
	RichText = fynewidget.NewRichTextFromMarkdown(
		`# SpaceMonkey!`               +"\n"+
		``                             +"\n"+
		`Hello world!`                 +"\n"+
		``                             +"\n"+
		`=> https://github.com/reiver` +"\n"+
		``,
	)
}
