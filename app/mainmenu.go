package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

const (
	// Opens a json definition file
	FileItemOpen string = "Open"
	// Saves the current state to a json definition file
	FileItemSave = "Save"
	// Generates source code from the current  state
	FileItemGen = "Generate Code"
	// Clear the grid and start fresh with verification if state changed
	EditItemClear = "Clear"
	// refresh the preview screen
	EditItemPrev = "Refresh Preview"
	// Copy the preview screen to the system clipboard
	EditItemCopy = "Copy to clipboard"
	// brief description of the app
	HelpItemAbout = "About"
	// view the readme file
	HelpItemHelp = "Help"
)

func constString(c string) string {
	switch c {
	case FileItemOpen:
		return "FileItemOpen"
	case FileItemSave:
		return "FileItemSave"
	case FileItemGen:
		return "FileItemGen"
	case EditItemClear:
		return "EditItemClear"
	case EditItemPrev:
		return "EditItemPrev"
	case EditItemCopy:
		return "EditItemCopy"
	case HelpItemAbout:
		return "HelpItemAbout"
	case HelpItemHelp:
		return "HelpItemHelp"
	}
	return ""
}

func assocString(c string) string {
	switch c {
	case FileItemOpen:
		return "Opens a json definition file"
	case FileItemSave:
		return "Saves the current state to a json definition file"
	case FileItemGen:
		return "Generates source code from the current  state"
	case EditItemClear:
		return "Clear the grid and start fresh with verification if state changed"
	case EditItemPrev:
		return "Refresh the preview screen"
	case EditItemCopy:
		return "Copy the preview screen to the system clipboard"
	case HelpItemAbout:
		return "Brief description of the app"
	case HelpItemHelp:
		return "View the readme file"
	}
	return ""
}

func fileItemOpen() *fyne.MenuItem {
	return fyne.NewMenuItem(FileItemOpen, func() { doFileOpen() })
}

func fileItemSave() *fyne.MenuItem {
	return fyne.NewMenuItem(FileItemSave, func() { doFileSave() })
}

func fileItemGen() *fyne.MenuItem {
	return fyne.NewMenuItem(FileItemGen, func() { doFileGen() })
}

func editItemClear() *fyne.MenuItem {
	return fyne.NewMenuItem(EditItemClear, func() { doEditClear() })
}

func editItemPrev() *fyne.MenuItem {
	return fyne.NewMenuItem(EditItemPrev, func() { doEditPreview() })
}

func editItemCopy() *fyne.MenuItem {
	return fyne.NewMenuItem(EditItemCopy, func() { doEditCopy() })
}

func helpItemAbout() *fyne.MenuItem {
	return fyne.NewMenuItem(HelpItemAbout, func() { doHelpAbout() })
}

func helpItemHelp() *fyne.MenuItem {
	return fyne.NewMenuItem(HelpItemHelp, func() { doHelpHelp() })
}

func doFileOpen() {
	loadFromFile()
}

func doFileSave() {
	saveToFile(true) // true that its json
}

func doFileGen() {
	writeConstants()
}

func doEditClear() {
	verifyClear()
}

func doEditPreview() {
	refreshPreview()
}
func doEditCopy() {
	copyToClipBoard()
}

func doHelpAbout() {
	about := `
Constantly will help you create blocks of constants with     
optional related functions with the bare minimum of typing.
Featuring a grid layout with copying, concatenation and       
string tranformation.                                              `
	dlg := dialog.NewInformation("About Constantly", about, applicationData.mainWindow)
	dlg.Show()
}

func doHelpHelp() {
	about := `
Constantly will help you create blocks of constants with     
optional related functions with the bare minimum of typing.
Featuring a grid layout with copying, concatenation and       
string tranformation.                                              `

	dlg := dialog.NewInformation("Help", about, applicationData.mainWindow)
	dlg.Show()
}

func setMainMenu() {
	file := fyne.NewMenu("File", fileItemOpen(), fileItemSave(), fileItemGen())
	edit := fyne.NewMenu("Edit", editItemClear(), editItemPrev(), editItemCopy())
	help := fyne.NewMenu("Help", helpItemAbout(), helpItemHelp())
	main := fyne.NewMainMenu(file, edit, help)
	applicationData.mainWindow.SetMainMenu(main)
}

func copyToClipBoard() {
	applicationData.mainWindow.Clipboard().SetContent(previewString())
}
