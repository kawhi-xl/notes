package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	t1 := time.Since(start).Nanoseconds()
	fmt.Println(s, t1)

	start = time.Now()
	s = strings.Join(os.Args[1:], " ")
	t2 := time.Since(start).Nanoseconds()
	fmt.Println(s, t2)
}
