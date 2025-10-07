package main

import (
	"fmt"
	"io/ioutil" // staticcheck: using deprecated package io/ioutil (SA1019)
	"net/http"
	"os"
)

var unusedVar = ""

var unusedVarTwo = ""

var unusedVarThree = ""

// unused: this function is unused
func unusedFunction() {
	// revive: empty-block: this block is empty
}

// cyclop: function with high cyclomatic complexity
func highComplexity(a int) {
	if a == 1 {
		fmt.Println("one")
	} else if a == 2 {
		fmt.Println("two")
	} else if a == 3 {
		fmt.Println("three")
	} else if a == 4 {
		fmt.Println("four")
	} else if a == 5 {
		fmt.Println("five")
	} else if a == 6 {
		fmt.Println("six")
	} else if a == 7 {
		fmt.Println("seven")
	} else if a == 8 {
		fmt.Println("eight")
	} else if a == 9 {
		fmt.Println("nine")
	} else {
		fmt.Println("other")
	}
}

// bodyclose: response body must be closed
func bodyCloseError() {
	resp, err := http.Get("https://example.com")
	if err != nil {
		return
	}
	// Missing resp.Body.Close()
	fmt.Println(resp.Status)
}

// nakedret: naked return
func nakedReturn(a int) (ret int) {
	ret = a * 2
	return
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

	// goconst: string literal repeated
	fmt.Println("repeated string")
	fmt.Println("repeated string")

	// prealloc: slice append in loop
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}

	// govet: unreachable code
	return

}
