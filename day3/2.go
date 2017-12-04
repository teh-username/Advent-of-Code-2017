package main

import (
    "fmt"
    "math"
)

func GetNextDirection(x int, y int) (string) {
    direction := ""
    // compute direction if needs changing
    if (x <= 0 && y <= 0) && x == y {
        direction = "right"
    } else if y <= 0 && x == (y*-1) {
        direction = "right"
    } else if x == y {
        direction = "left"
    } else if  x <= 0 && (x*-1) == y {
        direction = "down"
    } else if (x => 0 && y <= 0) && (x + y == 1) {
        direction = "up"
    }

    return direction
}

func GetNewCoor(x int, y int, direction string) (int, int) {
    switch direction {
    case "right":
        x++
    case "left":
        x--
    case "up":
        y++
    case "down":
        y--
    default:
        fmt.Println("WTF")
    }

    return x, y
}

func ComputeTotal(target int) (int, int) {
    spiral := map[[2]int]int {
        [2]int {0, 0}: 1,
        [2]int {1, 0}: 1,
        [2]int {1, 1}: 2,
        [2]int {0, 1}: 4,
        [2]int {-1, 1}: 5,
        [2]int {-1, 0}: 10,
        [2]int {-1, -1}: 11,
        [2]int {0, -1}: 23,
        [2]int {1, -1}: 25,
    }
    num, x, y := 0, 1, -1
    direction, lastDirection := "", "right"

    for num < target {
        direction := GetNextDirection(x, y)
        x, y := GetNewCoor(x, y, direction)

        // compute what needs to be placed on the thing
        nextDirection := GetNextDirection(x, y)

        // right cases
        if direction == "right" {
            num = spiral[[2]int{x-1, y}] + spiral[[2]int{x-1, y+1}]
            if nextDirection == "right" {
                num += spiral[[2]int{x, y+1}]
            }
        }

        // up cases
        if direction == "up" {
            num =
        }



        // put it in place
        spiral[[2]int{x, y}] = num

        // get last direction
        lastDirection = direction
    }

    return x, y
}

func main() {
    z := ComputeTotal(277678)
    fmt.Println(z)
}
