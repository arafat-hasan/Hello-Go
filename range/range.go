package main

import "fmt"

func main() {
	nums := []int{101, 201, 301, 401, 501}
	sum := 0
	for _, val := range nums {
		sum += val
	}
	fmt.Println(sum)

	for index, val := range nums {
		fmt.Println("index: ", index, "Value: ", val)
	}

	kvs := map[string]string{"a": "apple", "b": "banana", "c": "coconut"}

	for key, value := range kvs {
		fmt.Printf("%s  -> %s\n", key, value)
	}

	for byteIndex, rune := range "ABC" {
		fmt.Println(byteIndex, rune)
	}

	for byteIndex, rune := range "বাংলা" {
		fmt.Println(byteIndex, rune)
	}
}
