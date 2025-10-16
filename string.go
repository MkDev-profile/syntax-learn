package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func Main_string() {
	string_check_1()
}

// internal representation for type "string" (p.s. golang source code: unsafeheader.go)
type InternalStringStruct struct {
    Data uintptr
    Len  int
}

func printString(varName string, ptr *string) {
	val := *ptr
    internalStruct := (*InternalStringStruct)(unsafe.Pointer(ptr))

    // underlArrPtr := unsafe.StringData(val) // equals to: internalStruct.Data

    fmt.Printf("%s: Val = %s, Len = %d, UnderlArrayPtr = %#x\n", 
		varName, val, internalStruct.Len, internalStruct.Data)
}

func string_check_1() {
    s1 := "Hello"
    printString("s1", &s1)

    // this creates a new string, another underlying array
    s2 := s1 + " World"
	printString("s2", &s2)
    
    // slice from string, SAME underlying array
    s1_sl_1 := s1[:] 
    printString("s1_sl_1", &s1_sl_1)

    // slice from string, SAME underlying array
	s1_sl_2 := s1[1:] 
    printString("s1_sl_2", &s1_sl_2)

    // Same content, SAME underlying array
    s4 := "Hello" 
	printString("s4", &s4)

    fmt.Println("s4[0] =", s4[0])

    // assign/modify by index запрещено (Compilation error)!
    // s4[0] = 'h' 

    // get ptr by index запрещено (Compilation error)!
    //fmt.Println(&s4[0]) 

    // convert to byte type
    var b0 byte = byte(s4[0]) // (newly allocated) byte of string
    var b1 byte = byte(s4[1]) // (newly allocated) byte of string

    fmt.Printf("b0 = %p, b1 = %p\n", &b0, &b1) 

    // Convert to byte slice (allocates once)
    bytes := []byte(s4)
    fmt.Printf("bytes0 = %p\n", &bytes[0])

    // modify value of zero-byte
    bytes[0] = '*'

    s5 := string(bytes) // (newly allocated) string
    printString("s5", &s5) 

    s6 := string(b0) // (newly allocated) string
    printString("s6", &s6) 

    var sb strings.Builder
    sbStr := sb.String() // underlying array = nil
    printString("sb", &sbStr)

    sb.WriteByte(65)
    sbStr = sb.String() // underlying array created
    printString("sb", &sbStr)

    sb.WriteByte(70)
    sbStr = sb.String() // SAME underlying array
    printString("sb", &sbStr)

/*

output:

$ go run .
s1: Val = Hello, Len = 5, UnderlArrayPtr = 0x7ff64a856fcf
s2: Val = Hello World, Len = 11, UnderlArrayPtr = 0xc00000a0e0
s1_sl_1: Val = Hello, Len = 5, UnderlArrayPtr = 0x7ff64a856fcf
s1_sl_2: Val = ello, Len = 4, UnderlArrayPtr = 0x7ff64a856fd0
s4: Val = Hello, Len = 5, UnderlArrayPtr = 0x7ff64a856fcf
s4[0] = 72
b0 = 0xc00000a108, b1 = 0xc00000a109
bytes0 = 0xc00000a110
s5: Val = *ello, Len = 5, UnderlArrayPtr = 0xc00000a118
s6: Val = H, Len = 1, UnderlArrayPtr = 0xc00000a128
sb: Val = , Len = 0, UnderlArrayPtr = 0x0
sb: Val = A, Len = 1, UnderlArrayPtr = 0xc00000a138
sb: Val = AF, Len = 2, UnderlArrayPtr = 0xc00000a138

*/

}










