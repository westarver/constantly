package constantly

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

var AppData *SharedAppData

func init() {
	AppData = NewSharedAppData()
	setMainMenu()
	AppData.mainWindow.SetCloseIntercept(func() {
		verifyExit()
	})
}

func Run() {
	AppData.mainWindow.CenterOnScreen()
	AppData.mainWindow.Resize(fyne.Size{Width: 1240, Height: 680})

	tabs := container.NewAppTabs(
		container.NewTabItem("App Info", appInfoTab()),
		container.NewTabItem("Constants", columnGrid()),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	AppData.mainWindow.SetContent(tabs)
	AppData.mainWindow.ShowAndRun()
}
