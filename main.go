package main

import "fmt"

const (
	MAX_F     = 8
	MAX_C     = 8
	MAX_STEPS = 64
)

type posi struct {
	f int
	c int
}

func (p posi) in_board() bool {
	if (p.f >= 0) && (p.f < MAX_F) && (p.c >= 0) && (p.c < MAX_C) {
		return true
	} else {
		return false
	}
}

func (p posi) in_path(path [MAX_STEPS]posi, cant int) bool {
	for i := 0; i < cant; i++ {
		//fmt.Println(path[i])
		if p == path[i] {
			return true
		}
	}
	return false
}

func (p posi) possible_moves() []posi {
	var moves []posi
	possible := posi{f: p.f - 1, c: p.c - 2}
	if possible.in_board() {
		moves = append(moves, possible)
	}

	possible = posi{f: p.f - 2, c: p.c - 1}
	if possible.in_board() {
		moves = append(moves, possible)
	}

	possible = posi{f: p.f - 2, c: p.c + 1}
	if possible.in_board() {
		moves = append(moves, possible)

	}
	possible = posi{f: p.f - 1, c: p.c + 2}
	if possible.in_board() {
		moves = append(moves, possible)
	}

	possible = posi{f: p.f + 1, c: p.c + 2}
	if possible.in_board() {
		moves = append(moves, possible)
	}

	possible = posi{f: p.f + 2, c: p.c + 1}
	if possible.in_board() {
		moves = append(moves, possible)
	}

	possible = posi{f: p.f + 2, c: p.c - 1}
	if possible.in_board() {
		moves = append(moves, possible)
	}

	possible = posi{f: p.f + 1, c: p.c - 2}
	if possible.in_board() {
		moves = append(moves, possible)
	}

	return moves
}

func mov(current posi, path [MAX_STEPS]posi, step int) []posi {
	if step > 64 {
		fmt.Println("-")
		return make([]posi, 0)
	}
	moves := current.possible_moves()
	final_moves := make([]posi, 0, 8)
	for _, move := range moves {
		if !move.in_path(path, step) {
			final_moves = append(final_moves, move)
		}
	}
	if len(final_moves) == 0 {
		return make([]posi, 0)
	}
	now := make([]posi, 0, 64)
	path[step] = current
	for _, op := range final_moves {
		now = mov(op, path, step+1)
		if len(now)+step > 62 {
			return append(now, op)
		}
	}
	return append(now, final_moves[len(final_moves)-1])
}

func main() {
	x := posi{f: 0, c: 0}

	var path [MAX_STEPS]posi
	path[0] = x
	camino := mov(x, path, 1)

	for _, paso := range camino {
		fmt.Println(paso)
	}
	fmt.Println(x)

	fmt.Println("vim-go")
}
