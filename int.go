package gomathbits

import "unsafe"

// Int8bits returns uint8
func Int8bits(i int8) uint8 {
	return *(*uint8)(unsafe.Pointer(&i))
}

// Int8frombits return int8
func Int8frombits(b uint8) int8 {
	return *(*int8)(unsafe.Pointer(&b))
}

// Int16bits returns uint16
func Int16bits(i int16) uint16 {
	return *(*uint16)(unsafe.Pointer(&i))
}

// Int16frombits return int16
func Int16frombits(b uint16) int16 {
	return *(*int16)(unsafe.Pointer(&b))
}

// Int32bits returns uint32
func Int32bits(i int32) uint32 {
	return *(*uint32)(unsafe.Pointer(&i))
}

// Int32frombits return int32
func Int32frombits(b uint32) int32 {
	return *(*int32)(unsafe.Pointer(&b))
}

// Int64bits returns uint64
func Int64bits(i int64) uint64 {
	return *(*uint64)(unsafe.Pointer(&i))
}

// Int64frombits return int64
func Int64frombits(b uint64) int64 {
	return *(*int64)(unsafe.Pointer(&b))
}
