package main

import (
     "fmt"
     "time"
)

func echo(ch chan int){
    for {
        num := <- ch
        fmt.Println(num)
        
        // try a "select case" syntax here
    }
}

func wow(){
    for i:=0; i<=100; i++ {
        fmt.Println(i)
    }
    fmt.Println("Done normal printing")
}

func main() {
    ch := make(chan int)

    go echo(ch)

    go wow()

    for i:=101; i<=200; i++ {
        ch <- i
    }

    fmt.Println("Done channel printing")

    time.Sleep(time.Millisecond * 10000)
}
