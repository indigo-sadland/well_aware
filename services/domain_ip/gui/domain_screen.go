package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strings"
	"well_aware/services/domain_ip"
)

func whoisHistoryScreen(_ fyne.Window) fyne.CanvasObject  {
	var records []string

	tabs := container.NewAppTabs()
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter target domain")

	submitButton := widget.NewButton("Submit", func() {
		records = domain_ip.GetWhoisHistoryFree(input.Text)

		// Here we split the records slice into sub-slices
		// where each sub-slice is a record from whois history.
		size := 9
		var j int
		for i := 0; i < len(records); i += size {
			j += size
			if j > len(records) {
				j = len(records)
			}
			ss := strings.Join(records[i+1:j], "\n\n")

			text := widget.NewMultiLineEntry()
			text.SetText(ss)
			text.Disable()
			_ = text.TextStyle.Bold
			tabs.Append(container.NewTabItem(fmt.Sprintf("%s", records[i]), text))
		}
		tabs.SetTabLocation(container.TabLocationBottom)

	})

	topContainer := container.NewGridWithColumns(2, input, submitButton)

	return container.NewBorder(topContainer, nil, nil, nil, tabs)
}
