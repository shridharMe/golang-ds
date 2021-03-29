package main

import "fmt"

type ListInt struct {
	data []int
}

func (L *ListInt) bubbleSorting() {
	fmt.Println("\nBefore Bubble Sorting")
	fmt.Println(L.data)
	aLen := len(L.data)
	for i := 0; i < aLen-1; i++ {
		for j := 0; j < aLen-1; j++ {
			if L.data[j] > L.data[j+1] {
				L.data[j], L.data[j+1] = L.data[j+1], L.data[j]
			}
		}
	}
	fmt.Println("\nAfter Bubble Sorting")
	fmt.Println(L.data)
}

func main() {
	link := ListInt{
		data: []int{23, 12, 1, 56, 10, 5, 100},
	}
	link.bubbleSorting()
}
