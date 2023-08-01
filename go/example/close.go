package main

import "fmt"
import "time"

func main() {
    go func() {
        time.Sleep(time.Hour)
    }()
    c := make(chan int, 10)
    c <- 10086
    c <- 1
    c <- 2
    close(c)
    i, ok := <-c
    fmt.Printf("%d, %t\n", i, ok) //10086, false
    fmt.Println(<-c) // 1
    for i := range c {
        fmt.Println(i) // 2
    }
    fmt.Println(<-c) // 0
    fmt.Println(<-c) // 0
    fmt.Println(<-c) // 0
    fmt.Println(<-c) // 0
    fmt.Println(<-c) // 0
    fmt.Println(<-c) // 0
    i, ok = <-c
    fmt.Printf("%d, %t\n", i, ok) //0, false
    i, ok = <-c
    fmt.Printf("%d, %t\n", i, ok) //0, false
    i, ok = <-c
    fmt.Printf("%d, %t\n", i, ok) //0, false
    i, ok = <-c
    fmt.Printf("%d, %t\n", i, ok) //0, false
    i, ok = <-c
    fmt.Printf("%d, %t\n", i, ok) //0, false
}