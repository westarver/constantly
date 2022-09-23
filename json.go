package constantly

import (
	"encoding/json"

	"fyne.io/fyne/v2/dialog"
)

func form2json() []byte {
	hdr := makeHeader()
	jsonbytes := []byte(`{"appdata":`)
	jsonb, _ := json.Marshal(hdr)
	jsonbytes = append(jsonbytes, jsonb...)
	jsonbytes = append(jsonbytes, []byte(",\n")...)
	jsonbytes = append(jsonbytes, []byte(`"formdata":`)...)
	type body struct {
		Formdata []constData
	}

	var form body
	for i := 0; i < AppData.rows; i++ {
		obj := form2struct(i)
		form.Formdata = append(form.Formdata, *obj)
	}

	jsonbody, _ := json.Marshal(form.Formdata)
	jsonbytes = append(jsonbytes, jsonbody...)
	jsonbytes = append(jsonbytes, []byte("\n}\n")...)
	return jsonbytes
}

func form2struct(row int) *constData {
	cd := constData{
		Prefix: AppData.userEntries.cell(Prefix, row),
		BaseId: AppData.userEntries.cell(BaseID, row),
		Suffix: AppData.userEntries.cell(Suffix, row),
		Value:  AppData.userEntries.cell(Value, row),
		Assoc:  AppData.userEntries.cell(Assoc, row),
	}
	if cd.Value != "" {
		cd.ShowValue = true
	}
	if row == 0 {
		cd.ShowType = true
	}
	if cd.ShowValue {
		cd.ShowAssign = true
	}
	return &cd
}

func makeHeader() *fileHeader {
	t := AppData.userEntries.consType.Text
	u := AppData.userEntries.underlying.Text
	if !goType() {
		if u == "" {
			u = AppData.userEntries.underlying.PlaceHolder
		}
	} else {
		u = ""
	}
	return &fileHeader{
		App:        AppData.AppName(),
		Pkg:        AppData.pkg.Text,
		Author:     AppData.author.Text,
		Type:       t,
		Under:      u,
		GenStr:     AppData.userEntries.genStr.Checked,
		GenAssoc:   AppData.userEntries.genAssoc.Checked,
		GenComment: AppData.userEntries.genComment.Checked,
	}

}

func json2form() {
	type appd struct {
		Appdata fileHeader
	}
	var appdata appd
	err := json.Unmarshal(AppData.jsonData, &appdata)
	if err != nil {
		dlg := dialog.NewError(err, AppData.mainWindow)
		dlg.Show()
		return
	}

	AppData.appName.SetText(appdata.Appdata.App)
	AppData.pkg.SetText(appdata.Appdata.Pkg)
	AppData.author.SetText(appdata.Appdata.Author)
	AppData.userEntries.consType.SetText(appdata.Appdata.Type)
	AppData.userEntries.underlying.SetText(appdata.Appdata.Under)
	AppData.userEntries.genStr.SetChecked(appdata.Appdata.GenStr)
	AppData.userEntries.genAssoc.SetChecked(appdata.Appdata.GenAssoc)
	AppData.userEntries.genComment.SetChecked(appdata.Appdata.GenComment)

	type body struct {
		Formdata []constData
	}
	var form body
	err = json.Unmarshal(AppData.jsonData, &form)
	if err != nil {
		dlg := dialog.NewError(err, AppData.mainWindow)
		dlg.Show()
		return
	}
	for i, f := range form.Formdata {
		AppData.userEntries.columns[Prefix].entries[i].SetText(f.Prefix)
		AppData.userEntries.columns[BaseID].entries[i].SetText(f.BaseId)
		AppData.userEntries.columns[Suffix].entries[i].SetText(f.Suffix)
		// type column is disabled
		if f.ShowValue {
			AppData.userEntries.columns[Value].entries[i].SetText(f.Value)
		}
		AppData.userEntries.columns[Assoc].entries[i].SetText(f.Assoc)
	}
	refreshPreview()
	AppData.dirty = false
}

type fileHeader struct {
	App        string `json:"app"`
	Pkg        string `json:"pkg"`
	Author     string `json:"author"`
	GenStr     bool   `json:"genstr"`
	GenAssoc   bool   `json:"genassoc"`
	GenComment bool   `json:"gencomment"`
	Type       string `json:"type"`
	Under      string `json:"under"`
}

type constData struct {
	Prefix     string `json:"prefix"`
	BaseId     string `json:"baseid"`
	Suffix     string `json:"suffix"`
	ShowType   bool   `json:"showtype"`
	ShowAssign bool   `json:"showassign"`
	ShowValue  bool   `json:"showvalue"`
	Value      string `json:"value"`
	Assoc      string `json:"assoc"`
}
