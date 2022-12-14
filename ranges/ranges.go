package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Just iterate over keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Range on string iteraqtaes over Unicode code points.
	// Value will return as rune
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
