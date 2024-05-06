package app

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	//"github.com/westarver/fynewidgets"
)

const defaultRows = 15

type column struct {
	entries []*widget.Entry
}

type userEntries struct {
	columns                      map[string]column
	genStr, genAssoc, genComment *widget.Check
	genValue, genMarshal         *widget.Check
	consType                     *widget.Entry
	underlying                   *widget.SelectEntry
	valueBtn                     *contextMenuButton
	assocBtn                     *contextMenuButton
	valueshort                   bool
}

func (ue *userEntries) cell(c string, r int) string {
	return ue.columns[c].entries[r].Text
}

type appData struct {
	app         fyne.App
	mainWindow  fyne.Window
	preview     *fynewidgets.ReadOnlyEntry
	appName     *widget.Entry
	pkg         *widget.Entry
	author      *widget.Entry
	val         *contextMenuButton
	assoc       *contextMenuButton
	loadedPath  string
	jsonData    []byte
	userEntries *userEntries
	rows        int
	lastRow     int
	dirty       bool
}

func newUserEntries() *userEntries {
	c := make(map[string]column)
	c[Prefix] = column{entries: make([]*widget.Entry, defaultRows)}
	c[BaseID] = column{entries: make([]*widget.Entry, defaultRows)}
	c[Suffix] = column{entries: make([]*widget.Entry, defaultRows)}
	c[Type] = column{entries: make([]*widget.Entry, defaultRows)}
	c[Value] = column{entries: make([]*widget.Entry, defaultRows)}
	c[Assoc] = column{entries: make([]*widget.Entry, defaultRows)}

	return &userEntries{
		columns:    c,
		genStr:     nil,
		genAssoc:   nil,
		genComment: nil,
		genValue:   nil,
		genMarshal: nil,
		consType:   nil,
		underlying: nil,
		valueBtn:   nil,
		assocBtn:   nil,
		valueshort: true,
	}
}

func newApplicationData() *appData {
	app := app.New()
	win := app.NewWindow(MainWinTitle)

	return &appData{
		app:         app,
		mainWindow:  win,
		preview:     nil,
		appName:     nil,
		pkg:         nil,
		author:      nil,
		val:         nil,
		assoc:       nil,
		loadedPath:  "",
		jsonData:    []byte{},
		userEntries: newUserEntries(),
		rows:        defaultRows,
		lastRow:     1,
	}
}

// populatedCells() returns the number of rows in the Base Identifier column
func populatedCells() int {
	var num int
	for n, c := range applicationData.userEntries.columns[BaseID].entries {
		if c.Text == "" {
			continue
		}
		num = n
	}
	applicationData.lastRow = num
	return num
}

func (s appData) reset() {
	s.appName.SetText("")
	s.pkg.SetText("")
	s.author.SetText("")
	s.preview.SetText(strings.Repeat("\n", 16))
	s.loadedPath = ""
	s.userEntries.valueshort = true
	s.userEntries.consType.SetText("")
	s.userEntries.underlying.SetText("")
	s.userEntries.genStr.SetChecked(false)
	s.userEntries.genAssoc.SetChecked(false)
	s.userEntries.genComment.SetChecked(false)
	s.userEntries.genValue.SetChecked(false)
	s.userEntries.genMarshal.SetChecked(false)

	for _, v := range s.userEntries.columns {
		for i := 0; i < len(v.entries); i++ {
			v.entries[i].SetText("")
		}
	}
	s.lastRow = 1
	s.dirty = false
	s.jsonData = nil
}

func verifyExit() {
	if applicationData.dirty {
		dlg := dialog.NewConfirm("Work not saved", "The latest edits have not been saved. Do you want to exit anyway?", func(b bool) {
			if b {
				applicationData.mainWindow.Close()
			}
		}, applicationData.mainWindow)
		dlg.Show()
	} else {
		applicationData.mainWindow.Close()
	}
}

func verifyClear() {
	if applicationData.dirty {
		dlg := dialog.NewConfirm("Work not saved", "The latest edits have not been saved. Do you want to clear all?", func(b bool) {
			if b {
				applicationData.reset()
			}
		}, applicationData.mainWindow)
		dlg.Show()
	} else {
		applicationData.reset()
	}
}

/*
// (*SharedapplicationData) App returns a pointer to the unexported struct field app
func (S *sharedapplicationData) App() fyne.App {
	return S.app
}

// (*SharedapplicationData) Mainwindow returns a pointer to the unexported struct field mainWindow
func (S *SharedapplicationData) mainWindow fyne.Window {
	return S.mainWindow
}

// (*SharedapplicationData) Appname returns the value of the unexported struct field appName
func (S *SharedapplicationData) AppName() string {
	return S.appName.Text
}

// (*SharedapplicationData) LoadedPath returns the value of the unexported struct field loadedPath
func (S *SharedapplicationData) LoadedPath() string {
	return S.loadedPath
}
*/
/*
type SharedapplicationData struct {
	app         fyne.App
	mainWindow  fyne.Window
	preview     *fynewidgets.ReadOnlyEntry
	appName     *widget.Entry
	pkg         *widget.Entry
	author      *widget.Entry
	val         *contextMenuButton
	assoc       *contextMenuButton
	loadedPath  string
	jsonData    []byte
	userEntries *userEntries
	rows        int
	lastRow     int
	Dirty       bool
}
*/
