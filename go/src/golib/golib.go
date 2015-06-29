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
	mystruct.kv_put("key1", 40000)
	fmt.Println("PUT ", "key1", " Value ",  40000)

	mystruct1 := kv_open("/tmp/test1", 1024, 1024)
	mystruct1.kv_put("key2", 4)
	fmt.Println("PUT ", "key2", " Value ",  4)

	key := "KM"
	size := 0

	mystruct.kv_get(&key, &size)
	fmt.Println("\nGet ", key, size)

	mystruct1.kv_get(&key, &size)
	fmt.Println("Get ", key, size)
}
