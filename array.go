package main

import (
	"fmt"
	"unsafe"
)

/*

p.s.
[3]int and [5]int are different types.

*/

func Main_array() {
	check_createArray()
}

func printArray(varName string, ptr *[5]int) {
	val := *ptr

    fmt.Printf("%s: ptr = %p, Len = %d, Cap = %d, Val = %v\n", 
		varName, ptr, len(val), cap(val), val)
}

func check_createArray() {
	var a1 [5]int
	// значения автоматически инициализируются дефолтным значением для этого типа массива (например для типа int: default value = 0).
	printArray("a1", &a1)

	a2 := [5]int{}
	printArray("a2", &a2)

	a3 := [5]int{1, 2, 3, 4, 5}
	printArray("a3", &a3)

	a4 := [...]int{1, 2, 3, 4, 5}
	printArray("a4", &a4)

    fmt.Printf("ElementType: %T\n", a4[0])
	fmt.Printf("ElementSizeInBytes: %d\n", unsafe.Sizeof(a4[0]))
	fmt.Printf("ArraySizeInBytes: %d\n", unsafe.Sizeof(a4))
    
    // Memory addresses are contiguous():
    for i := 0; i < len(a4); i++ {
        fmt.Printf("arr[%d] address: %p\n", i, &a4[i])
    }

	// 2D array = матрица
	var matrix [2][3]int // [rows][columns]
    
    // Stored as: 
	// [row0-col0][row0-col1][row0-col2]
	// [row1-col0][row1-col1][row1-col2]
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            fmt.Printf("matrix[%d][%d] = %p\n", i, j, &matrix[i][j])
        }
    }
}
/*

// output:

a1: ptr = 0xc00000c300, Len = 5, Cap = 5, Val = [0 0 0 0 0]
a2: ptr = 0xc00000c360, Len = 5, Cap = 5, Val = [0 0 0 0 0]
a3: ptr = 0xc00000c3c0, Len = 5, Cap = 5, Val = [1 2 3 4 5]
a4: ptr = 0xc00000c420, Len = 5, Cap = 5, Val = [1 2 3 4 5]
ElementType: int
ElementSizeInBytes: 8
ArraySizeInBytes: 40
arr[0] address: 0xc00000c420
arr[1] address: 0xc00000c428
arr[2] address: 0xc00000c430
arr[3] address: 0xc00000c438
arr[4] address: 0xc00000c440
matrix[0][0] = 0xc00000c510
matrix[0][1] = 0xc00000c518
matrix[0][2] = 0xc00000c520
matrix[1][0] = 0xc00000c528
matrix[1][1] = 0xc00000c530
matrix[1][2] = 0xc00000c538

*/



















