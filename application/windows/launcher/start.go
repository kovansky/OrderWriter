package launcher

import (
	"github.com/kovansky/OrderWriter/application"
	"github.com/kovansky/OrderWriter/application/windows/writer"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var (
	window *widgets.QMainWindow
)

func RunStart() *widgets.QMainWindow {
	window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle(application.Language.App.Title)
	window.SetAutoFillBackground(true)
	window.SetPalette(application.Palettes.WindowPalette)
	window.SetGeometry2(0, 0, 300, 800)

	area := widgets.NewQWidget(window, core.Qt__Widget)

	titleLabel := widgets.NewQLabel2(application.Language.App.Title, area, core.Qt__Widget)
	titleLabel.SetPalette(application.Palettes.MainPalette)

	// TODO: center
	layout := widgets.NewQVBoxLayout2(area)
	layout.AddWidget(titleLabel, 0, core.Qt__AlignTop)

	area.SetLayout(layout)
	window.SetCentralWidget(area)
	window.Show()

	writer.RunWriterMain()

	return window
}
