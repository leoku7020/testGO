package main

import (
    "fmt"
)
//defined Channel
// ch := make(chan int)
//Send
// ch <- 5
//Receive
// x := <- ch
//Closr Channel
// close(ch)
func main() {
    c := make(chan int)
    // use go routine
    go func() {
        for i := 0; i < 10; i++ {
            c <- i
        }
        close(c)
    }()

    for n := range c {
        fmt.Println(n)
    }
}
