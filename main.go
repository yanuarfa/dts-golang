package main

import "fmt"

func main() {
	kalimat := "selamat malam"
	jumlah := make(map[string]int)

	for _, h := range kalimat {
		fmt.Printf("%c\n", h)
		jumlah[string(h)] += 1
	}

	fmt.Print(jumlah)
	fmt.Scanln()
}
