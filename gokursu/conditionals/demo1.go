package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekistenen float64 = 1900

	if cekilmekistenen > hesap {
		fmt.Println("Hesaptaki parayetersiz.")
	}
	if cekilmekistenen <= hesap {
		fmt.Println("paranız hazırlanıyor...")
	}
	fmt.Printf("Bitti. hesaptaki para : %F", hesap)
}
