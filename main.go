package main

import (
	"fmt"
	"log"
)

const MAX_MOVES = 9

type row = [3]string
type board = [3]row
type input = [2]uint8

func get_input () input {
    var result input
    
    fmt.Println("Enter the x and y coordinates")
    _, err := fmt.Scanln(&result[0], &result[1])

    if err != nil {
        log.Fatalln("[ERROR]:", err)
    }

    result[0] -= 1
    result[1] -= 1

    return result
}

func win(b board, move_num int) bool {
    log.Fatalln("win() not implemented")
    return false
}

func main() { 
    board := board{
        {" ", " ", " "}, 
        {" ", " ", " "}, 
        {" ", " ", " "},
    }

    move_number := 0
    var p1_input input
    var p2_input input
    
    
    for move_number < MAX_MOVES || win(board, move_number) {
        p1_input = get_input()
        p2_input = get_input()
        move_number += 2 
        fmt.Println(move_number)
    }
}

