package main

import "tictactoe/model"
import tea "github.com/charmbracelet/bubbletea"

func main() {
    
	p := tea.NewProgram(model.NewModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}

}


