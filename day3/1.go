package main

import (
    "fmt"
    "math"
)


func ComputeCoor(target int) (int, int) {
    ctr, x, y := 9, 1, -1
    direction := ""

    for ctr != target {
        // compute direction if needs changing
        if (x < 0 && y < 0) && x == y {
            direction = "right"
        } else if y < 0 && x == (y*-1) {
            direction = "right"
        } else if x == y {
            direction = "left"
        } else if  x < 0 && (x*-1) == y {
            direction = "down"
        } else if (x > 0 && y < 0) && (x + y == 1) {
            direction = "up"
        }

        // increment coor wrt to direction
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

        // keep going
        ctr++;
    }

    return x, y
}

func main() {
    x, y := ComputeCoor(277678)
    fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}
