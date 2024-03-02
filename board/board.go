package board

import (
	"slices"
	"tictactoe/styles"
	"tictactoe/utils"

	"github.com/charmbracelet/lipgloss"
)

var (
	horizontal = "─"
	coOrds     = map[int][2]int{
		1: {0, 0},
		2: {0, 1},
		3: {0, 2},
		4: {1, 0},
		5: {1, 1},
		6: {1, 2},
		7: {2, 0},
		8: {2, 1},
		9: {2, 2},
	}
)

type Board struct{}

var board = [3][3]string{
	{" ", " ", " "},
	{" ", " ", " "},
	{" ", " ", " "},
}

func (b Board) BoardStr() string {

	return b.constructBoard()

}

func (b Board) GetXY(pos int) (x, y int) {
	val := coOrds[pos]
	x, y = val[0], val[1]
	return
}

func (b *Board) Update(x, y int, marker string) (updated bool) {
	if b.canPlace(x, y) {
		board[x][y] = marker
		updated = true
	}
	return
}

func (b *Board) Reset() {
	board = [3][3]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
}

func (b Board) canPlace(x, y int) bool {
	if (x >= 0 && x < 3) && (y >= 0 && y < 3) {
		return board[x][y] == " "
	}
	return false
}

func (b Board) constructBoard() string {
	var _board string
	for _, row := range board {
		_board += constructRow(row[0], row[1], row[2])
		_board += utils.NewLine()
	}
	return _board
}

func constructRow(c1, c2, c3 string) string {
	var (
		b1 = styles.CellStyle.Render(c1)
		b2 = styles.CellStyle.Render(c2)
		b3 = styles.CellStyle.Render(c3)
	)

	return lipgloss.JoinHorizontal(0, b1, b2, b3)

}

func (b Board) Draw() bool {
	for _, row := range board {
		for _, column := range row {
			if " " == column {
				return false
			}
		}
	}
	return true
}

func (b Board) Won(x, y int) bool {
	var check func(i []int) bool = func(i []int) bool {
		if i[0] == x && i[1] == y {
			return true
		}
		return false
	}

	// row
	if board[x][0] == board[x][1] && board[x][0] == board[x][2] {
		return true
	}
	// column
	if board[0][y] == board[1][y] && board[0][y] == board[2][y] {
		return true
	}
	// diag
	if slices.ContainsFunc([][]int{{0, 0}, {2, 2}, {1, 1}}, check) {
		if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
			return true
		} else {
			return false
		}
	}

	if slices.ContainsFunc([][]int{{0, 2}, {1, 1}, {2, 0}}, check) {
		if board[0][2] == board[1][1] && board[2][0] == board[1][1] {
			return true
		}
		return false
	}

	return false

}
