package util

import (
	"testing"
)

var (
	str = "hello"
	bs = []byte(str)
)

func TestBytes2Str(t *testing.T) {
	t.Log(str==Bytes2Str(bs))
}

func TestStr2Bytes(t *testing.T) {
	t.Log(str==string(Str2Bytes(str)))
}

func BenchmarkBytes2Str(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes2Str(bs)
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(bs)
	}
}

func BenchmarkStr2Bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Str2Bytes(str)
	}
}

func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(str)
	}
}