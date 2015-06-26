package golib 

// #include "cppwrap.h"
import "C"
import "fmt"

type kv_interface interface {
	kv_get(key string, key_size int) int;
	kv_put(key string, key_size int) int;
}

type kv_struct struct {
	path string
	ssize, msize int
}

func (kv kv_struct) kv_get(key string, key_size int) int {
	fmt.Println("Key ", key, " Size ", key_size);
	return 0
}

func (kv kv_struct) kv_put(key string, key_size int) int {
	fmt.Println("Key ", key, " Size ", key_size);
	return 0
}

func kv_open(path string, ssize int, msize int) kv_struct {
	kv := kv_struct{path:path, ssize:ssize, msize:msize}	
	fmt.Println("kv open \n", kv)
	//C.CCppwrap()
	return kv
}
