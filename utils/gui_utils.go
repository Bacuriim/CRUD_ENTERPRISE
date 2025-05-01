package utils

import (
	"fyne.io/fyne/v2/widget"
)

func LimparCampos(entries ...*widget.Entry) {
	for _, entry := range entries {
		entry.SetText("")
		entry.Refresh()
	}
}
