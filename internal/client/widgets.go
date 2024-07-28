package client

import (
	"github.com/rivo/tview"
)

func (tui *Tui) InputMessage() *tview.TextArea {
	input := tview.NewTextArea().SetPlaceholder("Enter text here...")
	input.SetTitle("Message").SetBorder(true).SetInputCapture(tui.sendeMessage)
	return input
}

func (tui *Tui) Board() *tview.TextView {
	board := tview.NewTextView().SetDynamicColors(true)
	board.SetBorder(true).SetTitle("Chat")
	return board
}

func (tui *Tui) Users() *tview.TextView {
	users := tview.NewTextView()
	users.SetBorder(true).SetTitle("Users")
	return users
}

func (tui Tui) ChatWidget() *tview.Flex {
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tui.board, 0, 5, false).
			AddItem(tui.inputMessage, 0, 1, true), 0, 4, true).
		AddItem(tui.users, 0, 1, false)
	return flex
}
