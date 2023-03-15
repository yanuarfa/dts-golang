package main

import "fmt"

func main() {
	var i = 0

	for i < 5 {
		fmt.Println("Nilai i =", i)
		i++
	}

	for j := 0; j <= 10; j++ {
		if j == 5 {
			var russian = []rune{}
			russian = append(russian, 'С', 'А', 'Ш', 'А', 'Р', 'В', 'О')
			pos := 0
			for _, r := range russian {
				fmt.Printf("characters %U '%c' start at byte position %d\n", r, r, pos)
				pos += 2
			}
		} else {
			fmt.Println("Nilai j =", j)
		}
	}
	fmt.Scanln()
}
