package main

import (
	"cmp"
	"fmt"
	"slices"
	"unsafe"
)

func Main_slice() {
	check_slice_from_array()
}

// method "append" affect-ит underlying-array-of-current-slice. 

func check_slice_from_array() {
	a := [5]int{ 0, 1, 2, 3, 4}
	fmt.Printf("a  | %p | %v | len=%d | cap=%d\n", &a, a, len(a), cap(a))

	s1 := a[:2]
	fmt.Printf("s1 | %p | %p | %v | len=%d | cap=%d\n", &s1, &s1[0], s1, len(s1), cap(s1))

	s2 := a[1:3]
	fmt.Printf("s2 | %p | %p | %v | len=%d | cap=%d\n", &s2, &s2[0], s2, len(s2), cap(s2))

	fmt.Println("append nothing")
	s1 = append(s1)
	fmt.Printf("a  | %p | %v | len=%d | cap=%d\n", &a, a, len(a), cap(a))
	fmt.Printf("s1 | %p | %p | %v | len=%d | cap=%d\n", &s1, &s1[0], s1, len(s1), cap(s1))
	fmt.Printf("s2 | %p | %p | %v | len=%d | cap=%d\n", &s2, &s2[0], s2, len(s2), cap(s2))

	fmt.Println("append 7, not assign to slice")
	_ = append(s1, 7)
	fmt.Printf("a  | %p | %v | len=%d | cap=%d\n", &a, a, len(a), cap(a))
	fmt.Printf("s1 | %p | %p | %v | len=%d | cap=%d\n", &s1, &s1[0], s1, len(s1), cap(s1))
	fmt.Printf("s2 | %p | %p | %v | len=%d | cap=%d\n", &s2, &s2[0], s2, len(s2), cap(s2))

	fmt.Println("append 8, not assign to slice")
	_ = append(s1, 8)
	fmt.Printf("a  | %p | %v | len=%d | cap=%d\n", &a, a, len(a), cap(a))
	fmt.Printf("s1 | %p | %p | %v | len=%d | cap=%d\n", &s1, &s1[0], s1, len(s1), cap(s1))
	fmt.Printf("s2 | %p | %p | %v | len=%d | cap=%d\n", &s2, &s2[0], s2, len(s2), cap(s2))

	fmt.Println("append 9, assign to slice")
	s1 = append(s1, 9)
	fmt.Printf("a  | %p | %v | len=%d | cap=%d\n", &a, a, len(a), cap(a))
	fmt.Printf("s1 | %p | %p | %v | len=%d | cap=%d\n", &s1, &s1[0], s1, len(s1), cap(s1))
	fmt.Printf("s2 | %p | %p | %v | len=%d | cap=%d\n", &s2, &s2[0], s2, len(s2), cap(s2))

	fmt.Println("append 10, assign to slice")
	s1 = append(s1, 10)
	fmt.Printf("a  | %p | %v | len=%d | cap=%d\n", &a, a, len(a), cap(a))
	fmt.Printf("s1 | %p | %p | %v | len=%d | cap=%d\n", &s1, &s1[0], s1, len(s1), cap(s1))
	fmt.Printf("s2 | %p | %p | %v | len=%d | cap=%d\n", &s2, &s2[0], s2, len(s2), cap(s2))

	fmt.Println("append 11, assign to slice")
	s1 = append(s1, 11)
	fmt.Printf("a  | %p | %v | len=%d | cap=%d\n", &a, a, len(a), cap(a))
	fmt.Printf("s1 | %p | %p | %v | len=%d | cap=%d\n", &s1, &s1[0], s1, len(s1), cap(s1))
	fmt.Printf("s2 | %p | %p | %v | len=%d | cap=%d\n", &s2, &s2[0], s2, len(s2), cap(s2))

	fmt.Println("append 12, assign to slice")
	s1 = append(s1, 12)
	fmt.Printf("a  | %p | %v | len=%d | cap=%d\n", &a, a, len(a), cap(a))
	fmt.Printf("s1 | %p | %p | %v | len=%d | cap=%d\n", &s1, &s1[0], s1, len(s1), cap(s1))
	fmt.Printf("s2 | %p | %p | %v | len=%d | cap=%d\n", &s2, &s2[0], s2, len(s2), cap(s2))
}
/*

output:

a  | 0xc00000c300 | [0 1 2 3 4] | len=5 | cap=5
s1 | 0xc0000080a8 | 0xc00000c300 | [0 1] | len=2 | cap=5
s2 | 0xc0000080d8 | 0xc00000c308 | [1 2] | len=2 | cap=4
append nothing
a  | 0xc00000c300 | [0 1 2 3 4] | len=5 | cap=5
s1 | 0xc0000080a8 | 0xc00000c300 | [0 1] | len=2 | cap=5
s2 | 0xc0000080d8 | 0xc00000c308 | [1 2] | len=2 | cap=4
append 7, not assign to slice
a  | 0xc00000c300 | [0 1 7 3 4] | len=5 | cap=5
s1 | 0xc0000080a8 | 0xc00000c300 | [0 1] | len=2 | cap=5
s2 | 0xc0000080d8 | 0xc00000c308 | [1 7] | len=2 | cap=4
append 8, not assign to slice
a  | 0xc00000c300 | [0 1 8 3 4] | len=5 | cap=5
s1 | 0xc0000080a8 | 0xc00000c300 | [0 1] | len=2 | cap=5
s2 | 0xc0000080d8 | 0xc00000c308 | [1 8] | len=2 | cap=4
append 9, assign to slice
a  | 0xc00000c300 | [0 1 9 3 4] | len=5 | cap=5
s1 | 0xc0000080a8 | 0xc00000c300 | [0 1 9] | len=3 | cap=5
s2 | 0xc0000080d8 | 0xc00000c308 | [1 9] | len=2 | cap=4
append 10, assign to slice
a  | 0xc00000c300 | [0 1 9 10 4] | len=5 | cap=5
s1 | 0xc0000080a8 | 0xc00000c300 | [0 1 9 10] | len=4 | cap=5
s2 | 0xc0000080d8 | 0xc00000c308 | [1 9] | len=2 | cap=4
append 11, assign to slice
a  | 0xc00000c300 | [0 1 9 10 11] | len=5 | cap=5
s1 | 0xc0000080a8 | 0xc00000c300 | [0 1 9 10 11] | len=5 | cap=5
s2 | 0xc0000080d8 | 0xc00000c308 | [1 9] | len=2 | cap=4
append 12, assign to slice
a  | 0xc00000c300 | [0 1 9 10 11] | len=5 | cap=5
s1 | 0xc0000080a8 | 0xc00000e1e0 | [0 1 9 10 11 12] | len=6 | cap=10
s2 | 0xc0000080d8 | 0xc00000c308 | [1 9] | len=2 | cap=4

*/

// internal representation for type "slice" (p.s. golang source code: unsafeheader.go)
type SliceInternalStruct struct {
	// pointer to the starting element of underlying array.
	Data unsafe.Pointer

	// number of elements in this slice (slice = segment of underlying array)
	Len  int

	// the remaining elements from starting element of slice to the end of the underlying array.
	Cap  int 
}

// on append(to:curSlice, val:<valueToAppend>): 
// if curSlice.cap недостаточно, то создается новый underlyingarray (в новой области памяти) с cap=oldCap*2, в него копируются значения из старого слайса, затем в него добавляется (в конец) <valueToAppend>.
// p.s. т.е. чтобы не создавать(не аллоцировать) каждый раз новый underlyingarray => нужно сразу указывать требуемую cap при инициализации слайса.  
func check_append_to_nil_slice() {
	// only-declare (not initialize) slice
	var s []int // to "fix" alloc leaks: s := make([]int, 0, 9)
	fmt.Printf("s | %p | isNil: %t | cap=%d | len=%d | %v\n", &s, s == nil, cap(s), len(s), s)
	
	for i := 1; i < 10; i++ {
		s = append(s, i) // append value "i"
		fmt.Printf("s | %p | %p | cap=%d | len=%d | %v\n", &s, &s[0], cap(s), len(s), s)
	}

/*

// output:

s | 0xc0000080a8 | isNil: true | cap=0 | len=0 | []
s | 0xc0000080a8 | 0xc00000a0d0 | cap=1 | len=1 | [1]
s | 0xc0000080a8 | 0xc00000a0e0 | cap=2 | len=2 | [1 2]
s | 0xc0000080a8 | 0xc000012260 | cap=4 | len=3 | [1 2 3]
s | 0xc0000080a8 | 0xc000012260 | cap=4 | len=4 | [1 2 3 4]
s | 0xc0000080a8 | 0xc000010240 | cap=8 | len=5 | [1 2 3 4 5]
s | 0xc0000080a8 | 0xc000010240 | cap=8 | len=6 | [1 2 3 4 5 6]
s | 0xc0000080a8 | 0xc000010240 | cap=8 | len=7 | [1 2 3 4 5 6 7]
s | 0xc0000080a8 | 0xc000010240 | cap=8 | len=8 | [1 2 3 4 5 6 7 8]
s | 0xc0000080a8 | 0xc00001a400 | cap=16 | len=9 | [1 2 3 4 5 6 7 8 9]

*/

}

func check_append_to_slice() {
	s1 := make([]int, 2, 4)
  	fmt.Printf("s1 | %p | %p | cap=%d | len=%d | %v\n", &s1, &s1[0], cap(s1), len(s1), s1)

	s2 := s1[0:4]
	fmt.Printf("s2 | %p | %p | cap=%d | len=%d | %v\n", &s2, &s2[0], cap(s2), len(s2), s2)

	fmt.Println("modify index 1")
	s2[1] = 4
	fmt.Printf("s1 | %p | %p | cap=%d | len=%d | %v\n", &s1, &s1[0], cap(s1), len(s1), s1)
	fmt.Printf("s2 | %p | %p | cap=%d | len=%d | %v\n", &s2, &s2[0], cap(s2), len(s2), s2)

	fmt.Println("append to s1, not assign")
	_ = append(s1, 7)
	fmt.Printf("s1 | %p | %p | cap=%d | len=%d | %v\n", &s1, &s1[0], cap(s1), len(s1), s1)
	fmt.Printf("s2 | %p | %p | cap=%d | len=%d | %v\n", &s2, &s2[0], cap(s2), len(s2), s2)

	fmt.Println("append to s2, not assign")
	_ = append(s2, 8)
	fmt.Printf("s1 | %p | %p | cap=%d | len=%d | %v\n", &s1, &s1[0], cap(s1), len(s1), s1)
	fmt.Printf("s2 | %p | %p | cap=%d | len=%d | %v\n", &s2, &s2[0], cap(s2), len(s2), s2)

	fmt.Println("append to s2, assign to s1")
	s1 = append(s2, 9)
	fmt.Printf("s1 | %p | %p | cap=%d | len=%d | %v\n", &s1, &s1[0], cap(s1), len(s1), s1)
	fmt.Printf("s2 | %p | %p | cap=%d | len=%d | %v\n", &s2, &s2[0], cap(s2), len(s2), s2)

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

func check_syntax_slice() {
	fmt.Println("set len=3")
	s := make([]int, 3)
	fmt.Printf("s | %p | %p | cap=%d | len=%d | %v\n", &s, &s[0], cap(s), len(s), s)

	fmt.Println("set 3 values")
	s = []int{4, 5, 6}
	fmt.Printf("s | %p | %p | cap=%d | len=%d | %v\n", &s, &s[0], cap(s), len(s), s)

	fmt.Println("set len=3,cap=5")
	s = make([]int, 3, 5)
	fmt.Printf("s | %p | %p | cap=%d | len=%d | %v\n", &s, &s[0], cap(s), len(s), s)

	fmt.Println("append multiple values")
	s = append(s, 12, 13, 14)
	fmt.Printf("s | %p | %p | cap=%d | len=%d | %v\n", &s, &s[0], cap(s), len(s), s)

	fmt.Println("append all values from other slice")
	s2 := []int{15, 16}
	s = append(s, s2...)
	fmt.Printf("s | %p | %p | cap=%d | len=%d | %v\n", &s, &s[0], cap(s), len(s), s)
}

/*

// output:

set len=3
s | 0xc00009a090 | 0xc00009e048 | cap=3 | len=3 | [0 0 0]
set 3 values
s | 0xc00009a090 | 0xc00009e060 | cap=3 | len=3 | [4 5 6]
set len=3,cap=5
s | 0xc00009a090 | 0xc0000ac000 | cap=5 | len=3 | [0 0 0]
append multiple values
s | 0xc00009a090 | 0xc0000a20a0 | cap=10 | len=6 | [0 0 0 12 13 14]
append all values from other slice
s | 0xc00009a090 | 0xc0000a20a0 | cap=10 | len=8 | [0 0 0 12 13 14 15 16]

*/

func compare_alloc_array_slice() {
    arr := [5]int{1, 2, 3, 4, 5}
	// Array = 
	// {
	// 	 elem of Type,
	// 	 len
	// }
	// => alloc size = <elem alloc size> * len = <Type alloc size> * len
	// example: [5]int32 = 32 bits * 5 = 4 bytes * 5 = 20 bytes
	// p.s. example ram-memory:
	// address: | key: | value:       | alloc_size:
	// 0xA      | arr  | [1,2,3,4,5]  | <type of elem> * 5

	slice := []int{1, 2, 3, 4, 5, 6}
    // Slice = 
	// { 
	// 	 <Ptr to underlArray> int, 
	//   len int, 
	//   cap int 
	// }
	// => alloc size = 3 int fields = 3 * <4/8 bytes> = 12/24 bytes (on 32/64 machine)
	// p.s. example ram-memory:
	// address: | key:                 | value:                    | alloc_size:
	// 0xB      | internal_underlArray | [1,2,3,4,5,6]             | <type of elem> * 6
	// 0xC      | slice                | { ptr:0xB, len:6, cap:6 } | 3 int fields

	slice_from_arr := arr[:]
	// Slice - аналогично: 3 int fields.
	// p.s. example ram-memory:
	// address:  | key:            | value:                     | alloc_size:
	// 0xD       | slice_from_arr  | { ptr:0xA, len:5, cap:5 }  | 3 int fields

    fmt.Printf("Array: %v, size: %d\n", arr, unsafe.Sizeof(arr))
    fmt.Printf("Slice: %v, size: %d\n", slice, unsafe.Sizeof(slice))
    fmt.Printf("SliceFromArr: %v, size: %d\n", slice_from_arr, unsafe.Sizeof(slice_from_arr))
}

/*

// output: (64-bit laptop)

Array: [1 2 3 4 5], size: 40
Slice: [1 2 3 4 5 6], size: 24
SliceFromArr: [1 2 3 4 5], size: 24

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

    // Sort by age
    slices.SortFunc(people, func(a, b Person) int {
        return cmp.Compare(a.Age, b.Age)
    })
    
    fmt.Println("Sorted by age:", people)
    // Output: [{Bob 25} {Alice 30} {Charlie 35}]
}












