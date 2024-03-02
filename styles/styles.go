package styles

import "github.com/charmbracelet/lipgloss"

var (
	myCuteBorder = lipgloss.Border{
		Bottom: "─",
	}
	MainStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			PaddingRight(20).PaddingLeft(20)
	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("10")).
			PaddingLeft(5).
			PaddingRight(3).
			//Border(lipgloss.RoundedBorder()).
			Underline(true).
			Faint(true)
	CellStyle = lipgloss.NewStyle().
			Border(myCuteBorder).
			PaddingLeft(2).
			PaddingRight(2)
)
