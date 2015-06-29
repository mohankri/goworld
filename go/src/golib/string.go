package golib
//import "fmt"

func Reverse_string(reverse string) (new string) {
	runes := []rune(reverse)
	j := len(runes)
	for i:= 0; i < len(runes)/2; i++ {
		j = j-1
		ch := runes[i]
		runes[i] = runes[j]
		runes[j] = ch
	}
	return string(runes)
}
