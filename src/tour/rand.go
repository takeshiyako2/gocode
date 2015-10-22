package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // シードを与えないで実行
    for i := 0; i < 10; i++ {
        fmt.Print(rand.Intn(10), " ")
    }
    fmt.Println(" ")

    // randをNewSourceで初期化して実行
    my_rand := rand.New(rand.NewSource(1))
    for i := 0; i < 10; i++ {
        fmt.Print(my_rand.Intn(10), " ")
    }
    fmt.Println(" ")

    // シードを与えて実行
    my_rand.Seed(time.Now().UnixNano())
    for i := 0; i < 10; i++ {
        fmt.Print(my_rand.Intn(10), " ")
    }
}
