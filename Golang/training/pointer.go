package training

import "fmt"

func Pointer() {

	name := "Yusuf"

	p := &name
	fmt.Println("Bellek adresi : ", p)   // bellek adresi
	fmt.Println("Adresin değeri : ", *p) // adresin değeri
	*p = "Emir"                          // pointer ile müdahele
	fmt.Println("Pointer ile müdahele sonucu name : ", name)
	changeName(p)
	fmt.Println("Fonksiyon ile pointer kullanımı sonucu : ", name)

}

// Fonksiyon kullanımı
func changeName(m *string) {

	*m = "Mehmet"

}
