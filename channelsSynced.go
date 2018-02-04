package main

import (
     "fmt"
     "time"
)

func msgPrinter(ch chan int){
    for {
        num := <- ch
        fmt.Println(num)
    }
}

func main() {
    ch := make(chan int)
    go msgPrinter(ch)

    for i:=0; ;i++ {
        ch <- i
        time.Sleep(time.Millisecond * 10)
    }

    <- ch
}
