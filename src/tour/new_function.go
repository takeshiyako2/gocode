package main

import (
       "fmt"
)

type Vertes struct {
     X, Y int
}

func main() {
     v := new(Vertes)
     fmt.Println(v.X)
     fmt.Println(v.Y)
     v.X, v.Y = 11, 9
     fmt.Println(v.X)
     fmt.Println(v.Y)
}

