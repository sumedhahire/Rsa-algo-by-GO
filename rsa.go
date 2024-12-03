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

func rsa(str string) (string,string) {
	p, q := coPrime()
	fmt.Println(p, q)
	n := p * q
	z := (p - 1) * (q - 1)

	e := rand.IntN(z-2) + 2
	if 1 < e && e < z{
		fmt.Println("yes")
	}else{
		return "",""
	}

	d := (e - 1) % (p - 1) * (q - 1)



	txt:=[]rune(str)
	cipher:=""
	for i:=range txt{
		cipher+=string((i*e)%n)
	}

	txtToD:=[]rune(cipher)
	plaintxt:=""

	for i:= range txtToD{
		plaintxt+=string((i*d)%n)
	} 

	return cipher,plaintxt
}

// rsa
func main() {
	//fmt.Println(coPrime())
	fmt.Println(rsa("sumedh"))
}
