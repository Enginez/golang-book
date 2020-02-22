package main

import "fmt"

func a(n int) {
    go b(n-1)
    for i := 0; i < n; i++ {
        fmt.Println(n, ":", i)
        
        
    }
}

func b(n int) {
    go a(n-1)
    for i := 0; i < n; i++ {
        fmt.Println(n, ":", i)
  
    }
}


func main() {
    go a(5)
    var input string
    fmt.Scanln(&input)
}