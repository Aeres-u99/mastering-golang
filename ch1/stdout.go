package main
import (
    "os"
    "io"
)

func main() {
    myString := ""
    arguments := os.Args
    if len(arguments) == 1 {
        myString = "Specify an Argument atleast!"
    } else {
        myString = arguments[1]
    }
    io.WriteString(os.Stdout, myString)
    io.WriteString(os.Stdout, ": ~Pew")
    io.WriteString(os.Stdout, "\n")
}
