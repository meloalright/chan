package main

import "time"
import "fmt"

func main() {
    timer1 := time.NewTimer(time.Second * 2)
    <-timer1.C
    fmt.Println("Timer 1 expired")
    //
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
    //
    pending := make(chan int)
    ticker := time.NewTicker(time.Millisecond * 500)

    go func() {
        pending <- 1
    }()

    go func() {
        fmt.Println("Tick start")
        for t := range ticker.C {
            pending <- 1
            fmt.Println("Tick at", t)
        }
    }()
    for i := range pending {
        fmt.Println("Get from pending", i)
    }
}