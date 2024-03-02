package players

import "math/rand"

type Players struct{}

var turn = 0
var p1 = "X"
var p2 = "O"

func (p Players) P1() string {
	return p1
}

func (p Players) P2() string {
	return p2
}

func (p Players) DecideFirstPlayer() string {
	return [2]string{"p1", "p2"}[rand.Intn(2)]
}

func (p *Players) Assign(P1, P2 string) {
	p1, p2 = P1, P2
}

func (p *Players) AssignDefault() {
	p1, p2 = "X", "O"
}

func (p Players) Turn() int {
	return turn
}

func (p *Players) NextTurn() {
	turn = (turn + 1) % 2
}

func (p *Players) ResetTurns() {
	turn ^= turn
}
