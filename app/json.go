package app

import (
	"encoding/json"

	"fyne.io/fyne/v2/dialog"
)

func form2json() []byte {
	hdr := makeheader()
	jsonbytes := []byte(`{"applicationData":`)
	jsonb, _ := json.Marshal(hdr)
	jsonbytes = append(jsonbytes, jsonb...)
	jsonbytes = append(jsonbytes, []byte(",\n")...)
	jsonbytes = append(jsonbytes, []byte(`"formdata":`)...)
	type body struct {
		formdata []constData
	}

	var form body
	for i := 0; i < applicationData.rows; i++ {
		obj := form2struct(i)
		form.formdata = append(form.formdata, *obj)
	}

	jsonbody, _ := json.Marshal(form.formdata)
	jsonbytes = append(jsonbytes, jsonbody...)
	jsonbytes = append(jsonbytes, []byte("\n}\n")...)
	return jsonbytes
}

func form2struct(row int) *constData {
	cd := constData{
		Prefix: applicationData.userEntries.cell(Prefix, row),
		BaseId: applicationData.userEntries.cell(BaseID, row),
		Suffix: applicationData.userEntries.cell(Suffix, row),
		Value:  applicationData.userEntries.cell(Value, row),
		Assoc:  applicationData.userEntries.cell(Assoc, row),
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

func makeheader() *fileheader {
	t := applicationData.userEntries.consType.Text
	u := applicationData.userEntries.underlying.Text
	if !goType() {
		if u == "" {
			u = applicationData.userEntries.underlying.PlaceHolder
		}
	} else {
		u = ""
	}
	return &fileheader{
		App:        applicationData.appName.Text,
		Pkg:        applicationData.pkg.Text,
		Author:     applicationData.author.Text,
		Type:       t,
		Under:      u,
		GenStr:     applicationData.userEntries.genStr.Checked,
		GenAssoc:   applicationData.userEntries.genAssoc.Checked,
		GenComment: applicationData.userEntries.genComment.Checked,
		GenValue:   applicationData.userEntries.genValue.Checked,
		GenMarshal: applicationData.userEntries.genMarshal.Checked,
	}

}

func json2form() {
	var header fileheader

	err := json.Unmarshal(applicationData.jsonData, &applicationData)
	if err != nil {
		dlg := dialog.NewError(err, applicationData.mainWindow)
		dlg.Show()
		return
	}

	applicationData.appName.SetText(header.App)
	applicationData.pkg.SetText(header.Pkg)
	applicationData.author.SetText(header.Author)
	applicationData.userEntries.consType.SetText(header.Type)
	applicationData.userEntries.underlying.SetText(header.Under)
	applicationData.userEntries.genStr.SetChecked(header.GenStr)
	applicationData.userEntries.genAssoc.SetChecked(header.GenAssoc)
	applicationData.userEntries.genComment.SetChecked(header.GenComment)
	applicationData.userEntries.genValue.SetChecked(header.GenValue)
	applicationData.userEntries.genMarshal.SetChecked(header.GenMarshal)

	//type body struct {
	var formdata []constData
	//}
	//var form body
	err = json.Unmarshal(applicationData.jsonData, &formdata)
	if err != nil {
		dlg := dialog.NewError(err, applicationData.mainWindow)
		dlg.Show()
		return
	}
	for i, f := range formdata {
		applicationData.userEntries.columns[Prefix].entries[i].SetText(f.Prefix)
		applicationData.userEntries.columns[BaseID].entries[i].SetText(f.BaseId)
		applicationData.userEntries.columns[Suffix].entries[i].SetText(f.Suffix)
		// type column is disabled
		if f.ShowValue {
			applicationData.userEntries.columns[Value].entries[i].SetText(f.Value)
		}
		applicationData.userEntries.columns[Assoc].entries[i].SetText(f.Assoc)
	}
	refreshPreview()
	applicationData.dirty = false
}

type fileheader struct {
	App        string `json:"app"`
	Pkg        string `json:"pkg"`
	Author     string `json:"author"`
	GenStr     bool   `json:"genstr"`
	GenAssoc   bool   `json:"genassoc"`
	GenComment bool   `json:"gencomment"`
	GenValue   bool   `json:"genvalue"`
	GenMarshal bool   `json:"genmarshal"`
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
