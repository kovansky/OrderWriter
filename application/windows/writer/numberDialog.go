package writer

import (
	"github.com/kovansky/OrderWriter/application"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"strconv"
)

func runNumberDialog() *widgets.QDialog {
	dialog := widgets.NewQDialog(window, core.Qt__Dialog)
	dialog.SetWindowTitle(application.GetWinTitle(application.Language.App.Writer.NewOrder))

	label := widgets.NewQLabel2(application.Language.App.Writer.OrderNumber, dialog, 0)

	input := widgets.NewQLineEdit(dialog)
	input.SetValidator(gui.NewQIntValidator(input))

	button := widgets.NewQPushButton2(application.Language.App.Next, dialog)
	button.ConnectClicked(func(checked bool) {
		entry, err := strconv.Atoi(input.Text())

		if err != nil {
			message := widgets.NewQMessageBox(dialog)
			message.SetIcon(widgets.QMessageBox__Critical)
			message.SetWindowTitle(application.GetWinTitle(application.Language.App.Error))
			message.SetText(application.Language.App.Error)
			message.SetStandardButtons(widgets.QMessageBox__Ok)
			message.Exec()

			panic(err)
		}
		if entry <= 0 {
			message := widgets.NewQMessageBox(dialog)
			message.SetIcon(widgets.QMessageBox__Critical)
			message.SetWindowTitle(application.GetWinTitle(application.Language.App.Error))
			message.SetText(application.Language.App.Writer.OrderNumberError)
			message.SetStandardButtons(widgets.QMessageBox__Ok)

			message.Exec()
		} else {
			orderNumber = entry

			dialog.Close()
			showValues()
		}
	})

	dialogLayout := widgets.NewQFormLayout(dialog)
	dialogLayout.AddRow(label, input)
	dialogLayout.AddRow5(button)

	dialog.SetLayout(dialogLayout)

	dialog.Show()

	return dialog
}
