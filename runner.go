package constantly

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/westarver/constantly/app
)

func Run() {
	app.MainWindow().CenterOnScreen()
	app.MainWindow().Resize(fyne.Size{Width: 1240, Height: 680})

	tabs := container.NewAppTabs(
		container.NewTabItem("App Info", appInfoTab()),
		container.NewTabItem("Constants", columnGrid()),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	app.MainWindow().SetContent(tabs)
	app.MainWindow().ShowAndRun()
}


