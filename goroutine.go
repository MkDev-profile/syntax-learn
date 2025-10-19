package main

import (
	"fmt"
	"runtime"
)

func Main_examples() {

}

func check_example_1() {
	runtime.GOMAXPROCS(1) // 1 Processor(P) => 1 OsThread => concurrency of goroutines (чередование goroutines on this same underlying thread), (not parallelism).

	done := false

	go func() { // "concurrent" goroutine 1 (race condition on var "done")
		done = true
	}()

	for !done {
	} // "concurrent" goroutine "main" (race condition on var "done")

	fmt.Println("finished")
}

// output:
// finished
