package client

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Tui struct {
	tuiApp       *tview.Application
	board        *tview.TextView
	users        *tview.TextView
	inputMessage *tview.TextArea
	manager      *Manager
}

func (tui Tui) Init(manager *Manager) *Tui {
	tui.manager = manager
	tui.tuiApp = tview.NewApplication().EnableMouse(true)
	tui.board = tui.Board()
	tui.users = tui.Users()
	tui.inputMessage = tui.InputMessage()
	return &tui
}

func (tui *Tui) updateBoard(messages string) {
	go func() {
		tui.tuiApp.QueueUpdateDraw(func() {
			tui.board.SetText(messages)
			tui.board.ScrollToEnd()
		})
	}()
}

func (tui *Tui) updateUsers(users string) {
	go func() {
		tui.tuiApp.QueueUpdateDraw(func() {
			tui.users.SetText(users)
		})
	}()
}

func (tui *Tui) updateAll(users string, messages string) {
	go func() {
		tui.tuiApp.QueueUpdateDraw(func() {
			tui.board.SetText(messages)
			tui.board.ScrollToEnd()
			tui.users.SetText(users)
		})
	}()
}

func (tui *Tui) clearInput() {
	tui.inputMessage.SetText("", false)
}

func (tui *Tui) sendeMessage(event *tcell.EventKey) *tcell.EventKey {

	if event.Name() == "Enter" && event.Modifiers()&tcell.ModAlt == 0 {
		text := tui.inputMessage.GetText()
		tui.manager.sendMessage(text)
		tui.clearInput()
		return nil
	}
	return event
}

func (tui *Tui) Exit(event *tcell.EventKey) *tcell.EventKey {
	if event.Name() == "Esc" {
		tui.manager.close()
		return nil
	}
	return event
}

func (tui *Tui) Run() {
	err := tui.tuiApp.
		SetRoot(tui.ChatWidget(), true).
		SetInputCapture(tui.Exit).EnableMouse(false).
		Run()
	if err != nil {
		panic(err)
	}
}

func (tui *Tui) close() {
	tui.tuiApp.Stop()
}
