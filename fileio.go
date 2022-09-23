package constantly

import (
	"bytes"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

type listable struct {
	fyne.URI
}

func (l listable) List() ([]fyne.URI, error) {
	return storage.List(l.URI)
}

func saveToFile(json bool) {
	var err error
	var path string
	if json {
		path = AppData.loadedPath
	} else {
		path = ""
	}

	if len(path) == 0 {
		path, err = os.Getwd()
		if err != nil {
			dialog.NewError(err, AppData.mainWindow).Show()
			return
		}
		path = filepath.Join(path, "~")
	}

	uri, err := storage.Parent(storage.NewFileURI(path))
	if err != nil {
		dialog.NewError(err, AppData.mainWindow).Show()
		return
	}
	var fdialog *dialog.FileDialog
	var filter []string
	luri := listable{URI: uri}
	if json {
		fdialog = dialog.NewFileSave(doSaveJson, AppData.mainWindow)
		filter = []string{".json"}
	} else {
		fdialog = dialog.NewFileSave(doSaveGo, AppData.mainWindow)
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
	_, err := uwc.Write([]byte(previewString()))
	if err != nil {
		dialog.NewError(err, AppData.mainWindow).Show()
		return
	}

	uwc.Close()
}

func doSaveJson(uwc fyne.URIWriteCloser, e error) {
	if uwc == nil || e != nil {
		return
	}
	_, err := uwc.Write(form2json())
	if err != nil {
		dialog.NewError(err, AppData.mainWindow).Show()
		return
	}
	AppData.dirty = false
	uwc.Close()
}

func loadFromFile() {
	var err error
	path := AppData.loadedPath

	if len(path) == 0 {
		path, err = os.Getwd()
		if err != nil {
			dialog.NewError(err, AppData.mainWindow).Show()
			return
		}
		path = filepath.Join(path, "~")
	}

	uri, err := storage.Parent(storage.NewFileURI(path))
	if err != nil {
		dialog.NewError(err, AppData.mainWindow).Show()
		return
	}
	luri := listable{URI: uri}
	fdialog := dialog.NewFileOpen(doLoad, AppData.mainWindow)
	fdialog.SetLocation(luri)
	fdialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt", ".help", ".json"}))

	fdialog.Show()
}

func doLoad(uwc fyne.URIReadCloser, e error) {
	if uwc == nil {
		return
	}
	defer uwc.Close()

	if e != nil {
		dlg := dialog.NewError(e, AppData.mainWindow)
		dlg.Show()
	}

	stat, err := os.Stat(uwc.URI().Path())
	if err != nil {
		dlg := dialog.NewError(err, AppData.mainWindow)
		dlg.Show()
	}
	sz := stat.Size()

	b := make([]byte, sz+16)
	_, err = uwc.Read(b)
	if err != nil {
		dlg := dialog.NewError(e, AppData.mainWindow)
		dlg.Show()
		return
	}
	AppData.reset()
	AppData.jsonData = bytes.TrimRight(b, "\n\t\x00")
	AppData.loadedPath = uwc.URI().Path()
	json2form()
}

func reloadFile() {
	if AppData.loadedPath == "" {
		info := dialog.NewInformation("Reload file", "No file is currently loaded.", AppData.mainWindow)
		info.Show()
		return
	}
	uri := storage.NewFileURI(AppData.loadedPath)
	reader, err := storage.Reader(uri)
	doLoad(reader, err)
}
