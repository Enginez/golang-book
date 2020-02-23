package demo

import(
    "fmt"
    "time"
)
func Pinger(c chan string) {
    for i := 0; ; i++ {
        c <- "ping"
    }
}

func Ponger(c chan string) {
    for i := 0; ; i++ {
        c <- "pong"
    }
}
func Printer(c chan string) {
    for {
        msg := <- c
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
}