package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func welcomeScreen(_ fyne.Window) fyne.CanvasObject {
	pic, _ := fyne.LoadResourceFromPath("C:\\Users\\Public\\logo.png") // TODO: Change this
	logo := canvas.NewImageFromResource(pic)
	logo.FillMode = canvas.ImageFillContain
	logo.SetMinSize(fyne.NewSize(567, 167))


	return container.NewCenter(container.NewVBox(
		logo,
		container.NewHBox(
			widget.NewHyperlink("GitHub", parseURL("https://github.com/indigo-sadland/")),
		),
	))
}
