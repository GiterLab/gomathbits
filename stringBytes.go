package gomathbits

import "unsafe"

// StringToBytes 字符串转字节流
func StringToBytes(s string) []byte {
	// []byte(s)内部会进行复制过程，占用较多的内存
	return []byte(s)
}

// BytesToString 字节流转字符串
func BytesToString(b []byte) string {
	return string(b)
}

// StringToBytesPerf 字符串转字节流, 性能优化版
func StringToBytesPerf(s string) []byte {
	// 在内存中直接转换
	x := (*[2]uintptr)(unsafe.Pointer(&s)) // string x[0]: 内存起始位置x x[1]: 长度
	h := [3]uintptr{x[0], x[1], x[1]}      // []byte x[0]：内存起始位置x x[1]：实际长度 len(x) x[2]：容量长度 cap(x)
	return *(*[]byte)(unsafe.Pointer(&h))
}

// BytesToStringPerf 字节流转字符串，性能优化版
func BytesToStringPerf(b []byte) string {
	// byte 直接转换为指向字符串指针的指针
	return *(*string)(unsafe.Pointer(&b))
}
