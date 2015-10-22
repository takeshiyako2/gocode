package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Intnメソッドで擬似乱数を出力
	// シードを与えていないので、実行毎に同じ数字が出力される
	// rand.Seed(1)と同じ
	fmt.Println("no seed ->")
	for i := 0; i < 10; i++ {
		fmt.Print(rand.Intn(100), " ")
	}
	fmt.Println("")

	// シードを与えて、擬似乱数を出力
	fmt.Println("use seed ->")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Print(rand.Intn(100), " ")
	}
}
