package client

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (tui Tui) Init(manager *Manager) *Tui {
	tui.tuiApp = tview.NewApplication().EnableMouse(true)
	tui.manager = manager
	tui.board = tui.Board()
	tui.users = tui.Users()
	tui.inputMessage = tui.InputMessage()
	return &tui
}

func (tui *Tui) updateBoard(messages ChatMessages) {
	go func() {
		tui.tuiApp.QueueUpdateDraw(func() {
			tui.board.SetText(string(messages))
			tui.board.ScrollToEnd()
		})
	}()
}

func (tui *Tui) updateUsers(users ChatUsers) {
	go func() {
		tui.tuiApp.QueueUpdateDraw(func() {
			tui.users.SetText(string(users))
		})
	}()
}

func (tui *Tui) clearInput() {
	tui.inputMessage.SetText("", false)
}

func (tui *Tui) sendeMessage(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == 13 { // enter '\n'
		text := tui.inputMessage.GetText()
		tui.manager.sendMessage(text)
		tui.clearInput()
		return nil
	}
	if event.Rune() == 0 {
		tui.manager.close()
		return nil
	}
	return event
}

func (tui *Tui) Run() {
	if err := tui.tuiApp.SetRoot(tui.ChatWidget(), true).Run(); err != nil {
		panic(err)
	}
}

func (tui *Tui) close() {
	tui.tuiApp.Stop()
}
