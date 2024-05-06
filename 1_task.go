package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	output(a)
}

func output(n int) {
	if 11 <= n%100 && n%100 <= 14 {
		fmt.Println(n, "компьютеров")
	} else {
		t := n % 10
		if t == 1 {
			fmt.Println(n, "компьютер")
		} else if 2 <= t && t <= 4 {
			fmt.Println(n, "компьютера")
		} else {
			fmt.Println(n, "компьютеров")
		}
	}
}
