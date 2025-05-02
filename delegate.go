package main

import "github.com/charmbracelet/bubbles/list"

var (
	BaseDelegate     = list.NewDefaultDelegate()
	NoCursorDelegate = func() list.DefaultDelegate {
		d := list.NewDefaultDelegate()

		// make "selected" look just like "normal"
		d.Styles.SelectedTitle = d.Styles.NormalTitle
		d.Styles.SelectedDesc = d.Styles.NormalDesc
		return d
	}()
)
