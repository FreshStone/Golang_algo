package main

import (
	"fmt"
	"time"
)

func main(){
	board := [][]byte{{53,56,46,46,46,50,46,54,46},{50,49,46,51,46,46,46,46,46},{46,55,46,54,46,46,46,46,46},{46,46,46,46,46,51,46,46,53},{46,51,52,46,49,46,56,55,46},{49,46,46,57,46,46,46,46,46},{46,46,46,46,46,54,46,52,46},{46,46,46,46,46,57,46,56,49},{46,52,46,49,46,46,46,50,51}}
	printboard(board)
	t := time.Now()
	solveSudoku(board)
	fmt.Println("Runtime- ", time.Since(t))
	fmt.Println()
	printboard(board)
}

func solveSudoku(board [][]byte){
	var row [9][9]bool
	var col [9][9]bool
	var box [9][9]bool
	for i, r := range board{
		for j , _ := range r{
			if board[i][j] != 46{
				row[i][board[i][j]-49] = true
				col[j][board[i][j]-49] = true
				box[(i/3) + (j/3)*3][board[i][j]-49] = true
			}
		}
	}

	var probable [9][9][]uint8 //1-9
	var v uint8
	i := 0
	j := 0
	front := true

	for ; j < 9;{
		for ; i >= 0 && i < 9;{
			if board[i][j] != 46 && len(probable[i][j]) == 0{
				if front{
					i += 1
				}else{
					i -= 1
				}
			}else{
				/*  calculate probable[i][j]/remove last val from probale[i][j], assign a val to board[i][j], update row, col, box */
				if board[i][j] == 46 {
					for k := 0; k < 9; k++{
						if !row[i][k] && !col[j][k] && !box[(i/3) + (j/3)*3][k]{
							probable[i][j] = append(probable[i][j], uint8(k+1))
						}
					}
				}else{
					probable[i][j] = probable[i][j][:len(probable[i][j])-1]
					row[i][board[i][j]-49] = false
					col[j][board[i][j]-49] = false
					box[(i/3) + (j/3)*3][board[i][j]-49] = false
					board[i][j] = 46
				}

				if len(probable[i][j]) == 0{
					front = false
					i -= 1
				}else{
					front = true
					v = probable[i][j][len(probable[i][j])-1]
					if !row[i][v-1] && !col[j][v-1] && !box[(i/3) + (j/3)*3][v-1]{
						board[i][j] = v + 48
						row[i][v-1] = true
						col[j][v-1] = true
						box[(i/3) + (j/3)*3][v-1] = true
						i += 1
					}
				}
			}
		}
		if i == 9{
			j += 1
			i = 0
		}else{ // i = -1
			j -= 1
			i = 8
		}
	}
}

func printboard(board [][]byte){
	for i, v := range board{
		fmt.Print("| ")
		for j, b := range v{
			if b == 46{
				fmt.Print(".")
			}else{
				fmt.Print(b-48)
			}
			if (j+1)%3 == 0{
				fmt.Print(" | ")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
		if (i+1)%3 == 0{
			for j := 0; j < 13; j++{
				fmt.Print("- ")
			}
			fmt.Println()
		}
	}
}
