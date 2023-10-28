package arrays

import "fmt"

func Demo3() {
	sayilar := [5]int{20, 30, 40, 50, 60}
	fmt.Println(sayilar)
	big := sayilar[0]
	for i := 0; i < len(sayilar); i++ { //length = uzunluk
		if sayilar[i] > big {
			big = sayilar[i]
		}
	}
	fmt.Println("en büyük sayi : ", big)
}
