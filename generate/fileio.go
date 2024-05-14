package generate

import (
	"bytes"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"github.com/westarver/constantly/app"
)

type listable struct {
	fyne.URI
}

func (l listable) List() ([]fyne.URI, error) {
	return storage.List(l.URI)
}

func SaveToFile(json bool) {
	var err error
	var path string
	if json {
		path = app.LoadedPath()
	} else {
		path = ""
	}

	if len(path) == 0 {
		path, err = os.Getwd()
		if err != nil {
			dialog.NewError(err, app.MainWindow()).Show()
			return
		}
		path = filepath.Join(path, "~")
	}

	uri, err := storage.Parent(storage.NewFileURI(path))
	if err != nil {
		dialog.NewError(err, app.MainWindow()).Show()
		return
	}
	var fdialog *dialog.FileDialog
	var filter []string
	luri := listable{URI: uri}
	if json {
		fdialog = dialog.NewFileSave(doSaveJson, app.MainWindow())
		filter = []string{".json"}
	} else {
		fdialog = dialog.NewFileSave(doSaveGo, app.MainWindow())
		filter = []string{".go"}
	}

	fdialog.SetLocation(luri)
	fdialog.SetFilter(storage.NewExtensionFileFilter(filter))

	fdialog.Show()
}

func doSaveGo(uwc fyne.URIWriteCloser, e error) {
	if uwc == nil || e != nil {
		return
	}
	_, err := uwc.Write([]byte(PreviewString()))
	if err != nil {
		dialog.NewError(err, app.MainWindow()).Show()
		return
	}

	uwc.Close()
}

func doSaveJson(uwc fyne.URIWriteCloser, e error) {
	if uwc == nil || e != nil {
		return
	}

	app.SetJsonData(form2json())
	_, err := uwc.Write(app.JsonData())
	if err != nil {
		dialog.NewError(err, app.MainWindow()).Show()
		return
	}

	app.SetDirty(false)
	uwc.Close()
}

func LoadFromFile() {
	var err error
	path := app.LoadedPath

	if len(path) == 0 {
		path, err = os.Getwd()
		if err != nil {
			dialog.NewError(err, app.MainWindow()).Show()
			return
		}
		path = filepath.Join(path, "~")
	}

	uri, err := storage.Parent(storage.NewFileURI(path))
	if err != nil {
		dialog.NewError(err, app.MainWindow()).Show()
		return
	}
	luri := listable{URI: uri}
	fdialog := dialog.NewFileOpen(doLoad, app.MainWindow())
	fdialog.SetLocation(luri)
	fdialog.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))

	fdialog.Show()
}

func doLoad(uwc fyne.URIReadCloser, e error) {
	if uwc == nil {
		return
	}
	defer uwc.Close()

	if e != nil {
		dlg := dialog.NewError(e, app.MainWindow())
		dlg.Show()
	}

	stat, err := os.Stat(uwc.URI().Path())
	if err != nil {
		dlg := dialog.NewError(err, app.MainWindow())
		dlg.Show()
	}
	sz := stat.Size()

	b := make([]byte, sz+16)
	_, err = uwc.Read(b)
	if err != nil {
		dlg := dialog.NewError(e, app.MainWindow())
		dlg.Show()
		return
	}

	blob := bytes.TrimRight(b, "\n\t\x00")

	app.SetJsonData(bytes.TrimRight(b, "\n\t\x00"))
	app.SetLoadedPath(uwc.URI().Path())

	json2form(blob)

}

func reloadFile() {
	if app.LoadedPath == "" {
		info := dialog.NewInformation("Reload file", "No file is currently loaded.", app.MainWindow())
		info.Show()
		return
	}
	uri := storage.NewFileURI(app.LoadedPath())
	reader, err := storage.Reader(uri)
	doLoad(reader, err)
}
