package golib

// #cgo CFLAGS: -I./../../cppwrap
// #cgo LDFLAGS: -Wl,-rpath=./../../cppwrap -L./../../cppwrap -lcppwrap
// #include "cppwrap.h"
import "C"

import "fmt"
// exported interface starts with Capital Letter hence Golib_main instead of
// golib_main

func Golib_main() {
	mystruct := kv_open("/tmp/test", 1024, 1024)
	mystruct.kv_put("key1", 4)
	mystruct1 := kv_open("/tmp/test1", 1024, 1024)
	mystruct1.kv_put("key2", 4)
	fmt.Println("golib_main called ", mystruct.path, mystruct1.path)
	key := make([]string, 10)	
	size := 0
	mystruct1.kv_get(key, size)
	fmt.Println("key received 1 ", key, size)
	//mystruct.kv_get(key, size)
	//fmt.Println("key received 0 ", key, size)
}
