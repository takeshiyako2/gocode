package main

import (
       "fmt"
)

func process(ch chan<- int){
     ch <- 1
     ch	<- 2
     ch	<- 3
}

func main(){
     c:= make(chan int)
     go process(c)
     fmt.Printf("out=%d %d %d\n", <-c, <-c, <-c)
}