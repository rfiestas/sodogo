package sodogo

import (
	"reflect"
	"testing"
)

func test3x3BoardUnsolved() (b Board) {
	helper := NewHelperBoard(3)
	board := NewBoard(helper)
	_ = board.LoadFromString("004300209005009001070060043006002087190007400050083000600000105003508690042910300")
	return board
}

func test3x3BoardImpossible() (b Board) {
	helper := NewHelperBoard(3)
	board := NewBoard(helper)
	_ = board.LoadFromString("800000000003600000070090200050007000000045700000100030001000068008500010090000400")
	return board
}

func test2x2BoardSolved() (b Board) {
	helper := NewHelperBoard(2)
	board := NewBoard(helper)
	_ = board.LoadFromString("1234341221434321")
	return board
}

func test2x2BoardInvalidFlat() (b Board) {
	helper := NewHelperBoard(2)
	board := NewBoard(helper)
	_ = board.LoadFromString("1234141221434321")
	return board
}

func test2x2BoardInvalidY() (b Board) {
	helper := NewHelperBoard(2)
	board := NewBoard(helper)
	_ = board.LoadFromString("1214341221434321")
	return board
}

func test2x2BoardInvalidX() (b Board) {
	helper := NewHelperBoard(2)
	board := NewBoard(helper)
	_ = board.LoadFromString("1234341211434321")
	return board
}

func TestBoard_LoadFromString(t *testing.T) {
	type args struct {
		board string
	}
	tests := []struct {
		name    string
		b       Board
		args    args
		wantErr bool
	}{
		{
			name:    "2x2",
			b:       NewBoard(test2x2Board()),
			args:    args{"1234123412341234"},
			wantErr: false,
		},
		{
			name:    "2x2 wrong character is 0",
			b:       NewBoard(test2x2Board()),
			args:    args{"12X4123412341234"},
			wantErr: false,
		},
		{
			name:    "2x2 wrong size",
			b:       NewBoard(test2x2Board()),
			args:    args{"1234"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.LoadFromString(tt.args.board); (err != nil) != tt.wantErr {
				t.Errorf("Board.LoadFromString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBoard_String(t *testing.T) {
	tests := []struct {
		name    string
		b       Board
		wantRes string
	}{
		{
			name:    "2x2",
			b:       test2x2BoardSolved(),
			wantRes: "1234341221434321",
		},
		{
			name:    "2x2 wrong character is 0",
			b:       NewBoard(test2x2Board()),
			wantRes: "0000000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.b.String(); res != tt.wantRes {
				t.Errorf("Board.String() res = %v, wantRes %v", res, tt.wantRes)
			}
		})
	}
}

func TestBoard_isSolved(t *testing.T) {
	tests := []struct {
		name       string
		b          Board
		wantSolved bool
	}{
		{
			name:       "2x2",
			b:          test2x2BoardSolved(),
			wantSolved: true,
		},
		{
			name:       "2x2 wrong character is 0",
			b:          NewBoard(test2x2Board()),
			wantSolved: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.b.isSolved(); res != tt.wantSolved {
				t.Errorf("Board.isSolved() res = %v, wantRes %v", res, tt.wantSolved)
			}
		})
	}
}

func TestBoard_getValue(t *testing.T) {
	type args struct {
		pos int
	}

	tests := []struct {
		name string
		b    Board
		args args
		want int
	}{
		{
			name: "2x2",
			b:    test2x2BoardSolved(),
			args: args{4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.b.getValue(tt.args.pos); res != tt.want {
				t.Errorf("Board.getValue() res = %v, wantRes %v", res, tt.want)
			}
		})
	}
}

func TestBoard_setValue(t *testing.T) {
	type args struct {
		pos   int
		value int
	}

	tests := []struct {
		name string
		b    Board
		args args
		want int
	}{
		{
			name: "2x2",
			b:    test2x2BoardSolved(),
			args: args{4, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.setValue(tt.args.pos, tt.args.value)
			if res := tt.b.getValue(tt.args.pos); res != tt.want {
				t.Errorf("Board.setValue() res = %v, wantRes %v", res, tt.want)
			}
		})
	}
}

func TestBoard_getPotential(t *testing.T) {
	type args struct {
		pos int
	}

	tests := []struct {
		name string
		b    Board
		args args
		want []int
	}{
		{
			name: "2x2",
			b:    test2x2BoardSolved(),
			args: args{4},
			want: []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.b.getPotential(tt.args.pos); !reflect.DeepEqual(res, tt.want) {
				t.Errorf("Board.getValue() res = %v, wantRes %v", res, tt.want)
			}
		})
	}
}

func TestBoard_setPotential(t *testing.T) {
	type args struct {
		pos   int
		value []int
	}

	tests := []struct {
		name string
		b    Board
		args args
		want []int
	}{
		{
			name: "2x2",
			b:    test2x2BoardSolved(),
			args: args{4, []int{2, 3, 4}},
			want: []int{2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.setPotential(tt.args.pos, tt.args.value)
			if res := tt.b.getPotential(tt.args.pos); !reflect.DeepEqual(res, tt.want) {
				t.Errorf("Board.setValue() res = %v, wantRes %v", res, tt.want)
			}
		})
	}
}

func TestBoard_getAllNeighborsValues(t *testing.T) {
	type args struct {
		pos int
	}

	tests := []struct {
		name string
		b    Board
		args args
		want neighbors
	}{
		{
			name: "2x2",
			b:    test2x2BoardSolved(),
			args: args{1},
			want: neighbors{1, 2, 3, 4, 3, 4, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.b.getAllNeighborsValues(tt.args.pos); !reflect.DeepEqual(res, tt.want) {
				t.Errorf("Board.getAllNeighborsValues() res = %v, wantRes %v", res, tt.want)
			}
		})
	}
}

func TestBoard_Solve(t *testing.T) {
	type args struct {
		p int
	}

	tests := []struct {
		name string
		b    Board
		want bool
	}{
		{
			name: "3x3",
			b:    test3x3BoardUnsolved(),
			want: true,
		},
		{
			name: "3x3 imposible",
			b:    test3x3BoardImpossible(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.b.Solve(); res != tt.want {
				t.Errorf("Board.LoadFromString() error = %v, wantErr %v", res, tt.want)
			}
		})
	}
}

func Test_unique(t *testing.T) {
	type args struct {
		intSlice []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{[]int{1, 2}},
			want: true,
		},
		{
			args: args{[]int{1, 0, 0, 2}},
			want: true,
		},
		{
			args: args{[]int{1, 3, 0, 2, 2}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unique(tt.args.intSlice); got != tt.want {
				t.Errorf("unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_isValid(t *testing.T) {
	tests := []struct {
		name      string
		b         Board
		wantValid bool
	}{
		{
			name:      "2x2",
			b:         test2x2BoardSolved(),
			wantValid: true,
		},
		{
			name:      "2x2",
			b:         test2x2BoardInvalidFlat(),
			wantValid: false,
		},
		{
			name:      "2x2",
			b:         test2x2BoardInvalidY(),
			wantValid: false,
		},
		{
			name:      "2x2",
			b:         test2x2BoardInvalidX(),
			wantValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.b.isValid(); res != tt.wantValid {
				t.Errorf("Board.isSolved() res = %v, wantRes %v", res, tt.wantValid)
			}
		})
	}
}

func TestBoard_NicePrint(t *testing.T) {
	tests := []struct {
		name string
		b    Board
		res  string
	}{
		{
			name: "2x2",
			b:    test2x2BoardSolved(),
			res:  "╔═══╤═══╦═══╤═══╗\n║ 1 │ 2 ║ 3 │ 4 ║\n╟───┼───╫───┼───╢\n║ 3 │ 4 ║ 1 │ 2 ║\n╠═══╪═══╬═══╪═══╣\n║ 2 │ 1 ║ 4 │ 3 ║\n╟───┼───╫───┼───╢\n║ 4 │ 3 ║ 2 │ 1 ║\n╚═══╧═══╩═══╧═══╝\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.b.NicePrint(); res != tt.res {
				t.Errorf("Board.NicePrint() res = %v, res %v", res, tt.res)
			}
		})
	}
}
