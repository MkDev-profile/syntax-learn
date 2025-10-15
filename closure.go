// closure (замыкание)

package main

import (
	"fmt"
	"time"
)

func Main_closure() {
	values := []int{1, 2, 3, 4, 5}

	delayDur := time.Duration(len(values)) * time.Second

	//for val := 1; val < 6; val++ {
	for _, val := range values {
		go func() {
			time.Sleep(delayDur)
			fmt.Println(val)
		}()
	}
	fmt.Println("loop completed")

	time.Sleep(delayDur*2)
}
/*

// output:

loop completed
2
5
1
4
3

*/











