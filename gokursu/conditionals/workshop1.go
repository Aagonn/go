package conditionals

import "fmt"

func Workshop1() {
	var sayi1, sayi2, sayi3 int = 4523, 5213, 7342
	var big int = sayi1
	if sayi2 > big {
		big = sayi2
	}
	if big > sayi3 {
		big = sayi3
	}
	fmt.Printf("En büyük sayi : %v", big)
}
