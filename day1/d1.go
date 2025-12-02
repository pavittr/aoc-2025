package day1

import (
	"strconv"
	"strings"
)

type buff struct {
	prev  *buff
	next  *buff
	value int
}

func New() *buff {
	start := &buff{}

	var middle *buff

	current := start
	for i := 1; i < 100; i++ {
		// create the previous
		next := &buff{
			value: i,
			prev:  current,
		}
		current.next = next
		current = next
		if i == 50 {
			middle = current
		}

	}
	current.next = start
	start.prev = current
	return middle
}

func Solve(puzzle string) int {

	// split each lien and read the instruction into a turn map
	instructions := []Instruction{}
	for _, line := range strings.Split(puzzle, "\n") {
		dir := line[0] == 'L'
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, Instruction{left: dir, steps: steps})
	}

	middle := New()
	zeroCOunt := 0
	for _, instr := range instructions {
		middle = turn(middle, instr.left, instr.steps)
		if middle.value == 0 {
			zeroCOunt++
		}
	}

	return zeroCOunt

}

type Instruction struct {
	left  bool
	steps int
}

func turn(start *buff, left bool, steps int) *buff {
	returnedBuff := start
	for i := 0; i < steps; i++ {
		if left {
			returnedBuff = returnedBuff.prev
		} else {
			returnedBuff = returnedBuff.next
		}
	}
	return returnedBuff
}
