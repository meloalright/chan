package main

import "fmt"
import "time"

func main() {
    c := make(chan int)
    go func() {
        c <- 11
        time.Sleep(1 * time.Second)
        c <- 12
        close(c)
    }()
    go func() {
        for i := 0; i < 10; i = i + 1 {
            c <- i
        }
    }()
    for i := range c {
        fmt.Println(i)
    }
    fmt.Println("Finished")
}