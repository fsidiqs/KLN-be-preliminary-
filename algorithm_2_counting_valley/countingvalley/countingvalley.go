package countingvalley

import (
	"strings"
)

func CountingValley(n int, s string) int {
	valleys, altitude := 0, 0

	s = strings.ToLower(s)

	for i := 0; i < n; i++ {
		step := string(s[i])
		if altitude == -1 && step == "u" {
			valleys++
		}
		var temp1, temp2 int
		if step == "d" {
			temp1 = 1
		}
		if step == "u" {
			temp2 = 1
		}
		altitude += -1*temp1 + 1*temp2
	}
	return valleys
}
