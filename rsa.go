package main

import (
	"fmt"
	"math/rand/v2"
	_ "strings"
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

func rsa(str string) string {
	p, q := coPrime()
	fmt.Println(p, q)
	n := p * q
	z := (p - 1) * (q - 1)

	e := rand.IntN(z-2) + 2
	d := (e - 1) % (p - 1) * (q - 1)

	_ = d
	_=n
	//var cipher string=""
	// for  _,i := range str {
	// 	cipher+=string(i)
	// }
	//its like getting ascii so have to cast to string
	cipher:=[]rune(str)
	return string(cipher)
}

// rsa
func main() {
	//fmt.Println(coPrime())
	fmt.Println(rsa("sumedh"))
}
