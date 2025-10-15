package main

import (
	"unsafe"
)

func Main_conversion() {

}

func convertIntSliceToByteSlice(ints []int) []byte {
	if len(ints) == 0 {
		return nil
	}

	bytesCount := len(ints) * int(unsafe.Sizeof(ints[0]))

	internalSlice := &SliceInternalStruct{
		Data: unsafe.Pointer(&ints[0]),
		Len:  bytesCount,
		Cap:  bytesCount,
	}

	return *(*[]byte)(unsafe.Pointer(internalSlice))
}

func convertByteSliceToIntSlice(bytes []byte) []int {
	if len(bytes) == 0 {
		return nil
	}

	intSize := int(unsafe.Sizeof(int(0)))
	intLen := len(bytes) / intSize

	internalSlice := &SliceInternalStruct{
		Data: unsafe.Pointer(&bytes[0]),
		Len:  intLen,
		Cap:  intLen,
	}

	return *(*[]int)(unsafe.Pointer(&internalSlice))
}


















