package writer

import (
	"github.com/kovansky/OrderWriter/application"
	"github.com/kovansky/OrderWriter/types"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func runTypeDialog() *widgets.QDialog {
	dialog := widgets.NewQDialog(window, core.Qt__Dialog)
	dialog.SetWindowTitle(application.GetWinTitle(application.Language.App.Writer.NewOrder))

	label := widgets.NewQLabel2(application.Language.App.Writer.OrderType, dialog, 0)

	input := widgets.NewQComboBox(dialog)

	for _, element := range types.GetOrderTypes() {
		input.AddItem(element.String(), core.NewQVariant7(int(element)))
	}

	button := widgets.NewQPushButton2(application.Language.App.Next, dialog)
	button.ConnectClicked(func(checked bool) {
		orderType = types.OrderType(input.CurrentIndex())

		dialog.Close()
		runNumberDialog()
	})

	dialogLayout := widgets.NewQFormLayout(dialog)
	dialogLayout.AddRow(label, input)
	dialogLayout.AddRow5(button)

	dialog.SetLayout(dialogLayout)

	dialog.Show()

	return dialog
}
