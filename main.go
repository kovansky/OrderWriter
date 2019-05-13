package main

import (
	"fmt"
	"github.com/kovansky/OrderWriter/application/windows/launcher"
	"github.com/therecipe/qt/widgets"
	"os"
)

var (
	app            *widgets.QApplication
	launcherHolder *widgets.QMainWindow
)

func main() {
	fmt.Println("Starting an application")

	app = widgets.NewQApplication(len(os.Args), os.Args)

	launcherHolder = launcher.RunStart()

	launcherHolder.Move2(getCenterPoint())

	widgets.QApplication_Exec()
}

func getCenterPoint() (int, int) {
	screenRect := app.Desktop().ScreenGeometry(launcherHolder)
	centerPoint := screenRect.Center()
	return centerPoint.X(), centerPoint.Y()
}
