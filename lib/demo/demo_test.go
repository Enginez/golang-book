package demo

import "testing"

func TestPinger(t *testing.T) {
    var c chan string = make(chan string)
    go Pinger(c)
    for {
        msg := <- c
        if msg != "ping" {
            t.Error("Expected PING, got ", msg)
        }else{
            break;
        }
    }
}