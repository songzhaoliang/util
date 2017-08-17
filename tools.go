package util

import (
	"unsafe"
)

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Str2Bytes(str string) []byte {
	temp := (*[2]uintptr)(unsafe.Pointer(&str))
	res := [3]uintptr{temp[0],temp[1],temp[1]}
	return *(*[]byte)(unsafe.Pointer(&res))
}
