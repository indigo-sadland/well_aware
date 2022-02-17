package screens

import (
	"fyne.io/fyne/v2"
)

// ToolsPanel defines the data structure for a tutorial.
type ToolsPanel struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

var (

	// Tools defines the metadata for ToolsPanel.
	Tools = map[string]ToolsPanel{
		"main":         {"Main", "", welcomeScreen},
		"domain":       {"Domain Recon", "set of dorks and other tools for domain recon", domainReconScreen},
		"keyword":      {"Keyword Search", "set of dorks and other tools for proper search", keywordSearchScreen},
		"custom dorks": {"Custom Dorks", "upload your list of dorks", customDorksScreen},
	}

	// ToolsPanelIndex  defines how our tools should be laid out in the index tree.
	ToolsPanelIndex = map[string][]string{
		"":       {"main", "domain", "keyword", "custom dorks"},
	}
)
