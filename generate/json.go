package generate

import (
	"encoding/json"
	"errors"
	"fmt"

	"fyne.io/fyne/v2/dialog"
	"github.com/westarver/constantly/app"
)

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

func makeheader() *fileheader {

	return &fileheader{
		App:        app.AppName(),
		Pkg:       app.PackageName(),
		Author:     app.Author(),
		Type:       app.ConstType(),
		Under:      app.UnderlyingType(),
		GenStr:     app.GenStr(),
		GenAssoc:   app.GenAssoc(),
		GenComment: app.GenComment(),
		GenValue:   app.GenValue(),
		GenMarshal: app.GenMarshal(),
	}
}

func form2json() []byte {
	hdr := makeheader()
	jsonbytes := []byte(`{"applicationData":`)
	jsonb, _ := json.Marshal(hdr)
	jsonbytes = append(jsonbytes, jsonb...)
	jsonbytes = append(jsonbytes, []byte(",\n")...)
	jsonbytes = append(jsonbytes, []byte(`"formdata":`)...)

	r := app.LastRow() + 1
	var formdata []constData = make([]constData, r)

	for i := 0; i < r; i++ {
		obj := form2struct(i)
		formdata[i] = *obj
	}

	jsonbody, _ := json.Marshal(formdata)
	jsonbytes = append(jsonbytes, jsonbody...)
	jsonbytes = append(jsonbytes, []byte("\n}\n")...)

	return jsonbytes
}

func form2struct(row int) *constData {
	cd := constData{
		Prefix: app.CellText(Prefix, row),
		BaseId: app.CellText(BaseID, row),
		Suffix: app.CellText(Suffix, row),
		Value:  app.CellText(Value, row),
		Assoc:  app.CellText(Assoc, row),
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

func json2form(b []byte) {
	var appData map[string]interface{}
	err := json.Unmarshal(b, &appData)
	if err != nil {
		dlg := dialog.NewError(errors.New(err.Error()+" --could not unmarshal grid data"), MainWindow())
		dlg.Show()
		return
	}

	// fill in the AppInfoTab
	appdata := appData["applicationData"].(map[string]interface{})

	SetAppName(appdata["app"].(string))
	SetPackageName(appdata["pkg"].(string))
	SetAuthor(appdata["author"].(string))
	SetGenStr(appdata["genstr"].(bool))
	SetGenAssoc(appdata["genassoc"].(bool))
	SetGenComment(appdata["gencomment"].(bool))
	SetGenValue(appdata["genvalue"].(bool))
	SetGenMarshal(appdata["genmarshal"].(bool))
	SetConstType(appdata["type"].(string))
	SetUnderlyingType(appdata["under"].(string))

	// fill in the grid

	formdata := appData["formdata"]
	for n, c := range formdata.([]interface{}) {
		for label, r := range c.(map[string]interface{}) {
			switch label {
			case "showassign", "showvalue":
				continue
			case "showtype":
				if n > 0 {
					continue
				}
				Setapp.CellText(Type, n, ConstType())
			case "assoc":
				Setapp.CellText(Assoc, n, r.(string))
			case "value":
				Setapp.CellText(Value, n, fmt.Sprintf("%v", r))
			case "prefix":
				Setapp.CellText(Prefix, n, r.(string))
			case "baseid":
				Setapp.CellText(BaseID, n, r.(string))
			case "suffix":
				Setapp.CellText(Suffix, n, r.(string))
			}
		}
	}
	SetLastRow()
	SetDirty(false)
}
