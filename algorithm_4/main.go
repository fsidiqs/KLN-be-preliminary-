package main

import "fmt"

func main() {

	var n int
	fmt.Println("input total switches:")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err)
		return
	}
	switches := make([]int, n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j%i == 0 {
				if switches[j-1] == 0 {
					switches[j-1] = 1
				} else {
					switches[j-1] = 0
				}
			}
		}
	}
	result := sumArray(switches)
	fmt.Printf("result: %+v \n\n", result)
}

func sumArray(slices []int) int {
	result := 0
	for _, v := range slices {
		result += v
	}
	return result
}
