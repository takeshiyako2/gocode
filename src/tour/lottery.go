package main

import (
	"fmt"
	"github.com/kyokomi/lottery"
	"math/rand"
	"time"
)

func main() {
	lot := lottery.New(rand.New(rand.NewSource(time.Now().UnixNano())))

	check := 1000
	count := 0

	for i := 0; i < check; i++ {

		if lot.LotOf(1, 5) {
			// 1/5の確率
			fmt.Print("アタリ")
			count++
		} else {
			fmt.Print("ハズレ")
		}

	}

	fmt.Println(float32(count)/float32(check)*100, "%")

}
