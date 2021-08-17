package main

import (
  "fmt"
  "strings"
  "math/rand"
  "time"
  //"strconv"
)

func main() {
  // initialize boards
  board0 := init_board(40, 40)
  fill_board(board0)
  board1 := init_board(40, 40)

  // game loop
  for {
    clearscreen()
    print_board(board0)
    apply_rule(board0, board1)
    copy_board(board1, board0)
    time.Sleep(1 * time.Second)
  }
}

/*
Helper and utility functions
*/
func init_board(x,y int) [][]string {
  board := make([][]string, x)
  for i := 0; i < len(board); i++ {
    board[i] = make([]string, y)
  }
  for i := 0; i < x; i++ {
    for j := 0; j < y; j++ {
      board[i][j] = "O"
    }
  }
  return board
}

func fill_board(board [][]string) {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  for i := 0; i < len(board); i++ {
    for j := 0; j < len(board[i]); j++ {
      if r.Intn(2) == 1 {
        board[i][j] = "\033[31mX\033[0m"
      } else {
        board[i][j] = "O"
      }
    }
  }
}
func print_board(board [][]string) {
  for i := 0; i < len(board); i++ {
    fmt.Printf("%s\n", strings.Join(board[i], " "))
  }
}

func copy_board(src, dest [][]string) {
  for i := 0; i < len(src); i++ {
    for j := 0; j < len(src[i]); j++ {
      dest[i][j] = src[i][j]
    }
  }
}

func apply_rule(src, dest [][]string) {
  for i := 0; i < len(src); i++ {
    for j :=0; j < len(src[0]); j++ {
      /*
      if neighboring_xs(i, j, src) > 4 {
        dest[i][j] = "O"
      } else {
        dest[i][j] = src[i][j]
      }
      */
      switch n := neighboring_xs(i, j, src); n {
      case 2, 3:
        if src[i][j] == "O" {
          dest[i][j] = "\033[31mX\033[0m"
        }
      default:
        dest[i][j] = "O"
      }
    }
  }
}

func neighboring_xs(x,y int, board [][]string) int {
  neighbors := 0
  height := len(board)
  width := len(board[0])

  for i := -1; i < 2; i++ {
    for j := -1; j < 2; j++ {
      if x + i < 0 || x + i > height -1 { // stay inside array
        continue
      } else if y + j < 0 || y + j > width -1 { // stay inside array
        continue
      } else if i == 0 && j == 0 { // don't count yourself
        continue
      } else {
        if board[x + i][y + j] == "\033[31mX\033[0m" {
          neighbors++
        }
      }
    }
  }
  return neighbors
}

func clearscreen() {
  fmt.Print("\033[2J\033[H")
}