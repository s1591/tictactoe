package model

import (
	"slices"
	"tictactoe/board"
	"tictactoe/players"
	"tictactoe/styles"
	"tictactoe/utils"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	spinner   spinner.Model
	keys      keyMap
	help      help.Model
	textInput textinput.Model
	players   players.Players
	board     board.Board
	over      bool
}

func (m Model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	over := []string{"ctrl+c", "esc", "q"}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if pressed := msg.String(); slices.Contains(over, pressed) {
			return m, tea.Quit
		} else if pressed == "enter" {
			answer := m.textInput.Value()
			m.textInput.Reset()
			return m.runGame(answer)
		}
	}

	m.textInput, _ = m.textInput.Update(msg)
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd

}

func (m Model) View() string {
	var ui string
	ui += styles.TitleStyle.Render("Tic-Tac-Toe")
	ui += utils.NewLine()
	ui += m.board.BoardStr()
	ui += utils.NewLine()
	if !m.over {
		ui += m.askForInput()
	} else {
		ui += "Game Over"
	}
	ui += utils.NewLine()
	ui += utils.NewLine()
	ui += m.help.View(m.keys)
	return styles.MainStyle.Render(ui) + utils.NewLine()
}

func NewModel() Model {
	ti := textinput.New()
	ti.Focus()
	s := spinner.New()
	s.Spinner = spinner.Dot
	return Model{
		keys: keyMap{
			key.NewBinding(key.WithKeys("q", "ctrl+c", "esc"), key.WithHelp("q/esc/ctrl+q", "quit")),
		},
		textInput: ti,
		help:      help.New(),
		spinner:   s,
	}
}

func (m Model) askForInput() string {
	m.textInput.Placeholder = ""
	m.textInput.Focus()
	var text string
	text += m.spinnerStr() + " "
	text += m.getMarker() + "'s turn "
	text += m.textInput.View()
	return text
}

func (m Model) spinnerStr() string {
	return m.spinner.View()
}

func (m Model) getXY(pos string) (x, y int) {
	val := utils.Atoi(pos)
	x, y = m.board.GetXY(val)
	return
}

func (m Model) getMarker() (marker string) {
	return [2]string{"X", "O"}[m.players.Turn()]
}

func (m Model) canQuit(x, y int) bool {
	if m.board.Draw() || m.board.Won(x, y) {
		return true
	}
	return false
}

func (m Model) updateBoard(x, y int, marker string) bool {
	return m.board.Update(x, y, marker)
}

func (m Model) runGame(answer string) (tea.Model, tea.Cmd) {
	x, y := m.getXY(answer)
	marker := m.getMarker()
	if !m.updateBoard(x, y, marker) {
		return m, nil
	}
	if m.canQuit(x, y) {
		m.over = true
	}
	m.players.NextTurn()

	return m, nil

}
