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
	// check_...
}

// ----------------------------------------------------------------------

func check_createArray() {
	// only-declare (not initialize) array
	var a1 [3]int
	// значения автоматически инициализируются дефолтным значением для этого типа массива (например для типа int: default value = 0).
	fmt.Printf("a1 = %v\n", a1)           // [0 0 0]
	fmt.Printf("len(a1) = %d\n", len(a1)) // 3
	fmt.Printf("cap(a1) = %d\n", cap(a1)) // 3

	// declare and initialize (empty) array
	a2 := [3]int{}
	fmt.Printf("a2 = %v\n", a2)           // [0 0 0]
	fmt.Printf("len(a2) = %d\n", len(a2)) // 3
	fmt.Printf("cap(a2) = %d\n", cap(a2)) // 3

	// declare and initialize (not empty) array
	a3 := [3]int{1, 2, 3}
	fmt.Printf("a3 = %v\n", a3)           // [1 2 3]
	fmt.Printf("len(a3) = %d\n", len(a3)) // 3
	fmt.Printf("cap(a3) = %d\n", cap(a3)) // 3

	// declare and initialize (not empty) array
	a4 := [...]int{1, 2, 3}
	fmt.Printf("a4 = %v\n", a4)           // [1 2 3]
	fmt.Printf("len(a4) = %d\n", len(a4)) // 3
	fmt.Printf("cap(a4) = %d\n", cap(a4)) // 3
}

// ----------------------------------------------------------------------

func check_array() {
    arr := [3]int{10, 20, 30}
    
    fmt.Printf("Array size: %d bytes\n", unsafe.Sizeof(arr))
    fmt.Printf("Element size: %d bytes\n", unsafe.Sizeof(arr[0]))
    
    // Memory addresses are contiguous
    for i := 0; i < len(arr); i++ {
        fmt.Printf("arr[%d] address: %p\n", i, &arr[i])
    }

/*

Array size: 24 bytes // 3 elements * 8 bytes each (on 64-bit system)
Element size: 8 bytes
arr[0] address(hex): 0xc00009e030
arr[1] address(hex): 0xc00009e038  // +8 bytes
arr[2] address(hex): 0xc00009e040  // +8 bytes 
// p.s. 16(dec.) = 10(hex.), p.s. 30(hex)+16(dec.)=30(hex)+10(hex)=40(hex)

*/

}

// ----------------------------------------------------------------------

func check_multiDimensionalArray() {
    var matrix [2][3]int
    
    // Stored as: 
	// [row0-col0][row0-col1][row0-col2]
	// [row1-col0][row1-col1][row1-col2]
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            fmt.Printf("matrix[%d][%d] = %p\n", i, j, &matrix[i][j])
        }
    }

/*

output:

matrix[0][0] = 0xc00000c330
matrix[0][1] = 0xc00000c338
matrix[0][2] = 0xc00000c340
matrix[1][0] = 0xc00000c348
matrix[1][1] = 0xc00000c350
matrix[1][2] = 0xc00000c358

*/

}

// ----------------------------------------------------------------------















