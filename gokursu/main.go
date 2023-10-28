package main

/*
//.................................DEĞİŞKENLER VE VERİ TİPLERİ........................\\

func main() {
	fmt.Print("merhaba ")
	fmt.Print("dünya ")
	fmt.Println("!")
}

func main() {
	var metin string = "merhaba türkiye!" // "string" metin için kullanılır.
	fmt.Println(metin)
}

func main() {
	var kdv int = 18 // "int" tam sayılar için kullanılır
	fmt.Println(kdv)
	fmt.Println(100 + (100 * 10 / 100))

}

func main() {
	var otv float32 = 0.10 // "float32" ondalıklı sayılar için kullanılır
	fmt.Println(otv)
}

func main() {
	keltv := 25
	fmt.Printf("veri tipi : %T ", keltv) // veri tipini öğrenmek için "Printf" kullanıldı
	//%T virgülden sonraki değişkenin veri tipini göstermek için kullanıldı
}

func main() {
	keltv := 25.3 // işin kolayı için burada "var" ile atamak yerine ":=" bu işaretler ile hem tanımlıyorum hem de atıyorum anlamına gelen ifade kullanılıyor.
	fmt.Printf("veri tipi : %T ", keltv)
}

func main() {
	var durum bool
	var metin1 string = "araba"
	var metin2 string = "uçak"
	durum = metin1 == metin2 // "==" aynı mı ? sorusunu sorar.
	fmt.Println(durum)
}

func main() {
	var durum bool
	var metin1 string = "araba"
	var metin2 string = "uçak"
	durum = metin1 != metin2 // "!=" farklı mı ? sorusunu sorar.
	fmt.Println(durum)
}

func main() {
	var durum bool
	var metin1 string = "ev"
	var metin2 string = "araba"
	durum = metin1 == metin2
	fmt.Println(!durum) // durum değişkeninin önüne gelen "!" var olan durumun değilini gösterir.
}

//.................................MODÜLLER VE PAKETLER........................\\
import (
	"fmt"
	"golesson/variables"
)

func main() {
	variables.Demo1()
	fmt.Print()
}

//.................................ŞART BLOKLARI İLE ÇALIŞMAK........................\\

import (
	"golesson/conditionals"
)

func main() {
	conditionals.Demo1()
}

import "golesson/conditionals"

func main() {
	conditionals.Workshop1()
}

//.................................GÖNGÜLER İLE ÇALIŞMAK........................\\
import "golesson/loops"

func main() {
	loops.Demo1()
}

import "golesson/loops"

func main() {
	loops.Demo2()
}

import "golesson/loops"

func main() {
	loops.Demo3()
}

import "golesson/loops"

func main() {
	loops.Demo4()
}

import "golesson/loops"

func main() {
	loops.Demo5()
}

//.................................DİZİLER İLE ÇALIŞMAK........................\\
import "golesson/arrays"

func main() {
	arrays.Demo1()
}

import "golesson/arrays"

func main() {
	arrays.Demo2()
}

import "golesson/arrays"

func main() {
	arrays.Demo3()
}

import "golesson/arrays"

func main() {
	arrays.Demo4()
}

//.................................SLİCE YAPISIYLA ÇALIŞMAK........................\\
import "golesson/slices"

func main() {
	slices.Demo1()
}

import "golesson/slices"

func main() {
	slices.Demo2()
}

//.................................FONKSİYONLAR........................\\
import "golesson/functions"

func main() {
	functions.Selamver("Güzellik")
}

import "golesson/functions"

func main() {
	functions.Topla(2, 9)
}

import (
	"fmt"
	"golesson/functions"
)

func main() {
	sonuc1, sonuc2, sonuc3, sonuc4 := functions.Dortislem(27, 9)
	fmt.Println("Toplama : ", sonuc1)
	fmt.Println("Çıkarma: ", sonuc2)
	fmt.Println("Çarpma: ", sonuc3)
	fmt.Println("Bölme : ", sonuc4)
}

import (
	"fmt"
	"golesson/functions"
)

func main() {
	fmt.Println(functions.Toplavariadic(1, 4, 6, 3, 10))

	sayilar := []int{4, 6, 8, 2, 9}
	fmt.Println(functions.Toplavariadic(sayilar...))
}

//.................................MAP VE FOR-RANGE........................\\
import "golesson/maps"

func main() {
	maps.Demo1()
}

import "golesson/for_range"

func main() {
	for_range.Demo1()
}

import "golesson/for_range"

func main() {
	for_range.Demo2()
}

import "golesson/for_range"

func main() {
	for_range.Demo3()
}
*/
//.................................POİNTER VE STRUCT........................\\
