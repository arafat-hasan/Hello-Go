package main

import "fmt"

func main() {
	i := 2

	for i < 5 {
		fmt.Println(i)
		i++
	}

	for j := 30; j < 35; j++ {
		fmt.Println(j)
	}

	k := 98
	for {
		fmt.Println(k)
		if k == 100 {
			break
		}
		k++
	}

	for n := 200; n < 206; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
