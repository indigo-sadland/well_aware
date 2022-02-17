package screens

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/indigo-sadland/well_aware/utils"
	"io/ioutil"
	"log"
	"strings"
	"time"
)
// customDorksScreen represents the part of window with uploading custom dorks
// for domain recon or keyword search.
func customDorksScreen(w fyne.Window) fyne.CanvasObject {

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter target domain or keyword")

	fileInput := widget.NewButton("Select Dorks' File", func() {

		if input.Text == "" {
			dialog.ShowError(fmt.Errorf("enter target domain or keyword before loading dorks"), w)
			return
		}

		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {

			if err == nil && reader == nil {
				return
			}

			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			lines, err := openFile(reader)
			if err != nil {
				dialog.ShowError(err, w)
			}

			buttons := createButtons(lines, input.Text)

			a := fyne.CurrentApp()
			dorksWindow := a.NewWindow("Custom Dorks")
			dorksWindow.Resize(fyne.NewSize(600, 500))

			dorkButtons := container.NewVBox(buttons...)
			scrollContainer := container.NewVScroll(dorkButtons)

			dorksWindow.SetContent(container.NewBorder(nil, nil, nil, nil, scrollContainer))
			dorksWindow.Show()

		}, w)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		fd.Show()
	})

	main := container.NewAdaptiveGrid(2, container.NewVBox(input), container.NewVBox(fileInput))

	return main
}

func openFile(f fyne.URIReadCloser) ([]string, error) {

	var lines []string

	if f == nil {
		log.Println("Cancelled")
		return nil, nil
	}
	defer f.Close()

	body, err := ioutil.ReadFile(f.URI().Path())
	if err != nil {
		return nil, fmt.Errorf("failed to load the file: %s ", err.Error())

	}

	str := strings.Split(string(body), "\n")
	for _, s := range str {
		if s == "" {
			continue
		}
		lines = append(lines, s)
	}

	return lines, nil

}

func createButtons(lines []string, target string) []fyne.CanvasObject {

	var url string
	var err error
	var dorkButtons []fyne.CanvasObject

	for n := 0; n < len(lines); n++ {
		index := n

		b := widget.NewButton(lines[index], func() {
			url = "https://www.google.com/search?q=" + lines[index] + " site:" + target
			utils.Cmd.Args = append(utils.CmdArgs, url)

			cmd := utils.Cmd
			err = cmd.Start()
			if err != nil {
				utils.ErrorLogger.Println(err.Error())
			}

			time.Sleep(500 * time.Millisecond)
			err = cmd.Process.Kill()
			if err != nil {
				utils.ErrorLogger.Println(err.Error())
			}
		})

		dorkButtons = append(dorkButtons, b)
	}

	return dorkButtons
}
