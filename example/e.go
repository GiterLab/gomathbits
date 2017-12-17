package main

import (
	"fmt"
	"math"
	"os"

	"github.com/GiterLab/gomathbits"
)

func main() {
	var b []byte
	var i int64
	var err error
	var f32 float32
	var f64 float64

	// 80
	b = gomathbits.Uint8ToBytes(0x80)
	fmt.Println(gomathbits.EncodeToString(b)) // 80

	// ff00
	b = gomathbits.Uint16ToBytes(0xFF00)
	fmt.Println(gomathbits.EncodeToString(b)) // ff00

	// 12345678
	b = gomathbits.Uint32ToBytes(0x12345678)
	fmt.Println(gomathbits.EncodeToString(b)) // 12345678

	// 123456789abcdef0
	b = gomathbits.Uint64ToBytes(0x123456789abcdef0)
	fmt.Println(gomathbits.EncodeToString(b)) // 123456789abcdef0

	// 8bits -128
	i, err = gomathbits.ParseInt("80", 8)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(i) // -128

	// 16bits -256
	i, err = gomathbits.ParseInt("ff00", 16)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(i) // -256

	// 32bits -16777216
	i, err = gomathbits.ParseInt("ff000000", 32)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(i) // -16777216

	// 64bits -72057589759737856
	i, err = gomathbits.ParseInt("ff000000ff000000", 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(i) // -72057589759737856

	// float32 C3B23A5E
	f32, err = gomathbits.ParseFloat32(fmt.Sprintf("%02X", math.Float32bits(-356.456))) // C3B23A5E
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(f32, fmt.Sprintf("%02X", math.Float32bits(-356.456))) // -356.456

	// float64 C076474BC6A7EF9E
	f64, err = gomathbits.ParseFloat64(fmt.Sprintf("%02X", math.Float64bits(356.456))) // C076474BC6A7EF9E
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(f64, fmt.Sprintf("%02X", math.Float64bits(356.456))) // -356.456
}
