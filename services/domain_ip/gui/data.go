package gui

import "fyne.io/fyne/v2"

// Tutorial defines the data structure for a tutorial
type ToolsPanel struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

var (
	// Tools defines the metadata for ToolsPanel
	Tools = map[string]ToolsPanel{
		"main": {"Main", "", welcomeScreen},
		"domain": {"Domain", "", whoisHistoryScreen},
	}

	// ToolsPanelIndex  defines how our tools should be laid out in the index tree
	ToolsPanelIndex = map[string][]string{
		"":            {"main", "domain", "animations", "icons", "widgets", "collections", "containers", "dialogs", "windows", "binding", "advanced"},
		"collections": {"list", "table", "tree"},
		"containers":  {"apptabs", "border", "box", "center", "doctabs", "grid", "scroll", "split"},
		"widgets":     {"accordion", "button", "card", "entry", "form", "input", "progress", "text", "toolbar"},
	}
)
