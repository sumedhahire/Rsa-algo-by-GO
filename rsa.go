package main

import (
	"fmt"
	"math/rand/v2"
)

func coPrime() (int, int) {
	p := rand.IntN(100)
	for p%2 != 0 && p != 1 {
		p = rand.IntN(100)
	}
	q := rand.IntN(100)

	for i := p; i < q; i++ {
		if q%i == 0 && q%2 == 0 {
			q = rand.IntN(100)
		}
	}

	return p, q

}

// rsa
func main() {
	fmt.Println(coPrime())
}
