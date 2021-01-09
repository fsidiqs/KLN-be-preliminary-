package util

type Sock struct {
	Value    int
	IsPaired bool
}

type Socks []Sock

// func (socks *Socks) removeSock(idx int) {
// 	// var tempSocks Socks

// 	copy((*socks)[idx:], (*socks)[idx+1:])
// 	(*socks)[len(*socks)-1] = 0
// 	(*socks) = (*socks)[:len(*socks)-1]
// }
