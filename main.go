package main

import (
	"fmt"
	"io/ioutil" // staticcheck: using deprecated package io/ioutil (SA1019)


	"os"
)

// unused: this function is unused
func unusedFunction() {
	// revive: empty-block: this block is empty
}

func anotherFunction() {
	// errcheck: error value not checked
	f, _ := os.Open("another-file")
	if f != nil {
		defer f.Close()
	}
}

func main() {
	// lll: this line is way too long and should trigger a linting error because it is more than 100 characters which is the limit that has been configured in the .golangci.yaml file.
	fmt.Println("Hello, linter!")

	var x int
	x = 1 // ineffassign: ineffectual assignment, x is reassigned before use.
	x = 2
	fmt.Println("value of x:", x)

	// errcheck: unchecked error
	os.Open("non-existent-file")

	// govet: wrong argument type in Printf
	fmt.Printf("This is a number: %s\n", 42)

	// asciicheck: non-ASCII character in identifier
	var nãme = "gopher"
	fmt.Println(nãme)

	// forcetypeassert: forced type assertion
	var i interface{} = "a string"
	_ = i.(string)

	var y = 5
	// staticcheck: self-assignment (SA4005)
	_ = y

	// staticcheck: using deprecated function
	ioutil.ReadFile("test.txt") // SA1019

	// govet: unreachable code
	return

}
