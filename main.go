package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	message "github.com/reiver/spacemonkey-gui/src/components"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")

	message1, _ := message.DefaultMessage.CreateURLImageTextStyle("Image from the web", "This image is loaded from the web", "Google Search", "https://cdn.shopify.com/s/files/1/0507/4354/1965/products/loch-fyne-sunset-or-paintings-and-art-prints-of-scotland-or-by-scottish-artist-kevin-hunter-or-framed-canvas-1.jpg?v=1655920015")
	card1 := message1.GenerateCard()

	message2, _ := message.DefaultMessage.CreateLocalImageTextStyle("Image from local storage", "This image is loaded from local storage", "./twitter.jpg", "twitter.jpg")
	card2 := message2.GenerateCard()

	content := container.New(layout.NewGridLayout(1), card1, card2)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
