package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    var sum int64
    var avg int64
    sum = 0
    avg = 0
    arguments := os.Args
    if len(os.Args) < 2 {
        fmt.Println("Please provide numbers in format: [s|a] n1 n2 n3 ... ")
        os.Exit(1)
    }
    if (arguments[1] == "s") {
        fmt.Println("Entering Sum mode ")
        for i := 2 ; i < len(os.Args); i++ {
            n, _ := strconv.ParseInt(arguments[i], 10, 0)
            sum += n
        }
        fmt.Printf("Sum: %d", sum)
    } else if (arguments[1] == "a") {
        fmt.Println("Entering Average mode")
        for i :=2 ; i < len(os.Args); i++ {
            n, _ := strconv.ParseInt(arguments[i], 10, 32)
            sum += n
        }
        avg = sum / 2
        fmt.Printf("Average: %d", avg)
    } else {
        fmt.Println("Only options s and a are allowed")
        os.Exit(1)
    }
}

