package hello

/*
#include "cgoexample/hello/hello.h"
*/
import "C"

func Half(x int) int {
	return int(C.half(C.int(x)))
}
