package main

/*compiles the C files with the gcc options
-g(enable debug symbols) and -Wall(enable all warnings) enabled*/

/*
#cgo CFLAGS: -g -Wall
#include "cgoexample/hello.h"
#include "cgoexample/sum.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := Hello()
	if err != nil {
		log.Fatal(err)
	}
	//Call to int function with two params
	res, err := makeSum(5, 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sum of 5 + 4 is %d\n", res)
}

// Hello is a C binding to the Hello World "C" programm
// As a Go user you could use now the Hello function transparently
// without knowing that it is calling a C function
func Hello() error {
	_, err := C.Hello() // We ignore first result as it is a void function
	if err != nil {
		return errors.New("error calling Hello function:" + err.Error())
	}
	return nil
}

func makeSum(a, b int) (int, error) {
	//Convert Go ints to C ints
	aC := C.int(a)
	bC := C.int(b)
	sum, err := C.sum(aC, bC)
	if err != nil {
		return 0, errors.New("error calling Sum function: " + err.Error())
	}
	//Convert C.int result to Go int
	res := int(sum)
	return res, nil
}
