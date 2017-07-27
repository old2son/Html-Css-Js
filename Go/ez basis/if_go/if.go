package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func main() {
    fmt.Println(sqrt(2), sqrt(-4))
    // 1.4142135623730951 2i
}

/*
    if

    就像 for 循环一样, Go 的 if 语句也不要求用 () 将条件括起来, 同时, {} 还是必须有的
*/
