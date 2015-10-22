package main

import (
	"fmt"
	"time"
	"math/rand"
)

func test(n int, name string) {
	for i := 1; i <= n; i++ {
		fmt.Println(i, name)
		time.Sleep(500 * time.Millisecond)
		fmt.Println(rand.Intn(100), " ")
	}
}

func main() {
	test(5, "foo")
	test(5, "bar")
}
