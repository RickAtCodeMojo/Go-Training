package main

import (
	"fmt"
	"unsafe"
)

type SpaceMonster struct {
	Speed  int32 // 4 bytes
	Health int64 // 8 bytes
	Class  byte  //1 bytes

}

type SpaceMonster2 struct {
	Health int64 // 8 bytes
	Speed  int32 // 4 bytes
	Class  byte  //1 bytes
}

type TestStruct struct {
	b1 byte
	b2 byte
}

func main() {
	var tim SpaceMonster
	var b TestStruct
	fmt.Println(unsafe.Sizeof(b))
	//var tim2 SpaceMonster
	fmt.Println(unsafe.Sizeof(tim)) // should be 8+4+2 = 10 bytes in theory

	fmt.Println(unsafe.Pointer(&tim.Speed))
	fmt.Println(unsafe.Pointer(&tim.Health))
	fmt.Println(unsafe.Pointer(&tim.Class))

	var bob SpaceMonster2
	fmt.Println(unsafe.Sizeof(bob)) // should be the same, but it isn't, why?
	fmt.Println(unsafe.Pointer(&bob.Health))
	fmt.Println(unsafe.Pointer(&bob.Speed))
	fmt.Println(unsafe.Pointer(&bob.Class))

}
