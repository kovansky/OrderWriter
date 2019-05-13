package writer

import (
	"github.com/kovansky/OrderWriter/application"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func runNumberDialog(parent *widgets.QMainWindow) *widgets.QDialog {
	dialog := widgets.NewQDialog(parent, core.Qt__Dialog)
	dialog.SetWindowTitle(application.GetWinTitle(application.Language.App.Writer.NewOrder))

	label := widgets.NewQLabel2(application.Language.App.Writer.OrderNumber, dialog, 0)

	input := widgets.NewQLineEdit(dialog)
	input.SetValidator(gui.NewQIntValidator(input))

	dialogLayout := widgets.NewQFormLayout(dialog)
	dialogLayout.AddRow(label, input)

	dialog.SetLayout(dialogLayout)

	dialog.Show()

	return dialog
}
