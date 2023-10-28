package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekistenen float64 = 1000
	if cekilmekistenen > hesap {
		fmt.Println("hesaptaki para yetersiz")
	} else if cekilmekistenen == hesap {
		fmt.Println("paranız hazırlanıyor\nparanız kalmadı.") // burada kullanılan "\n" aynı satırda yazılan iki cümle için alt satıra geç komutu olarak kullanılır.
	} else {
		fmt.Println("paranız hazırlanıyor.")
	}
}
