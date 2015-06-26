package golib

// #cgo CFLAGS: -I./../../cppwrap
// #cgo LDFLAGS: -Wl,-rpath=./../../cppwrap -L./../../cppwrap -lcppwrap
// #include "cppwrap.h"
import "C"

import "fmt"
// exported interface starts with Capital Letter hence Golib_main instead of
// golib_main

func Golib_main() {
	C.CCppwrap()
	fmt.Println("golib_main called")
}
