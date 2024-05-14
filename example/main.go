package main

import (
	"fmt"

	"github.com/franklange/go-ringbuf"
)

func main() {
	r := ringbuf.NewRingBuf[int](3)

	r.Put(1, 2)
	fmt.Println(r.Raw()) // [1 2]

	r.Put(3, 4)
	fmt.Println(r.Raw()) // [4 2 3]

	// iterator maintains insertion order
	iter := r.Iter()
	for iter.Next() {
		fmt.Print(*iter.Get(), " ") // 2 3 4
	}
	fmt.Println()
}
