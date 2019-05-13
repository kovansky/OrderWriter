package writer

import (
	"github.com/kovansky/OrderWriter/application"
	"github.com/kovansky/OrderWriter/types"
	"github.com/therecipe/qt/widgets"
	"log"
)

var (
	window *widgets.QMainWindow

	orderType   types.OrderType
	orderNumber int
)

func RunWriterMain() {
	window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle(application.GetWinTitle(application.Language.App.Writer.Title))
	window.SetAutoFillBackground(true)
	window.SetPalette(application.Palettes.WindowPalette)

	runTypeDialog()

	window.Show()
}

func showValues() {
	log.Println("Order type:", orderType.String())
	log.Println("Order number:", orderNumber)
}
