package utils

import "unsafe"

// NocopyStr2ByteSlice use nocopy to convert the string to byte slice
func NocopyStr2ByteSlice(str string) []byte {
	ptr := (*[2]uintptr)(unsafe.Pointer(&str))
	return *(*[]byte)(unsafe.Pointer(&[3]uintptr{ptr[0], ptr[1], ptr[1]}))
}

// NocopyByteSlice2Str use nocopy to convert the byte slice to string
func NocopyByteSlice2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
