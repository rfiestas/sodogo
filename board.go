package sodogo

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

/*
     Board example

  0 0 4  3 0 0  2 0 9
  0 0 5  0 0 9  0 0 1
  0 7 0  0 6 0  0 4 3

  0 0 6  0 0 2  0 8 7
  1 9 0  0 0 7  4 0 0
  0 5 0  0 8 3  0 0 0

  6 0 0  0 0 0  1 0 5
  0 0 3  5 0 8  6 9 0
  0 4 2  9 1 0  3 0 0

*/

//Board sudoku board data
type Board struct {
	data    []cell        // cell value and potential values
	helpers HelperBoard   // helpers to calculate neighbors
	Steps   int           // 0 steps
	Elapsed time.Duration // 0 elapsed time
}
type cell struct {
	value     int       // cell value
	potential potential // potential cell values
}

type neighbors []int // neighbors values
type potential []int // potential values

type flatNeighborsPotential struct {
	value int
}

type streetYNeighborsPotential struct {
	value int
}

type streetXNeighborsPotential struct {
	value int
}

type neighborsPotential interface {
	getNeighborsPotentialValues(h HelperBoard) (int, []int)
}

func (p flatNeighborsPotential) getNeighborsPotentialValues(h HelperBoard) (int, []int) {
	return h.flatGroups[p.value], h.flatNeighbors

}

func (p streetYNeighborsPotential) getNeighborsPotentialValues(h HelperBoard) (int, []int) {
	return p.value - p.value%h.maxValue, h.streetYNeighbors
}

func (p streetXNeighborsPotential) getNeighborsPotentialValues(h HelperBoard) (int, []int) {
	return p.value % h.maxValue, h.streetXNeighbors
}

// NewBoard create a new board
func NewBoard(h HelperBoard) (b Board) {
	b = Board{
		data:    make([]cell, h.boardSize),
		helpers: h,
		Steps:   0,
		Elapsed: 0,
	}

	return b
}

// LoadFromString converts a string to a board
func (b Board) LoadFromString(board string) error {
	if len(board) != b.helpers.boardSize {
		return fmt.Errorf("A valid board definition contains %d caracters, not %d", b.helpers.boardSize, len(board))
	}

	for inc := 0; inc < len(board); inc++ {
		value, err := strconv.Atoi(string(board[inc]))
		if err != nil {
			value = 0
		}
		b.data[inc] = cell{
			value:     value,
			potential: []int{value},
		}
	}
	return nil
}

// Solve the Sudoku
func (b *Board) Solve() bool {
	start := time.Now()
	var np []neighborsPotential
	step := 1
	for !b.isSolved() {
		stepChanges := 0
		for pos := 0; pos < b.helpers.boardSize; pos++ {

			if value := b.getValue(pos); value == 0 {
				allNeighborsValues := b.getAllNeighborsValues(pos)
				value, potentialValues := allNeighborsValues.getPotentialValues(b.helpers.validValues)
				if len(b.getPotential(pos)) != len(potentialValues) {
					b.setPotential(pos, potentialValues)
					stepChanges++
				}
				if value != 0 {
					b.setValue(pos, value)
					stepChanges++
					continue
				}
				flat := flatNeighborsPotential{pos}
				streetY := streetYNeighborsPotential{pos}
				streetX := streetXNeighborsPotential{pos}
				np = []neighborsPotential{flat, streetY, streetX}
				for _, f := range np {
					inc, helperNeighbors := f.getNeighborsPotentialValues(b.helpers)
					neighborsPotentialValue := b.getNeighborsPotentialValues(&pos, helperNeighbors, inc)
					value = neighborsPotentialValue.getPotentialValues(potentialValues)

					if value != 0 {
						b.setValue(pos, value)
						stepChanges++
						break
					}

				}
			}
		}
		if stepChanges == 0 {
			break
		}
		step++
	}
	b.Steps = step
	b.Elapsed = time.Since(start)
	return b.isSolved()
}

// String returns the board as string
func (b *Board) String() (res string) {
	var buffer bytes.Buffer

	for pos := 0; pos < b.helpers.boardSize; pos++ {
		buffer.WriteString(strconv.Itoa(b.data[pos].value))
	}
	return buffer.String()
}

// NicePrint print the sudoku human representation
func (b *Board) NicePrint() string {
	output := []interface{}{}
	for pos := 0; pos < b.helpers.boardSize; pos++ {
		value := " "
		if v := b.getValue(pos); v != 0 {
			value = fmt.Sprintf("%d", v)
		}
		output = append(output, value)
	}
	return fmt.Sprintf(b.helpers.generateNicePrint(), output...)
}

// IsSolved returns if the board is solved
func (b *Board) isSolved() (solved bool) {
	for pos := 0; pos < b.helpers.boardSize; pos++ {
		value := b.getValue(pos)
		if value == 0 {
			return false
		}
	}

	return true
}

// IsValid returns if the board is valid
func (b *Board) IsValid() (solved bool) {
	for pos := 0; pos < b.helpers.boardSize; pos++ {
		r := unique(b.getFlatNeighborsValues(pos))
		if !r {
			return false
		}
		inc := pos - (pos % b.helpers.maxValue)
		r = unique(b.getNeighborsValues(pos, b.helpers.streetYNeighbors, inc, false))
		if !r {
			return false
		}
		inc = pos % b.helpers.maxValue
		r = unique(b.getNeighborsValues(pos, b.helpers.streetXNeighbors, inc, false))
		if !r {
			return false
		}
	}
	return true
}

// unique check duplicated numbers on a list
func unique(intSlice []int) bool {
	keys := make(map[int]bool)
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
		} else {
			if entry != 0 {
				return false
			}
		}
	}
	return true
}

// getValue returns the cell value
func (b *Board) getValue(pos int) (value int) {
	return b.data[pos].value
}

// setValue save a value on a cell
func (b *Board) setValue(pos int, value int) {
	b.data[pos].value = value
	if value != 0 {
		b.setPotential(pos, []int{value})
	}
}

// getPotential returns the cell potential values
func (b *Board) getPotential(pos int) (value []int) {
	return b.data[pos].potential
}

// setPotential save a the potential values on a cell
func (b *Board) setPotential(pos int, values []int) {
	b.data[pos].potential = values
}

// getAllNeighborsValues returns all neighbors values
func (b *Board) getAllNeighborsValues(pos int) (n neighbors) {
	n = b.getFlatNeighborsValues(pos)
	n = append(n, b.getStreetYNeighborsValues(pos)...)
	n = append(n, b.getStreetXNeighborsValues(pos)...)
	return n
}

// getFlatNeighborsValues returns the flat neighbors values
func (b *Board) getFlatNeighborsValues(p int) (n neighbors) {
	inc := b.helpers.flatGroups[p]
	return b.getNeighborsValues(p, b.helpers.flatNeighbors, inc, false)
}

// getStreetYNeighborsValues returns the street Y neighbors values
func (b *Board) getStreetYNeighborsValues(p int) (n neighbors) {
	inc := p - p%b.helpers.maxValue
	return b.getNeighborsValues(p, b.helpers.streetYNeighbors, inc, true)
}

// getStreetXNeighborsValues returns the street X neighbors values
func (b *Board) getStreetXNeighborsValues(p int) (n neighbors) {
	inc := p % b.helpers.maxValue
	return b.getNeighborsValues(p, b.helpers.streetXNeighbors, inc, true)
}

// getNeighborsValues returns the neighbors values
func (b *Board) getNeighborsValues(p int, helpersNeighbors []int, inc int, skipFlatNeighbors bool) (n neighbors) {
	flatGroup := b.helpers.flatGroups[p]
	for _, pos := range helpersNeighbors {
		if skipFlatNeighbors && flatGroup == b.helpers.flatGroups[pos+inc] {
			continue
		}
		n = append(n, b.getValue(pos+inc))
	}
	return n
}

// getNeighborsPotentialValues returns the neighbors potential values
func (b *Board) getNeighborsPotentialValues(p *int, helpersNeighbors []int, inc int) (potential potential) {
	for _, pos := range helpersNeighbors {
		if *p == pos+inc {
			continue
		}
		potential = append(potential, b.getPotential(pos+inc)...)
	}
	return potential
}
