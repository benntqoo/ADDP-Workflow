//go:build windows
// +build windows

package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"ai-launcher/internal/project"
)

type ProjectConfigDialog struct {
	window fyne.Window
}

func NewProjectConfigDialog(parent fyne.Window, pm *project.ConfigManager, onConfigured func(project.ProjectConfig, project.AIModelType)) *ProjectConfigDialog {
	return &ProjectConfigDialog{window: parent}
}

func (d *ProjectConfigDialog) Show() { dialog.ShowInformation("Info", "Project configuration dialog is temporarily unavailable on Windows build.", d.window) }
func (d *ProjectConfigDialog) Hide() { /* no-op */ }
