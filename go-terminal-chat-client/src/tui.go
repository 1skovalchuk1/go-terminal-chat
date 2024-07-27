package src

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

func (tui *Tui) UpdateBoard(messages ChatMessages) {
	go func() {
		tui.tuiApp.QueueUpdateDraw(func() {
			tui.board.SetText(string(messages))
		})
	}()
}

func (tui *Tui) UpdateUsers(users ChatUsers) {
	go func() {
		tui.tuiApp.QueueUpdateDraw(func() {
			tui.users.SetText(string(users))
		})
	}()
}

func (tui *Tui) UpdateAll(messages string, users string) {
	go func() {
		tui.tuiApp.QueueUpdateDraw(func() {
			tui.board.SetText(messages)
			tui.users.SetText(users)
		})
	}()
}
func (tui *Tui) ClearInput() {
	tui.inputMessage.SetText("", false)
}

func (tui *Tui) SendeMessage(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == 13 { // enter '\n'
		text := tui.inputMessage.GetText()
		tui.manager.sendMessage(text)
		tui.ClearInput()
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
