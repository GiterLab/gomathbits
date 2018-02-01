package gomathbits

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"bytes"
)

// BytesToUint8 return uint8 from bits
func BytesToUint8(bits []byte) (u uint8) {
	if len(bits) < 1 {
		return 0
	}
	return bits[0]
}

// Uint8ToBytes returns bytes
func Uint8ToBytes(u uint8) []byte {
	b := make([]byte, 1)
	b[0] = byte(u)
	return b
}

// BytesToUint16 return uint16 from bits
func BytesToUint16(bits []byte) (u uint16) {
	if len(bits) < 2 {
		return binary.LittleEndian.Uint16(BytesCombine(bits,[]byte{0,0}))
	}
	return binary.LittleEndian.Uint16(bits)
}

// Uint16ToBytes returns bytes
func Uint16ToBytes(u uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, u)
	return b
}

// BytesToUint32 return uint32 from bits
func BytesToUint32(bits []byte) (u uint32) {
	if len(bits) < 4 {
		return binary.LittleEndian.Uint32(BytesCombine(bits,[]byte{0,0,0,0}))
	}
	return binary.LittleEndian.Uint32(bits)
}

// Uint32ToBytes returns bytes
func Uint32ToBytes(u uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, u)
	return b
}

// BytesToUint64 return uint64 from bits
func BytesToUint64(bits []byte) (u uint64) {
	if len(bits) < 8 {
		return binary.LittleEndian.Uint64(BytesCombine(bits,[]byte{0,0,0,0,0,0,0,0}))
	}
	return binary.LittleEndian.Uint64(bits)
}

// Uint64ToBytes returns bytes
func Uint64ToBytes(u uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, u)
	return b
}

// EncodeToString returns the hexadecimal encoding of b.
func EncodeToString(b []byte) string {
	return hex.EncodeToString(bytesSwap(b))
}

// ParseUInt parse hexadecimal encoding of string to uint64
// The bitSize argument specifies the integer type
// that the result must fit into. Bit sizes 8, 16, 32, and 64
// correspond to uint8, uint16, uint32, and uint64.
func ParseUInt(s string, bitSize int) (u uint64, err error) {
	bits, err := hex.DecodeString(s)
	if err != nil {
		return 0, err
	}
	bytesSwap(bits)
	switch bitSize {
	case 8:
		uu := BytesToUint8(bits)
		return uint64(uu), nil
	case 16:
		uu := BytesToUint16(bits)
		return uint64(uu), nil
	case 32:
		uu := BytesToUint32(bits)
		return uint64(uu), nil
	case 64:
		uu := BytesToUint64(bits)
		return uint64(uu), nil
	}
	return 0, errors.New("invalid syntax")
}

// ParseInt parse hexadecimal encoding of string to int64
// The bitSize argument specifies the integer type
// that the result must fit into. Bit sizes 8, 16, 32, and 64
// correspond to int8, int16, int32, and int64.
func ParseInt(s string, bitSize int) (i int64, err error) {
	bits, err := hex.DecodeString(s)
	if err != nil {
		return 0, err
	}
	bytesSwap(bits)
	switch bitSize {
	case 8:
		uu := BytesToUint8(bits)
		return int64(Int8frombits(uu)), nil
	case 16:
		uu := BytesToUint16(bits)
		return int64(Int16frombits(uu)), nil
	case 32:
		uu := BytesToUint32(bits)
		return int64(Int32frombits(uu)), nil
	case 64:
		uu := BytesToUint64(bits)
		return Int64frombits(uu), nil
	}
	return 0, errors.New("invalid syntax")
}

// ParseFloat converts the string s to a floating-point number
// with the precision specified by bitSize: 32 for float32, or 64 for float64.
// When bitSize=32, the result still has type float64, but it will be
// convertible to float32 without changing its value.
//
// If s is well-formed and near a valid floating point number,
// ParseFloat returns the nearest floating point number rounded
// using IEEE754 unbiased rounding.
func ParseFloat(s string, bitSize int) (f float64, err error) {
	bits, err := hex.DecodeString(s)
	if err != nil {
		return 0, err
	}
	bytesSwap(bits)
	switch bitSize {
	case 32:
		uu := BytesToUint32(bits)
		return float64(Float32frombits(uu)), nil
	case 64:
		uu := BytesToUint64(bits)
		return Float64frombits(uu), nil
	}
	return 0, errors.New("invalid syntax")
}

// ParseFloat32 converts the string s to a floating-point number
// with the precision specified by 32 for float32
func ParseFloat32(s string) (f float32, err error) {
	bits, err := hex.DecodeString(s)
	if err != nil {
		return 0, err
	}
	bytesSwap(bits)
	uu := BytesToUint32(bits)
	return Float32frombits(uu), nil
}

// ParseFloat64 converts the string s to a floating-point number
// with the precision specified by 64 for float64
func ParseFloat64(s string) (f float64, err error) {
	bits, err := hex.DecodeString(s)
	if err != nil {
		return 0, err
	}
	bytesSwap(bits)
	uu := BytesToUint64(bits)
	return Float64frombits(uu), nil
}

func bytesSwap(b []byte) []byte {
	for from, to := 0, len(b)-1; from < to; from, to = from+1, to-1 {
		b[from], b[to] = b[to], b[from]
	}
	return b
}

// BytesCombine 组合数据包
func BytesCombine(pBytes ...[]byte) []byte {
	lenth := len(pBytes)
	s := make([][]byte, lenth)
	for i, v := range pBytes {
		s[i] = v
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}
