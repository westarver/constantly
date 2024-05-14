// package constantly contains the data and logic to build the GUI
// and serve core data to the IO package
package app

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/westarver/fynewidgets"
)

const DefaultRows = 15
const MainWinTitle = "Constantly"

func init() {
	ApplicationData = NewAppData()
	ApplicationData.rows = DefaultRows
	ApplicationData.lastRow = 0
	SetMainMenu()
	ApplicationData.mainWindow.SetCloseIntercept(func() {
		VerifyExit()
	})
}

// Cell type to populate userEntries
type Cell struct {
	row   int
	col   string
	entry *widget.Entry
}

func (c Cell) Text() string {
	return c.entry.Text
}

func (c *Cell) SetText(t string) {
	c.entry.Text = t
}

func (c Cell) Row() int {
	return c.row
}

func (c Cell) Col() string {
	return c.col
}

func (c Cell) Entry() *widget.Entry { // careful with this
	return c.entry
}

// End Cell type

type column struct {
	entries []Cell
}

// userEntries type holds all of the entries input by the user
type userEntries struct {
	columns                      map[string]column
	genStr, genAssoc, genComment *widget.Check
	genValue, genMarshal         *widget.Check
	underlying                   *widget.SelectEntry
	consType                     *widget.Entry
}

func NewUserEntries() *userEntries {
	c := make(map[string]column)
	c[Prefix] = column{entries: make([]Cell, DefaultRows)}
	c[BaseID] = column{entries: make([]Cell, DefaultRows)}
	c[Suffix] = column{entries: make([]Cell, DefaultRows)}
	c[Type] = column{entries: make([]Cell, DefaultRows)}
	c[Value] = column{entries: make([]Cell, DefaultRows)}
	c[Assoc] = column{entries: make([]Cell, DefaultRows)}

	return &userEntries{
		columns:    c,
		genStr:     nil,
		genAssoc:   nil,
		genComment: nil,
		genValue:   nil,
		genMarshal: nil,
		consType:   nil,
		underlying: nil,
	}
}

// End userEntries type

var ApplicationData *appData

// appData struct stores core data shared at the application level
type appData struct {
	app         fyne.App
	mainWindow  fyne.Window
	preview     *fynewidgets.ReadOnlyEntry
	appName     *widget.Entry
	pkg         *widget.Entry
	author      *widget.Entry
	UserEntries *userEntries
	val         *contextMenuButton
	assoc       *contextMenuButton
	loadedPath  string
	jsonData    []byte
	rows        int
	lastRow     int
	dirty       bool
	valueshort  bool
}

func NewAppData() *appData {
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
		UserEntries: NewUserEntries(),
		rows:        DefaultRows,
		lastRow:     0,
		valueshort:  true,
	}
}

// End of appData

//exported functions to set/get userEntries data

// ResetApp will reset everything -- all data will disappear
func ResetApp() {
	ApplicationData.appName.SetText("")
	ApplicationData.pkg.SetText("")
	ApplicationData.author.SetText("")
	ApplicationData.preview.Text = strings.Repeat("\n", 16)
	ApplicationData.loadedPath = ""
	UserEntries().consType.SetText("Int")
	UserEntries().underlying.SetText("int")
	UserEntries().genStr.Checked = false
	UserEntries().genAssoc.Checked = false
	UserEntries().genComment.Checked = false
	UserEntries().genValue.Checked = false
	UserEntries().genMarshal.Checked = false
	//ApplicationData.val = nil
	//ApplicationData.assoc = nil
	for _, v := range UserEntries().columns {
		for i := 0; i < len(v.entries); i++ {
			v.entries[i].Entry().SetText("")
		}
	}
	ApplicationData.lastRow = 0
	ApplicationData.dirty = false
	ApplicationData.jsonData = []byte{}
	ApplicationData.valueshort = true
}

func UserEntries() *userEntries {
	return ApplicationData.UserEntries
}

func Column(label string) *column {
	c := UserEntries().columns[label]
	return &c
}

func SetGenStr(b bool) {
	UserEntries().genStr.Checked = b
	UserEntries().genStr.Refresh()
}
func GenStr() bool {
	return UserEntries().genStr.Checked
}

func SetGenAssoc(b bool) {
	UserEntries().genAssoc.Checked = b
	UserEntries().genAssoc.Refresh()
}
func GenAssoc() bool {
	return UserEntries().genAssoc.Checked
}

func SetGenComment(b bool) {
	UserEntries().genComment.Checked = b
	UserEntries().genComment.Refresh()
}
func GenComment() bool {
	return UserEntries().genComment.Checked
}

func SetGenValue(b bool) {
	UserEntries().genValue.Checked = b
	UserEntries().genValue.Refresh()
}
func GenValue() bool {
	return UserEntries().genValue.Checked
}

func SetGenMarshal(b bool) {
	UserEntries().genMarshal.Checked = b
	UserEntries().genMarshal.Refresh()
}
func GenMarshal() bool {
	return UserEntries().genMarshal.Checked
}

func SetConstType(t string) {
	UserEntries().consType.Text = t
}
func ConstType() string {
	return UserEntries().consType.Text
}

func SetUnderlyingType(t string) {
	UserEntries().underlying.Text = t
}
func UnderlyingType() string {
	return UserEntries().underlying.Text
}

func CellText(c string, r int) string {
	return CellAt(c, r).Text()
}

func CellColumn(c string, r int) string {
	return CellAt(c, r).Col()
}

func CellRow(c string, r int) int {
	return CellAt(c, r).Row()
}

func CellAt(c string, r int) *Cell {
	return &UserEntries().columns[c].entries[r]
}

func EntryAt(c string, r int) *widget.Entry {
	return CellAt(c, r).Entry()
}

func SetCellText(c string, r int, text string) {
	CellAt(c, r).Entry().Text = text
	CellAt(c, r).Entry().Refresh()
}

func CellRefresh(c string, r int) {
	EntryAt(c, r).Refresh()
}

func SetLastRow() {
	ApplicationData.lastRow = populatedCells()
}
func LastRow() int {
	return ApplicationData.lastRow
}

func SetDirty(b bool) {
	ApplicationData.dirty = b
}
func IsDirty() bool {
	return ApplicationData.dirty
}

func MainWindow() fyne.Window {
	return ApplicationData.mainWindow
}

func AppName() string {
	return ApplicationData.appName.Text
}
func SetAppName(a string) {
	ApplicationData.appName.Text = a
	ApplicationData.appName.Refresh()
}

func PackageName() string {
	return ApplicationData.pkg.Text
}
func SetPackageName(p string) {
	ApplicationData.pkg.Text = p
	ApplicationData.pkg.Refresh()
}

func Author() string {
	return ApplicationData.author.Text
}
func SetAuthor(a string) {
	ApplicationData.author.Text = a
	ApplicationData.author.Refresh()
}

func LoadedPath() string {
	return ApplicationData.loadedPath
}
func SetLoadedPath(p string) {
	ApplicationData.loadedPath = p
}

func JsonData() []byte {
	return ApplicationData.jsonData
}
func SetJsonData(j []byte) {
	ApplicationData.jsonData = j
}

// populatedCells() returns the number of rows in the Base Identifier column
func populatedCells() int {
	var num int
	for n, c := range UserEntries().columns[BaseID].entries {
		if c.Text() == "" {
			break
		}
		num = n
	}
	return num
}
func VerifyExit() {
	if IsDirty() {
		dlg := dialog.NewConfirm("Work not saved", "The latest edits have not been saved. Do you want to exit anyway?", func(b bool) {
			if b {
				MainWindow().Close()
			}
		}, MainWindow())
		dlg.Show()
	} else {
		MainWindow().Close()
	}
}

func VerifyClear() {
	if IsDirty() {
		dlg := dialog.NewConfirm("Work not saved", "The latest edits have not been saved. Do you want to clear all?", func(b bool) {
			if b {
				ResetApp()
			}
		}, MainWindow())
		dlg.Show()
	} else {
		ResetApp()
	}
}
