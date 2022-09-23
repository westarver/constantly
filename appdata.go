package constantly

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	fynewidgets "github.com/westarver/fyne-widgets"
)

const defaultRows = 15

type column struct {
	entries []*widget.Entry
}

type userEntries struct {
	columns                      map[string]column
	genStr, genAssoc, genComment *widget.Check
	consType                     *widget.Entry
	underlying                   *widget.SelectEntry
	valueBtn                     *contextMenuButton
	assocBtn                     *contextMenuButton
	valueshort                   bool
}

func (ue *userEntries) cell(c string, r int) string {
	return ue.columns[c].entries[r].Text
}

type SharedAppData struct {
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
		consType:   nil,
		underlying: nil,
		valueBtn:   nil,
		assocBtn:   nil,
		valueshort: true,
	}
}

func NewSharedAppData() *SharedAppData {
	app := app.New()
	win := app.NewWindow(MainWinTitle)

	return &SharedAppData{
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
	}
}

// populatedCells() returns the number of rows in the Base Identifier column
func populatedCells() int {
	var num int
	for _, c := range AppData.userEntries.columns[BaseID].entries {
		if c.Text == "" {
			break
		}
		num++
	}
	return num
}

func (s *SharedAppData) reset() {
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

	for _, v := range s.userEntries.columns {
		for i := 0; i < len(v.entries); i++ {
			v.entries[i].SetText("")
		}
	}
	AppData.dirty = false
	AppData.jsonData = nil
}

func verifyExit() {
	if AppData.dirty {
		dlg := dialog.NewConfirm("Work not saved", "The latest edits have not been saved. Do you want to exit anyway?", func(b bool) {
			if b {
				AppData.MainWindow().Close()
			}
		}, AppData.MainWindow())
		dlg.Show()
	} else {
		AppData.MainWindow().Close()
	}
}

func verifyClear() {
	if AppData.dirty {
		dlg := dialog.NewConfirm("Work not saved", "The latest edits have not been saved. Do you want to clear all?", func(b bool) {
			if b {
				AppData.reset()
			}
		}, AppData.MainWindow())
		dlg.Show()
	} else {
		AppData.reset()
	}
}

// (*SharedAppData) App returns a pointer to the unexported struct field app
func (S *SharedAppData) App() fyne.App {
	return S.app
}

// (*SharedAppData) Mainwindow returns a pointer to the unexported struct field mainWindow
func (S *SharedAppData) MainWindow() fyne.Window {
	return S.mainWindow
}

// (*SharedAppData) Appname returns the value of the unexported struct field appName
func (S *SharedAppData) AppName() string {
	return S.appName.Text
}

// (*SharedAppData) Loadedpath returns the value of the unexported struct field loadedPath
func (S *SharedAppData) LoadedPath() string {
	return S.loadedPath
}
