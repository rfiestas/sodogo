package sodogo

import (
	"reflect"
	"testing"
)

func test3x3Board() (h HelperBoard) {
	return HelperBoard{
		flats:     3,
		maxValue:  9,
		boardSize: 81,
	}
}

func test2x2Board() (h HelperBoard) {
	return HelperBoard{
		flats:     2,
		maxValue:  4,
		boardSize: 16,
	}
}

func TestBoard_generateValidValues(t *testing.T) {
	tests := []struct {
		name    string
		h       HelperBoard
		wantRes []int
	}{
		{
			name:    "3x3",
			h:       test3x3Board(),
			wantRes: neighbors{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:    "2x2",
			h:       test2x2Board(),
			wantRes: neighbors{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := tt.h.generateValidValues(); !reflect.DeepEqual(gotN, tt.wantRes) {
				t.Errorf("Board.generateValidValues = %v, want %v", gotN, tt.wantRes)
			}
		})
	}
}

func TestBoard_generateFlatGroups(t *testing.T) {
	tests := []struct {
		name    string
		h       HelperBoard
		wantRes []int
	}{
		{
			name:    "3x3",
			h:       test3x3Board(),
			wantRes: neighbors{0, 0, 0, 3, 3, 3, 6, 6, 6, 0, 0, 0, 3, 3, 3, 6, 6, 6, 0, 0, 0, 3, 3, 3, 6, 6, 6, 27, 27, 27, 30, 30, 30, 33, 33, 33, 27, 27, 27, 30, 30, 30, 33, 33, 33, 27, 27, 27, 30, 30, 30, 33, 33, 33, 54, 54, 54, 57, 57, 57, 60, 60, 60, 54, 54, 54, 57, 57, 57, 60, 60, 60, 54, 54, 54, 57, 57, 57, 60, 60, 60},
		},
		{
			name:    "2x2",
			h:       test2x2Board(),
			wantRes: neighbors{0, 0, 2, 2, 0, 0, 2, 2, 8, 8, 10, 10, 8, 8, 10, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := tt.h.generateFlatGroups(); !reflect.DeepEqual(gotN, tt.wantRes) {
				t.Errorf("Board.generateValidValues = %v, want %v", gotN, tt.wantRes)
			}
		})
	}
}

func TestBoard_generateFlatNeighbors(t *testing.T) {
	tests := []struct {
		name  string
		h     HelperBoard
		wantN neighbors
	}{
		{
			name:  "3x3",
			h:     test3x3Board(),
			wantN: neighbors{0, 1, 2, 9, 10, 11, 18, 19, 20},
		},
		{
			name:  "2x2",
			h:     test2x2Board(),
			wantN: neighbors{0, 1, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := tt.h.generateFlatNeighbors(); !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("Board.generateFlatNeighbors() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestBoard_generateStreetYNeighbors(t *testing.T) {
	tests := []struct {
		name  string
		h     HelperBoard
		wantN neighbors
	}{
		{
			name:  "3x3",
			h:     test3x3Board(),
			wantN: neighbors{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:  "2x2",
			h:     test2x2Board(),
			wantN: neighbors{0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := tt.h.generateStreetYNeighbors(); !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("Board.generateStreetYNeighbors() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
func TestBoard_generateStreetXNeighbors(t *testing.T) {
	tests := []struct {
		name  string
		h     HelperBoard
		wantN neighbors
	}{
		{
			name:  "3x3",
			h:     test3x3Board(),
			wantN: neighbors{0, 9, 18, 27, 36, 45, 54, 63, 72},
		},
		{
			name:  "2x2",
			h:     test2x2Board(),
			wantN: neighbors{0, 4, 8, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := tt.h.generateStreetXNeighbors(); !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("Board.generateStreetXNeighbors() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_neighbors_getPotentialValues(t *testing.T) {
	type args struct {
		validValues []int
	}
	tests := []struct {
		name  string
		n     neighbors
		args  args
		want  int
		want1 []int
	}{
		{
			name: "Get value",
			n:    neighbors{2, 3, 4, 5, 6, 7, 8, 9},
			args: args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want:  1,
			want1: []int{1},
		},
		{
			name: "Get potential values",
			n:    neighbors{0, 0, 3, 4, 5, 6, 7, 8, 9},
			args: args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want:  0,
			want1: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.n.getPotentialValues(tt.args.validValues)
			if got != tt.want {
				t.Errorf("neighbors.getPotentialValues() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("neighbors.getPotentialValues() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_potential_getPotentialValues(t *testing.T) {
	type args struct {
		myPotential potential
	}
	tests := []struct {
		name string
		p    potential
		args args
		want int
	}{
		{
			name: "Get potential value",
			p:    potential{2, 3, 4, 5},
			args: args{
				potential{1, 2, 5},
			},
			want: 1,
		},
		{
			name: "Get 0 when some potential is 0",
			p:    potential{2, 3, 4, 5, 0},
			args: args{
				potential{1, 2, 5},
			},
			want: 0,
		},
		{
			name: "No potential values",
			p:    potential{2, 3, 4, 5, 1},
			args: args{
				potential{1, 2, 5},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.getPotentialValues(tt.args.myPotential); got != tt.want {
				t.Errorf("potential.getPotentialValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
