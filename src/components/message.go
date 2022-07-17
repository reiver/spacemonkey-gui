package message

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var DefaultMessage = Message{
	Title:    "",
	SubTitle: "",
	Body:     nil,
}

type Message struct {
	Title    string
	SubTitle string
	Body     *fyne.Container
}

func (receiver Message) Create(title, subtitle string, body *fyne.Container) (Message, error) {
	return Message{
		Title:    title,
		SubTitle: subtitle,
		Body:     body,
	}, nil
}

func (receiver Message) CreateImageTextStyle(title, subtitle, text string, img *canvas.Image) (Message, error) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(500, 500))
	bodyText := canvas.NewText(text, green)
	body := container.NewVBox(img, bodyText)
	return receiver.Create(title, subtitle, body)
}

func (receiver Message) CreateLocalImageTextStyle(title, subtitle, text, imgPath string) (Message, error) {
	img := canvas.NewImageFromFile(imgPath)
	return receiver.CreateImageTextStyle(title, subtitle, text, img)
}

func (receiver Message) CreateURLImageTextStyle(title, subtitle, text, imgURL string) (Message, error) {
	res, err := fyne.LoadResourceFromURLString(imgURL)
	if err != nil {
		return Message{}, err
	}
	img := canvas.NewImageFromResource(res)
	return receiver.CreateImageTextStyle(title, subtitle, text, img)
}

func (receiver Message) GenerateCard() *widget.Card {
	return widget.NewCard(receiver.Title, receiver.SubTitle, receiver.Body)
}
