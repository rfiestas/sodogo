package sodogo

// HelperBoard a collection of helpers, Examples for a 3x3 soduku
type HelperBoard struct {
	flats            int    //  3
	maxValue         int    //  9
	boardSize        int    // 81
	validValues      []int  // [1,2,3,4,5,6,7,8,9]
	flatGroups       []int  // [0,0,0,3,3,3,6,6,6,0,0,0,3,3,3,6,6,6,0,0,0,3,3,3,6,6,6,27,27,27,30,30,30,33,...]
	flatNeighbors    []int  // [0,1,2,9,10,11,18,19,20]
	streetYNeighbors []int  // [0,1,2,3,4,5,6,7,8]
	streetXNeighbors []int  // [0,9,18,27,36,45,54,63,72]
	nicePrint        string // Table caracters
}

// NewHelperBoard create a the board helpers
func NewHelperBoard(size int) (h HelperBoard) {
	maxValue := size * size
	boardSize := maxValue * maxValue
	h = HelperBoard{
		flats:     size,
		maxValue:  maxValue,
		boardSize: boardSize,
	}
	h.validValues = h.generateValidValues()
	h.flatGroups = h.generateFlatGroups()
	h.flatNeighbors = h.generateFlatNeighbors()
	h.streetYNeighbors = h.generateStreetYNeighbors()
	h.streetXNeighbors = h.generateStreetXNeighbors()
	h.nicePrint = h.generateNicePrint()
	return h
}

func (h HelperBoard) generateValidValues() (res []int) {
	res = []int{}
	for pos := 0; pos < h.maxValue; pos++ {
		res = append(res, pos+1)
	}
	return res
}

func (h HelperBoard) generateFlatGroups() (flatGroups []int) {
	flatGroups = []int{}
	for y := 0; y < h.maxValue; y++ {
		for x := 0; x < h.maxValue; x++ {
			flatY := (y / h.flats) * h.flats
			flatX := (x / h.flats) * h.flats
			group := (flatY * h.maxValue) + flatX

			flatGroups = append(flatGroups, group)
		}
	}
	return flatGroups
}

func (h HelperBoard) generateFlatNeighbors() (n neighbors) {
	n = []int{}
	for y := 0; y < h.flats; y++ {
		for x := 0; x < h.flats; x++ {
			value := (y * h.maxValue) + x
			n = append(n, value)
		}
	}
	return n
}

func (h HelperBoard) generateStreetYNeighbors() (n neighbors) {
	n = []int{}
	for pos := 0; pos < h.maxValue; pos++ {
		n = append(n, pos)
	}
	return n
}

func (h HelperBoard) generateStreetXNeighbors() (n neighbors) {
	n = []int{}
	for pos := 0; pos < h.maxValue; pos++ {
		n = append(n, pos*h.maxValue)
	}
	return n
}

func (n neighbors) getPotentialValues(validValues []int) (int, []int) {
	res := []int{}
	for num, vVal := range validValues {
		found := false
		for _, nVal := range n {
			if nVal == vVal {
				found = true
				break
			}
		}
		if !found {
			res = append(res, validValues[num])
		}
	}

	if len(res) != 1 {
		return 0, res
	}

	return res[0], res
}

func (p potential) getPotentialValues(myPotential potential) int {
	res := []int{}
	for num, vVal := range myPotential {
		found := false
		for _, nVal := range p {
			if nVal == 0 {
				return 0
			}
			if nVal == vVal {
				found = true
				break
			}
		}
		if !found {
			res = append(res, myPotential[num])
		}
	}

	if len(res) != 1 {
		return 0
	}

	return res[0]
}

func (h HelperBoard) generateNicePrint() (res string) {
	tableData := [25]string{"╔", "═══", "╤", "╦", "╗", "║", " %s ", "│", "║", "║", "╚", "═══", "╧", "╩", "╝", "╠", "═══", "╪", "╬", "╣", "╟", "───", "┼", "╫", "╢"}
	hIndex := 0
	inc := 0
	doubleMaxValue := h.maxValue * 2
	doubleFlats := h.flats * 2

	for y := 0; y < doubleMaxValue+1; y++ {
		if (y+1)%2 == 0 {
			inc = 5
		} else if y == doubleMaxValue {
			inc = 10
		} else if (y)%(doubleFlats) == 0 && y != 0 {
			inc = 15
		}
		for x := 0; x < doubleMaxValue+1; x++ {
			if x == 0 {
				hIndex = 0
			} else if x == doubleMaxValue {
				hIndex = 4
			} else {
				hIndex = 3
			}
			if (x+1)%2 == 0 {
				hIndex = 1
			} else {
				if (x)%(doubleFlats) != 0 {
					hIndex = 2
				}
			}
			res += tableData[hIndex+inc]

		}
		res += "\n"
		inc = 20
	}
	return res
}
