package client

import (
	"time"

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

func NewTui(manager *Manager) *Tui {
	tui := Tui{
		manager: manager,
		tuiApp:  tview.NewApplication(),
	}
	tui.board = tui.Board()
	tui.users = tui.Users()
	tui.inputMessage = tui.InputMessage()
	return &tui
}

func (t *Tui) updateBoard(messages string) {
	// TODO CONST
	time.Sleep(time.Second / 10)
	//
	go func() {
		t.tuiApp.QueueUpdateDraw(func() {

			t.board.SetText(messages)
			t.board.ScrollToEnd()

		})
	}()
}

func (t *Tui) updateUsers(users string) {
	// TODO CONST
	time.Sleep(time.Second / 10)
	//
	go func() {
		t.tuiApp.QueueUpdateDraw(func() {

			t.users.SetText(users)

		})
	}()
}

func (t *Tui) updateAll(users string, messages string) {
	time.Sleep(time.Second / 10)
	go func() {
		t.tuiApp.QueueUpdateDraw(func() {
			t.board.SetText(messages)
			t.board.ScrollToEnd()
			t.users.SetText(users)
		})
	}()
}

func (t *Tui) clearInput() {
	t.inputMessage.SetText("", false)
}

func (t *Tui) sendeMessage(event *tcell.EventKey) *tcell.EventKey {

	if event.Name() == "Enter" && event.Modifiers()&tcell.ModAlt == 0 {
		text := t.inputMessage.GetText()
		t.manager.send(text)
		t.clearInput()
		return nil
	}
	return event
}

func (t *Tui) exit(event *tcell.EventKey) *tcell.EventKey {
	if event.Name() == "Esc" {
		t.manager.close()
		return nil
	}
	return event
}

func (t *Tui) Run() {
	err := t.tuiApp.
		SetRoot(t.ChatWidget(), true).
		EnableMouse(true).
		SetInputCapture(t.exit).
		Run()
	if err != nil {
		panic(err)
	}
}

func (t *Tui) close() {
	t.tuiApp.Stop()
}
