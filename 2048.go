package main

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/nsf/termbox-go"
)

const (
    boardSize = 4
)

var (
    board [boardSize][boardSize]int
    gameOver bool
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    initBoard()
    for {
        drawBoard()
        handleInput()
        if gameOver {
            fmt.Println("Game Over!")
            break
        }
    }
}

func initBoard() {
 // Initialize the board with two random "2" or "4" values
    addRandomTile()
    addRandomTile()
}

func addRandomTile() {
    // Choose a random empty cell on the board
    var emptyCells []struct{x, y int}
    for i := 0; i < boardSize; i++ {
        for j := 0; j < boardSize; j++ {
            if board[i][j] == 0 {
                emptyCells = append(emptyCells, struct{x, y int}{i, j})
            }
        }
    }
}

    if len(emptyCells) == 0 {
        gameOver = true
        return
    }

    // Add a "2" or "4" to the empty cell
    randIndex := rand.Intn(len(emptyCells))
    randomCell := emptyCells[randIndex]
    if rand.Float32() < 0.9 {
        board[randomCell.x][randomCell.y] = 2
    } else {
        board[randomCell.x][randomCell.y] = 4
    }}

func handleInput() {
    var input string
    fmt.Scan(&input)
    switch input {
    case "u":
        moveTiles("up")
    case "d":
        moveTiles("down")
    case "l":
        moveTiles("left")
    case "r":
        moveTiles("right")
    default:
        fmt.Println("Invalid input, please use 'w', 'a', 's', 'd' keys.")
    }
}


func moveTiles() {
    var moved bool

    switch direction {
    case "up":
        moved = moveUp()
    case "down":
        moved = moveDown()
    case "left":
        moved = moveLeft()
    case "right":
        moved = moveRight()
    }

    if moved {
        addRandomTile()
    }
}
func moveUp() bool {
    var moved bool

    // Move tiles up
    for j := 0; j < boardSize; j++ {
        for i := 1; i < boardSize; i++ {
            if board[i][j] == 0 {
                continue
            }
            for k := i; k > 0; k-- {
                if board[k-1][j] == 0 {
                    board[k-1][j] = board[k][j]
                    board[k][j] = 0
                    moved = true
                }
            }
        }
    }

    // Merge equal tiles
    for i := 1; i < boardSize; i++ {
        for j := 0; j < boardSize; j++ {
            if board[i][j] != 0 && board[i][j] == board[i-1][j] {
                board[i-1][j] += board[i][j]
                board[i][j] = 0
                moved = true
            }
        }
    }

    // Move tiles up again
    for j := 0; j < boardSize; j++ {
        for i := 1; i < boardSize; i++ {
            if board[i][j] == 0 {
                continue
            }
            for k := i; k > 0; k-- {
                if board[k-1][j] == 0 {
                    board[k-1][j] = board[k][j]
                    board[k][j] = 0
                    moved = true
                }
            }
        }
    }

    return moved
}

func moveDown() bool {
    var moved bool

    // Move tiles down
    for j := 0; j < boardSize; j++ {
        for i := boardSize - 2; i >= 0; i-- {
            if board[i][j] == 0 {
                continue
            }
            for k := i; k < boardSize-1; k++ {
                if board[k+1][j] == 0 {
                    board[k+1][j] = board[k][j]
                    board[k][j] = 0
                    moved = true
                }
            }
        }
    }

    // Merge equal tiles
    for i := boardSize - 2; i >= 0; i-- {
        for j := 0; j < boardSize; j++ {
            if board[i][j] != 0 && board[i][j] == board[i+1][j] {
                board[i+1][j] += board[i][j]
                board[i][j] = 0
                moved = true
            }
        }
    }

    // Move tiles down again
    for j := 0; j < boardSize; j++ {
        for i := boardSize - 2; i >= 0; i-- {
            if board[i][j] == 0 {
                continue
            }
            for k := i; k < boardSize-1; k++ {
                if board[k+1][j] == 0 {
                    board[k+1][j] = board[k][j]
                    board[k][j] = 0
                    moved = true
                }
            }
        }
    }

    return moved
}
func moveTilesLeft(grid [][]int) [][]int {
    // Iterate through each row
    for i := 0; i < len(grid); i++ {
        // Iterate through each column starting from the second column
        for j := 1; j < len(grid[i]); j++ {
            // If the current tile is not empty
            if grid[i][j] != 0 {
                // Iterate through each column to the left of the current column
                for k := 0; k < j; k++ {
                    // If the current tile is empty, move the tile to the left
                    if grid[i][k] == 0 {
                        grid[i][k] = grid[i][j]
                        grid[i][j] = 0
                        break
                    }
                    // If the current tile is equal to the tile to the left, merge them
                    if grid[i][k] == grid[i][j] {
                        grid[i][k] *= 2
                        grid[i][j] = 0
                        break
                    }
                }
            }
        }
    }
    // Move the tiles left again after the merge
    grid = moveTilesLeft(grid)
    return grid
}

func moveTilesRight(grid [][]int) (bool, [][]int) {
    // variable to keep track if move was made
    moveMade := false
    // Iterate through each row
    for i := 0; i < len(grid); i++ {
        // Iterate through each column starting from the second to last column
        for j := len(grid[i]) - 2; j >= 0; j-- {
            // If the current tile is not empty
            if grid[i][j] != 0 {
                // Iterate through each column to the right of the current column
                for k := len(grid[i]) - 1; k > j; k-- {
                    // If the current tile is empty, move the tile to the right
                    if grid[i][k] == 0 {
                        grid[i][k] = grid[i][j]
                        grid[i][j] = 0
                        moveMade = true
                        break
                    }
                    // If the current tile is equal to the tile to the right, merge them
                    if grid[i][k] == grid[i][j] {
                        grid[i][k] *= 2
                        grid[i][j] = 0
                        moveMade = true
                        break
                    }
                }
            }
        }
    }
    return moveMade, grid
}

func isGameOver(grid [][]int) bool {
    // Check if the board is full
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 0 {
                return false
            }
        }
    }
    // Check if any moves are possible
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if i > 0 && grid[i-1][j] == grid[i][j] {
                return false
            }
            if i < len(grid)-1 && grid[i+1][j] == grid[i][j] {
                return false
            }
            if j > 0 && grid[i][j-1] == grid[i][j] {
                return false
            }
            if j < len(grid[i])-1 && grid[i][j+1] == grid[i][j] {
                return false
            }
        }
    }
    return true
}

func drawBoard(grid [][]int) {
    // Initialize the termbox library
    err := termbox.Init()
    if err != nil {
        panic(err)
    }
    defer termbox.Close()

    // Iterate through each row
    for i := 0; i < len(grid); i++ {
        // Iterate through each column
        for j := 0; j < len(grid[i]); j++ {
            // Get the value of the current tile
            value := grid[i][j]
            // Set the color of the tile based on its value
            var fg, bg termbox.Attribute
            switch {
            case value == 0:
                fg, bg = termbox.ColorDefault, termbox.ColorDefault
            case value < 4:
                fg, bg = termbox.ColorWhite, termbox.ColorBlack
            case value < 8:
                fg, bg = termbox.ColorYellow, termbox.ColorBlack
            case value < 16:
                fg, bg = termbox.ColorGreen, termbox.ColorBlack
            case value < 32:
                fg, bg = termbox.ColorCyan, termbox.ColorBlack
            case value < 64:
                fg, bg = termbox.ColorMagenta, termbox.ColorBlack
            case value < 128:
                fg, bg = termbox.ColorBlue, termbox.ColorBlack
            case value < 256:
                fg, bg = termbox.ColorRed, termbox.ColorBlack
            default:
                fg, bg = termbox.ColorBlack, termbox.ColorWhite
            }

            // Draw the current tile
            if value != 0 {
                termbox.SetCell(j, i, ' ', fg, bg)
                fmt.Fprintf(termbox.CellBuffer(), "%4d", value)
            } else {
                termbox.SetCell(j, i, ' ', termbox.ColorDefault, termbox.ColorDefault)
            }
        }
    }
    termbox.Flush()
}
