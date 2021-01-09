package main

import (
	"algorithm_1/util"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Println("Input the number of socks in the pile")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err)
		return
	}
	if n < 1 || n > 100 {
		fmt.Println("number of socks must be: 1 ≤ n ≤ 100")

		return
	}
	fmt.Println("Input the colors of each sock, separated by space (10 20 20 10 10 30 50 10 20)")
	// socks, err := socksScanln(n)
	// if err != nil {
	// 	fmt.Println(err)
	// 	// return
	// }
	// fmt.Printf("test: %+v \n", socks)

	in := bufio.NewReader(os.Stdin)

	socksInput, err := in.ReadString('\n')
	fmt.Println(socksInput)

	socksSliceInput := strings.Fields(socksInput)
	if n != len(socksSliceInput) {
		fmt.Println("total number of socks must be: 0 ≤ i < numberOfSocks")
		return
	}
	socks, err := toSocksType(socksSliceInput)

	if err != nil {
		fmt.Println(err)
		return
	}
	totalPair := util.SockMerchant(n, socks)
	fmt.Printf("total pairs of socks: %+v \n\n", totalPair)

}

// func socksScanln(n)
func toSocksType(socksInput []string) (util.Socks, error) {
	fmt.Println(socksInput)
	result := make(util.Socks, len(socksInput))
	for i, sockInput := range socksInput {
		temp, err := strconv.Atoi(sockInput)
		if err != nil {
			return nil, err
		}
		if temp < 1 || temp > 100 {
			return nil, errors.New("value of socks must be: 1 ≤ ar [i] ≤ 100 ")
		}
		result[i] = util.Sock{Value: temp}
	}
	return result, nil
}
