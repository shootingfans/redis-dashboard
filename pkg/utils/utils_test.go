package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NocopyStr2ByteSlice(t *testing.T) {
	str := "bcd"
	b := NocopyStr2ByteSlice(str)
	assert.Equal(t, []byte{0x62, 0x63, 0x64}, b)
	assert.Equal(t, len(b), 3)
	assert.Equal(t, cap(b), 3)
}

func Test_NocopyByteSlice2Str(t *testing.T) {
	b := []byte{0x61, 0x64, 0x65}
	str := NocopyByteSlice2Str(b)
	assert.Equal(t, "ade", str)
	assert.Equal(t, len(str), 3)
}

func Benchmark_NocopyStr2ByteSlice(b *testing.B) {
	str := "abcde"
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		NocopyStr2ByteSlice(str)
	}
}

func Benchmark_NocopyByteSlice2Str(b *testing.B) {
	by := []byte{0x63, 0x64, 0x65, 0x66, 0x67}
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		NocopyByteSlice2Str(by)
	}
}
