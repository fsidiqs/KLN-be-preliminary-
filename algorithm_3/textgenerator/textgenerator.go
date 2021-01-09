package textgenerator

import (
	"fmt"
	"strings"
)

func GenerateText(s string) string {
	length := len(s)
	var generated string
	index := 0
	for i := length - 1; i >= 0; i-- {
		if i == length-1 {
			generated = fmt.Sprintf("%v", padRightSide(string(s[index]), "0", i))

		} else {

			generated = fmt.Sprintf("%v\n%v", generated, padRightSide(string(s[index]), "0", i))
		}
		index++
	}

	return generated
}

func padRightSide(str string, item string, count int) string {
	return str + strings.Repeat(item, count)
}
