package util

func SockMerchant(n int, socks Socks) int {
	socksLen := len(socks)
	var totalPair int
	for i := 0; i < socksLen; i++ {
		if socks[i].IsPaired {
			continue
		}
		for j := i + 1; j < socksLen; j++ {
			if socks[i].Value == socks[j].Value && !socks[j].IsPaired {
				socks[i].IsPaired = true
				socks[j].IsPaired = true
				totalPair++
				break
			}
		}
	}
	return totalPair
}
