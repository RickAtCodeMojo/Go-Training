package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var vuint8 uint8     //the set of all unsigned  8-bit integers (0 to 255)
	var vuint16 uint16   //the set of all unsigned 16-bit integers (0 to 65535)
	var vuint32 uint32   //the set of all unsigned 32-bit integers (0 to 4294967295)
	var vuint64 uint64   //the set of all unsigned 64-bit integers (0 to 18446744073709551615)
	var vint8 int8       //the set of all signed  8-bit integers (-128 to 127)
	var vint16 int16     //the set of all signed 16-bit integers (-32768 to 32767)
	var vint32 int32     //the set of all signed 32-bit integers (-2147483648 to 2147483647)
	var vint64 int64     //the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	var vfloat32 float32 //the set of all IEEE-754 32-bit floating-point numbers
	var vfloat64 float64 //the set of all IEEE-754 64-bit floating-point numbers

	var vcomplex64 complex64   //the set of all complex numbers with float32 real and imaginary parts
	var vcomplex128 complex128 //the set of all complex numbers with float64 real and imaginary parts

	var vbyte byte //alias for uint8
	var vrune rune //alias for int32
	//The value of an n-bit integer is n bits wide and represented using two's complement arithmetic.

	//There is also a set of predeclared numeric types with implementation-specific sizes:

	var vuint uint       //either 32 or 64 bits
	var vint int         //same size as uint
	var vuintptr uintptr //an unsigned integer large enough to store the uninterpreted bits of a pointer value

	vuint8 = math.MaxUint8
	fmt.Println("MaxUint8 ", vuint8)
	vuint16 = math.MaxUint16
	fmt.Println("MaxUint16 ", vuint16)
	vuint32 = math.MaxUint32
	fmt.Println("MaxUint32 ", vuint32)
	vuint64 = math.MaxUint64
	fmt.Println("MaxUint64 ", vuint64)
	vint8 = math.MaxInt8
	fmt.Println("MaxInt8 ", vint8)
	vint16 = math.MaxInt16
	fmt.Println("MaxInt16 ", vint16)
	vint32 = math.MaxInt32
	fmt.Println("MaxInt32 ", vint32)
	vint64 = math.MaxInt64
	fmt.Println("MaxInt64 ", vint64)
	vfloat32 = math.MaxFloat32
	fmt.Println("MaxFloat32 ", vfloat32)
	vfloat64 = math.MaxFloat64
	fmt.Println("MaxFloat64 ", vfloat64)
	vbyte = vuint8
	fmt.Println("byte", vbyte)
	vrune = vint32
	fmt.Println("rune", vrune)

	vcomplex64 = complex(vfloat32, vfloat32)
	fmt.Println("Complex 64:", vcomplex64)
	vcomplex128 = complex(vfloat64, vfloat64)
	fmt.Println("Complex 128:", vcomplex128)
	vuint = uint(vuint64)
	fmt.Println("uint", vuint)
	vint = int(vint64)
	fmt.Println("int", vint)
	vuintptr = (uintptr)(unsafe.Pointer(&vint))
	fmt.Println(vuintptr)
}
