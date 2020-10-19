# Sodogo

A Sudoku solver written in [Go](https://golang.org/).

From [Wikipedia](https://en.wikipedia.org/wiki/Sudoku):
> Sudoku (数独 sūdoku, digit-single) (/suːˈdoʊkuː/, /-ˈdɒk-/, /sə-/, originally
> called Number Place) is a logic-based, combinatorial number-placement puzzle.
> The objective is to fill a 9×9 grid with digits so that each column, each row,
> and each of the nine 3×3 subgrids that compose the grid (also called "boxes",
> "blocks", or "regions") contains all of the digits from 1 to 9. The puzzle
> setter provides a partially completed grid, which for a well-posed puzzle has
> a single solution.


## Example

```go
package main

import (
    "fmt"
    "os"
    "github.com/rfiestas/sodogo"
)

func main() {
    helper := sodogo.NewHelperBoard(3) // Creates a 3x3 board helpers
    board := sodogo.NewBoard(helper)   // Creates an empty 3x3 board
    
    // Load sodoku from string
    err := board.LoadFromString("004300209005009001070060043006002087190007400050083000600000105003508690042910300")
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
    
    fmt.Println("Sodoku to solve:")
    fmt.Println(board.NicePrint())
    
    fmt.Println("Go,...")
    if board.Solve() {
        fmt.Printf("Solved in %d steps, elapsed time: %s\n", board.Steps, board.Elapsed)
        fmt.Println(board.NicePrint())
    } else {
        fmt.Println("Impossible to solve :(")
    }
}
```

```bash
$ go run .

Sodoku to solve:
╔═══╤═══╤═══╦═══╤═══╤═══╦═══╤═══╤═══╗
║   │   │ 4 ║ 3 │   │   ║ 2 │   │ 9 ║
╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
║   │   │ 5 ║   │   │ 9 ║   │   │ 1 ║
╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
║   │ 7 │   ║   │ 6 │   ║   │ 4 │ 3 ║
╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣
║   │   │ 6 ║   │   │ 2 ║   │ 8 │ 7 ║
╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
║ 1 │ 9 │   ║   │   │ 7 ║ 4 │   │   ║
╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
║   │ 5 │   ║   │ 8 │ 3 ║   │   │   ║
╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣
║ 6 │   │   ║   │   │   ║ 1 │   │ 5 ║
╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
║   │   │ 3 ║ 5 │   │ 8 ║ 6 │ 9 │   ║
╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
║   │ 4 │ 2 ║ 9 │ 1 │   ║ 3 │   │   ║
╚═══╧═══╧═══╩═══╧═══╧═══╩═══╧═══╧═══╝

Go,...
Solved in 4 steps, elapsed time: 189.21µs
╔═══╤═══╤═══╦═══╤═══╤═══╦═══╤═══╤═══╗
║ 8 │ 6 │ 4 ║ 3 │ 7 │ 1 ║ 2 │ 5 │ 9 ║
╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
║ 3 │ 2 │ 5 ║ 8 │ 4 │ 9 ║ 7 │ 6 │ 1 ║
╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
║ 9 │ 7 │ 1 ║ 2 │ 6 │ 5 ║ 8 │ 4 │ 3 ║
╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣
║ 4 │ 3 │ 6 ║ 1 │ 9 │ 2 ║ 5 │ 8 │ 7 ║
╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
║ 1 │ 9 │ 8 ║ 6 │ 5 │ 7 ║ 4 │ 3 │ 2 ║
╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
║ 2 │ 5 │ 7 ║ 4 │ 8 │ 3 ║ 9 │ 1 │ 6 ║
╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣
║ 6 │ 8 │ 9 ║ 7 │ 3 │ 4 ║ 1 │ 2 │ 5 ║
╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
║ 7 │ 1 │ 3 ║ 5 │ 2 │ 8 ║ 6 │ 9 │ 4 ║
╟───┼───┼───╫───┼───┼───╫───┼───┼───╢
║ 5 │ 4 │ 2 ║ 9 │ 1 │ 6 ║ 3 │ 7 │ 8 ║
╚═══╧═══╧═══╩═══╧═══╧═══╩═══╧═══╧═══╝
```
