package gomathbits

import "errors"
import "encoding/hex"

// BytesToUint8 return uint8 from bits
func BytesToUint8(bits []byte) (u uint8, err error) {
	if len(bits) != 1 {
		return 0, errors.New("invalid syntax")
	}
	return bits[0], nil
}

// Uint8ToBytes returns bytes
func Uint8ToBytes(u uint8) []byte {
	b := make([]byte, 1)
	b[0] = byte(u)
	return b
}

// BytesToUint16 return uint16 from bits
func BytesToUint16(bits []byte) (u uint16, err error) {
	if len(bits) != 2 {
		return 0, errors.New("invalid syntax")
	}
	u = (uint16)(bits[1])<<8 | uint16(bits[0])
	return u, nil
}

// Uint16ToBytes returns bytes
func Uint16ToBytes(u uint16) []byte {
	b := make([]byte, 2)
	b[0] = byte(u & 0xFF)
	b[1] = byte((u >> 8) & 0xFF)
	return b
}

// BytesToUint32 return uint32 from bits
func BytesToUint32(bits []byte) (u uint32, err error) {
	if len(bits) != 4 {
		return 0, errors.New("invalid syntax")
	}
	u = (uint32)(bits[3])<<24 | (uint32)(bits[2])<<16 |
		(uint32)(bits[1])<<8 | uint32(bits[0])
	return u, nil
}

// Uint32ToBytes returns bytes
func Uint32ToBytes(u uint32) []byte {
	b := make([]byte, 4)
	b[0] = byte(u & 0xFF)
	b[1] = byte((u >> 8) & 0xFF)
	b[2] = byte((u >> 16) & 0xFF)
	b[3] = byte((u >> 24) & 0xFF)
	return b
}

// BytesToUint64 return uint64 from bits
func BytesToUint64(bits []byte) (u uint64, err error) {
	if len(bits) != 8 {
		return 0, errors.New("invalid syntax")
	}
	u = (uint64)(bits[7])<<56 | (uint64)(bits[6])<<48 |
		(uint64)(bits[5])<<40 | (uint64)(bits[4])<<32 |
		(uint64)(bits[3])<<24 | (uint64)(bits[2])<<16 |
		(uint64)(bits[1])<<8 | uint64(bits[0])
	return u, nil
}

// Uint64ToBytes returns bytes
func Uint64ToBytes(u uint64) []byte {
	b := make([]byte, 8)
	b[0] = byte(u & 0xFF)
	b[1] = byte((u >> 8) & 0xFF)
	b[2] = byte((u >> 16) & 0xFF)
	b[3] = byte((u >> 24) & 0xFF)
	b[4] = byte((u >> 32) & 0xFF)
	b[5] = byte((u >> 40) & 0xFF)
	b[6] = byte((u >> 48) & 0xFF)
	b[7] = byte((u >> 56) & 0xFF)
	return b
}

// EncodeToString returns the hexadecimal encoding of src.
func EncodeToString(src []byte) string {
	return hex.EncodeToString(src)
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

	switch bitSize {
	case 8:
		uu, err := BytesToUint8(bits)
		if err != nil {
			return 0, err
		}
		return uint64(uu), nil
	case 16:
		uu, err := BytesToUint16(bits)
		if err != nil {
			return 0, err
		}
		return uint64(uu), nil
	case 32:
		uu, err := BytesToUint32(bits)
		if err != nil {
			return 0, err
		}
		return uint64(uu), nil
	case 64:
		uu, err := BytesToUint64(bits)
		if err != nil {
			return 0, err
		}
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

	switch bitSize {
	case 8:
		uu, err := BytesToUint8(bits)
		if err != nil {
			return 0, err
		}
		return int64(Int8frombits(uu)), nil
	case 16:
		uu, err := BytesToUint16(bits)
		if err != nil {
			return 0, err
		}
		return int64(Int16frombits(uu)), nil
	case 32:
		uu, err := BytesToUint32(bits)
		if err != nil {
			return 0, err
		}
		return int64(Int32frombits(uu)), nil
	case 64:
		uu, err := BytesToUint64(bits)
		if err != nil {
			return 0, err
		}
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

	switch bitSize {
	case 32:
		uu, err := BytesToUint32(bits)
		if err != nil {
			return 0, err
		}
		return float64(Float32frombits(uu)), nil
	case 64:
		uu, err := BytesToUint64(bits)
		if err != nil {
			return 0, err
		}
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

	uu, err := BytesToUint32(bits)
	if err != nil {
		return 0, err
	}
	return Float32frombits(uu), nil
}

// ParseFloat64 converts the string s to a floating-point number
// with the precision specified by 64 for float64
func ParseFloat64(s string) (f float64, err error) {
	bits, err := hex.DecodeString(s)
	if err != nil {
		return 0, err
	}

	uu, err := BytesToUint64(bits)
	if err != nil {
		return 0, err
	}
	return Float64frombits(uu), nil
}
