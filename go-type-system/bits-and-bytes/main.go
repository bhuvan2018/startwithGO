package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Different integer sizes (in bits)
	var tiny int8 = 100        // 8 bits  = 1 byte
	var medium int32 = 100000  // 32 bits = 4 bytes
	var huge int64 = 100000000 // 64 bits = 8 bytes

	// Default 'int' matches your CPU architecture (64-bit on modern PCs = 8 bytes)
	var standard int = 42

	fmt.Println("--- Memory Size in Bytes ---")
	fmt.Printf("int8   takes %d byte  (8 bits)\n", unsafe.Sizeof(tiny))
	fmt.Printf("int32  takes %d bytes (32 bits)\n", unsafe.Sizeof(medium))
	fmt.Printf("int64  takes %d bytes (64 bits)\n", unsafe.Sizeof(huge))
	fmt.Printf("int    takes %d bytes (64 bits on 64-bit OS)\n", unsafe.Sizeof(standard))

	// Bonus: See how a number looks in binary (bits) using %b!
	number := 13
	fmt.Printf("\nNumber %d in binary (bits) is: %b\n", number, number)
}
