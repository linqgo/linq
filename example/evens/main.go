package main

import (
	"fmt"

	"github.com/linqgo/linq/v2"
)

func main() {
	even := func(i int) bool { return i%2 == 0 }
	for i := range linq.From(1, 2, 3, 4, 5).Where(even).Seq() {
		fmt.Println(i)
	}
}
