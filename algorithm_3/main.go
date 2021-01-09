package main

import (
	"algorithm_3/textgenerator"
	"fmt"
	"regexp"
)

func main() {
	var input string
	fmt.Println("input some numbers")
	fmt.Scanf("%v", &input)
	matched, _ := regexp.MatchString(`^[0-9]*$`, input)
	if !matched {
		fmt.Println("only numeric input is allowed")
		return
	}
	result := textgenerator.GenerateText(input)
	fmt.Println(result)
}
