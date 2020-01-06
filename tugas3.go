package main

import "fmt"

func main() {

	var buah = []string{"apel", "jeruk", "anggur", "melon"}
	buah = append(buah, "pepaya")

	fmt.Println("Jumlah Element : ", len(buah))
	fmt.Println("Isi Element : ", buah)

	for i, buah := range buah {
		fmt.Println("Element ke - : ", i, buah)
	}

}
