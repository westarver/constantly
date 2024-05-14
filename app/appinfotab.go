package app

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/westarver/constantly/bridge"
	"github.com/westarver/fynewidgets"
)

func appInfoTab() *fyne.Container {
	infoLabel1 := widget.NewLabelWithStyle("Application Data", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	info1 := widget.NewLabel("Application Name")
	info1.TextStyle = fyne.TextStyle{Bold: true}
	text1 := widget.NewEntry()
	text1.OnSubmitted = func(s string) { SetDirty(true) }
	info2 := widget.NewLabel("Package Name For Generated Code")
	info2.TextStyle = fyne.TextStyle{Bold: true}
	text2 := widget.NewEntry()
	text2.OnSubmitted = func(s string) { SetDirty(true) }
	info3 := widget.NewLabel("Author")
	info3.TextStyle = fyne.TextStyle{Bold: true}
	text3 := widget.NewEntry()
	text3.OnSubmitted = func(s string) { SetDirty(true) }
	f := container.New(layout.NewFormLayout(), info1, text1, info2, text2, info3, text3)
	form1 := container.NewVBox(infoLabel1, f)

	divider1 := widget.NewSeparator()
	divider2 := widget.NewSeparator()

	infoLabel2 := widget.NewLabelWithStyle("Preview (read only)", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	text4 := fynewidgets.NewReadOnlyEntry()
	text4.SetText(strings.Repeat("\n", 16))
	preview := container.NewVScroll(container.NewVBox(text4, layout.NewSpacer()))
	preview.SetMinSize(fyne.Size{
		Width:  text4.Size().Width,
		Height: 400,
	})

	ApplicationData.appName = text1
	ApplicationData.pkg = text2
	ApplicationData.author = text3
	ApplicationData.preview = text4

	button1 := widget.NewButton("Load From File", func() {
		bridge.LoadFromFile()
	})
	button2 := widget.NewButton("Save", func() {
		bridge.SaveToFile(true)
	})
	button3 := widget.NewButton("Generate Source File", func() { bridge.WriteConstants() })
	button4 := widget.NewButton("Reload", func() { bridge.ReloadFile() })
	btnbox := container.NewHBox(button1, button2, layout.NewSpacer(), button4, button3)
	return container.NewVBox(form1, divider1, infoLabel2, preview, layout.NewSpacer(), divider2, btnbox)
}
