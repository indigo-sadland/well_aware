package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/indigo-sadland/well_aware/screens"
	"github.com/indigo-sadland/well_aware/utils"
)

const preferenceCurrentTool = ""

func init() {

	utils.OsDetection()
	utils.ErrorLogInit()

}

func main() {

	newApp := app.NewWithID("Well Aware")
	newApp.Settings().SetTheme(theme.LightTheme())

	window := newApp.NewWindow("Well Aware")
	title := widget.NewLabel("")
	intro := widget.NewLabel("")
	window.Resize(fyne.NewSize(700, 500))

	content := container.NewMax()
	tool := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)
	setTutorial := func(t screens.ToolsPanel) {
		title.SetText(t.Title)
		intro.SetText(t.Intro)

		content.Objects = []fyne.CanvasObject{t.View(window)}
		content.Refresh()
	}
	split := container.NewHSplit(makeNav(setTutorial, true), tool)
	split.Offset = 0.2
	window.SetContent(split)
	window.ShowAndRun()
}

// makeNav creates tree view for the left panel
func makeNav(setTutorial func(tutorial screens.ToolsPanel), loadPrevious bool) fyne.CanvasObject {

	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return screens.ToolsPanelIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := screens.ToolsPanelIndex[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := screens.Tools[uid]
			if !ok {
				fyne.LogError("Missing tutorial panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
		},
		OnSelected: func(uid string) {
			if t, ok := screens.Tools[uid]; ok {
				a.Preferences().SetString(preferenceCurrentTool, uid)
				setTutorial(t)
			}
		},
	}

	if loadPrevious {
		//currentPref := a.Preferences().StringWithFallback(preferenceCurrentTool, "main")
		currentPref := "main"
		tree.Select(currentPref)
	}

	return container.NewBorder(nil, nil, nil, nil, tree)
}
