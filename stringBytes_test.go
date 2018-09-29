package gomathbits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const StringTest = "hello, world!"

func TestStringToBytes(t *testing.T) {
	assert.Equal(t, StringToBytes(StringTest), []byte(StringTest))
}

func TestBytesToString(t *testing.T) {
	assert.Equal(t, BytesToString([]byte(StringTest)), StringTest)
}

func BenchmarkStringToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToBytes(StringTest)
	}
}

func BenchmarkBytesToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BytesToString([]byte(StringTest))
	}
}

func TestStringToBytesPerf(t *testing.T) {
	assert.Equal(t, StringToBytesPerf(StringTest), []byte(StringTest))
}

func TestBytesToStringPerf(t *testing.T) {
	assert.Equal(t, BytesToStringPerf([]byte(StringTest)), StringTest)
}

func BenchmarkStringToBytePerf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToBytesPerf(StringTest)
	}
}

func BenchmarkByteToStringPerf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BytesToStringPerf([]byte(StringTest))
	}
}
