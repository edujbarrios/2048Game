package main

import (
    "fmt"
    "math/rand"
    "time"
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


func moveLeft() bool {
    // Move tiles left and merge if possible
    // ...
    return true // or false if no move was made
}

func moveRight() bool {
    // Move tiles right and merge if possible
    // ...
    return true // or false if no move was made
}


func checkGameOver() {
    // Check if the game is over (i.e. the board is full and no more moves are possible)
    // ...
}

func drawBoard() {
    // Draw the current state of the board to the terminal
    // ...
}
