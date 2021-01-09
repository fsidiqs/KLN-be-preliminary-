package main

import (
	"counting_valley/countingvalley"
	"fmt"
)

func main() {

	var n int
	fmt.Println("Input the number of steps Gary takes")
	fmt.Scanf("%d", &n)
	if n < 2 || n > 1000000 {
		fmt.Println("number of steps must be: 2 < n <= 10^6")
		return
	}
	var s string
	fmt.Println("Input a string describing his path")

	fmt.Scanf("%s", &s)
	if len(s) != n {
		fmt.Println("number of steps must be: 2 < n <= 10^6")
		return
	}
	isValid := true
	for i := range s {
		if string(s[i]) != "u" && string(s[i]) != "d" {
			isValid = false
		}
	}
	if !isValid {
		fmt.Printf("invalid path")
		return
	}
	fmt.Printf("Output: \n%+v\n", countingvalley.CountingValley(n, s))
}
