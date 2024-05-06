package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

var applicationData *appData

func init() {
	applicationData = newApplicationData()
	setMainMenu()
	applicationData.mainWindow.SetCloseIntercept(func() {
		verifyExit()
	})
}

func Run() {
	applicationData.mainWindow.CenterOnScreen()
	applicationData.mainWindow.Resize(fyne.Size{Width: 1240, Height: 680})

	tabs := container.NewAppTabs(
		container.NewTabItem("App Info", appInfoTab()),
		container.NewTabItem("Constants", columnGrid()),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	applicationData.mainWindow.SetContent(tabs)
	applicationData.mainWindow.ShowAndRun()
}
