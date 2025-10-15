package main

import (
	"fmt"
	"unsafe"
)

type Point struct {
	X, Y int
}

func Main_struct() {
	points := [2]Point{
		{1, 2},
		{3, 4},
	}

	// stores as: [X1][Y1][X2][Y2] (contiguous)
	fmt.Printf("Point size: %d\n", unsafe.Sizeof(Point{})) // 16
	fmt.Printf("Array size: %d\n", unsafe.Sizeof(points))  // 32
}











