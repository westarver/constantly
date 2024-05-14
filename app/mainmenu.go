package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/westarver/constantly/bridge"
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
	EditItemCopy = "Copy to Clipboard"
	// brief description of the app
	HelpItemAbout = "About"
	// view the readme file
	HelpItemHelp = "Help"
)

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
	bridge.LoadFromFile()
}

func doFileSave() {
	bridge.SaveToFile(true) // true that its json
}

func doFileGen() {
	bridge.WriteConstants()
}

func doEditClear() {
	VerifyClear()
}

func doEditPreview() {
	bridge.RefreshPreview()
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
	dlg := dialog.NewInformation("About Constantly", about, MainWindow())
	dlg.Show()
}

func doHelpHelp() {
	about := `
Constantly will help you create blocks of constants with     
optional related functions with the bare minimum of typing.
Featuring a grid layout with copying, concatenation and       
string tranformation.                                              `

	dlg := dialog.NewInformation("Help", about, MainWindow())
	dlg.Show()
}

func SetMainMenu() {
	file := fyne.NewMenu("File", fileItemOpen(), fileItemSave(), fileItemGen())
	edit := fyne.NewMenu("Edit", editItemClear(), editItemPrev(), editItemCopy())
	help := fyne.NewMenu("Help", helpItemAbout(), helpItemHelp())
	main := fyne.NewMainMenu(file, edit, help)
	MainWindow().SetMainMenu(main)
}

func copyToClipBoard() {
	MainWindow().Clipboard().SetContent(bridge.PreviewString())
}
