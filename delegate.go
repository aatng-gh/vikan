package main

import "github.com/charmbracelet/bubbles/list"

var (
	BaseDelegate     = list.NewDefaultDelegate()
	NoCursorDelegate = func() list.DefaultDelegate {
		d := BaseDelegate
		d.Styles.SelectedTitle = d.Styles.NormalTitle
		d.Styles.SelectedDesc = d.Styles.NormalDesc
		return d
	}()
)
