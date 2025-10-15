package main

import (
	"cmp"
	"fmt"
	"slices"
	"unsafe"
)

func Main_slice() {
	
}

// internal representation for type "slice" (p.s. golang source code: unsafeheader.go)
type SliceInternalStruct struct {
	// pointer to the starting element of underlying array.
	Data unsafe.Pointer

	// number of elements in this slice (slice = segment of underlying array)
	Len  int

	// the remaining elements from starting element of slice to the end of the underlying array.
	Cap  int 
}

func printSlice(varName string, ptr *[]int) {
	val := *ptr
    internalStruct := (*SliceInternalStruct)(unsafe.Pointer(ptr))
 
	// len(val) // equals to: internalStruct.Len
	// cap(val) // equals to: internalStruct.Cap

    // underlArrPtr := unsafe.SliceData(val) // equals to: &val[0].
	// printf("%p",underlArrPtr) equals to: internalStruct.Data.

    fmt.Printf("%s: Val = %s, Len = %d, Cap = %d, isNil = %t, ptr = %p, UnderlArrayPtr = %#x\n", 
		varName, val, internalStruct.Len, internalStruct.Cap, 
		val == nil, ptr, internalStruct.Data)
}

func slice_check_1() {
	// only-declare (not initialize) slice
	var s []int
	printSlice("s", &s)	
	
	s = append(s, 8) // append value "8"
	// p.s. создался новый underlying array т.к. capacity было недостаточно
	printSlice("s", &s) 

/*
s: [], len = 0, cap = 0, isNil = true
s: [8], len = 1, cap = 1, isNil = false
*/

}

func slice_check_3() {
	// only-declare (not initialize) slice
	s := make([]int, 3, 4)
	printSlice("s", &s)
	
	fmt.Println("append nothing")
	s = append(s)
	printSlice("s", &s)
	
	fmt.Println("append 8, not assign")
	_ = append(s, 8) // append value "8"
	//printSlice("s1", &s1)
	printSlice("s", &s)
	
	fmt.Println("append 9, assign to other slice")
	s1 := append(s, 9)
	printSlice("s1", &s1) 
	printSlice("s", &s) 

	fmt.Println("modify other slice")
	s1[0] = 5
	printSlice("s1", &s1) 
	printSlice("s", &s) 

	fmt.Println("append 10, assign to initial slice")
	s = append(s, 10)
	printSlice("s", &s)
	printSlice("s1", &s1)

	fmt.Println("append 11, capacity already full")
	s = append(s, 11)
	printSlice("s", &s)

	fmt.Println("modify other slice")
	s1[0] = 6
	printSlice("s1", &s1) 
	printSlice("s", &s) 

	fmt.Println("append multiple values")
	s = append(s, 12, 13, 14)
	printSlice("s", &s)

	fmt.Println("append all values from other slice")
	s2 := []int{15,16}
	s = append(s, s2...)
	printSlice("s", &s)


/*
go run .
s: [0 0 0], len = 3, cap = 4, isNil = false, addr0 = 0xc000098040
append nothing
s: [0 0 0], len = 3, cap = 4, isNil = false, addr0 = 0xc000098040
append 8, not assign
s: [0 0 0], len = 3, cap = 4, isNil = false, addr0 = 0xc000098040
append 9, assign to other slice
s1: [0 0 0 9], len = 4, cap = 4, isNil = false, addr0 = 0xc000098040
s: [0 0 0], len = 3, cap = 4, isNil = false, addr0 = 0xc000098040
modify other slice
s1: [5 0 0 9], len = 4, cap = 4, isNil = false, addr0 = 0xc000098040
s: [5 0 0], len = 3, cap = 4, isNil = false, addr0 = 0xc000098040
append 10, assign to initial slice
s: [5 0 0 10], len = 4, cap = 4, isNil = false, addr0 = 0xc000098040
s1: [5 0 0 10], len = 4, cap = 4, isNil = false, addr0 = 0xc000098040
append 11, capacity already full
s: [5 0 0 10 11], len = 5, cap = 8, isNil = false, addr0 = 0xc0000a40c0
modify other slice
s1: [6 0 0 10], len = 4, cap = 4, isNil = false, addr0 = 0xc000098040
s: [5 0 0 10 11], len = 5, cap = 8, isNil = false, addr0 = 0xc0000a40c0
append multiple values
s: [5 0 0 10 11 12 13 14], len = 8, cap = 8, isNil = false, addr0 = 0xc0000a40c0
append all values from other slice
s: [5 0 0 10 11 12 13 14 15 16], len = 10, cap = 16, isNil = false, addr0 = 0xc0000a2200
*/

}

func slice_check_2() {
	// declare and initialize (empty) slice
	s2 := make([]int, 0)                     // set length=0
	printSlice("s2", &s2)

	// declare and initialize (not empty) slice
	s3 := make([]int, 3)                  // set length=3
	printSlice("s3", &s3)

	// declare and initialize slice with capacity
	s7 := make([]int, 3, 20)                  // set length=3, capacity=20
	printSlice("s7", &s7)
	// fmt.Println(s7[4]) // panic: runtime error: index out of range

	// declare and initialize (not empty) slice with setting values
	s4 := []int{4, 5, 6}
	printSlice("s4", &s4)

	// create slice From array
	a1 := [5]int{0,1,2,3,4}
	fmt.Printf("len(a1) = %d\n", len(a1)) // 5
	fmt.Printf("cap(a1) = %d\n", cap(a1)) // 5

	s5 := a1[:]
	printSlice("s5", &s5)

	s6 := a1[2:4] // = [i2;i4) = from i2 (including) to i4 (not including)
	printSlice("s6", &s6)

/*

s2: [], len = 0, cap = 0, isNil = false, addr0 = nil
s3: [0 0 0], len = 3, cap = 3, isNil = false, addr0 = 0xc000014150
s7: [0 0 0], len = 3, cap = 20, isNil = false, addr0 = 0xc000018140
s4: [4 5 6], len = 3, cap = 3, isNil = false, addr0 = 0xc000014168
len(a1) = 5
cap(a1) = 5
s5: [0 1 2 3 4], len = 5, cap = 5, isNil = false, addr0 = 0xc00000c330
s6: [2 3], len = 2, cap = 3, isNil = false, addr0 = 0xc00000c340

*/

}

func compareArraySlice() {
    // Array - fixed size, value type
    arr := [4]int{1, 2, 3, 4}
    
    // Slice = reference to underlying array
    slice := []int{1, 2, 3, 4, 5, 6}

	slice_from_arr := arr[:]
    
    fmt.Printf("Array: %v, size: %d\n", arr, unsafe.Sizeof(arr))
    fmt.Printf("Slice: %v, size: %d\n", slice, unsafe.Sizeof(slice))
    fmt.Printf("SliceFromArr: %v, size: %d\n", slice_from_arr, unsafe.Sizeof(slice_from_arr))
}

/*

output:
Array: [3]int, size: 32 // 4 int elements
Slice: []int, size: 24  // Slice header (ptr, len, cap) // p.s. 3 int fields
SliceFromArr: []int, size: 24 // аналогично: 3 int fields

*/


// -- SortFunc

type Person struct {
    Name string
    Age  int
}

func example_for_sort_func() {
    people := []Person{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }

    // Sort by age in ascending order
    slices.SortFunc(people, func(a, b Person) int {
        return cmp.Compare(a.Age, b.Age)
    })
    
    fmt.Println("Sorted by age:", people)
    // Output: [{Bob 25} {Alice 30} {Charlie 35}]
}












