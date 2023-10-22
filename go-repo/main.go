package main

/*
//.................................DEĞİŞKENLER VE VERİ TİPLERİ........................\\
// DEĞŞKENLER:
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
*/
//.................................MODÜLLER VE PAKETLER........................\\
import (
	"fmt"
	"golesson/variables"
)

func main() {
	variables.Demo1()
	fmt.Print()
}
