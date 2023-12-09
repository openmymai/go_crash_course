package main

import "fmt"

func main() {
    var myAge int32 = 21
    const isCool = true
    name, email := "Gabriel", "test@test.com"
    size := 1.3

    fmt.Println(name, myAge, isCool, size, email)
    fmt.Printf("%T\n", name)

}

// All types:
// string
// bool
// int
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte - alias for uint8
// rune - alias for int32
// float32 float64
// complex64 complex128
