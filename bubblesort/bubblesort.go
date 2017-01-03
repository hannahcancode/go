//a bubblesort fuction in go

package main

import ()

func Bubblesort(numbers []int) []int {
	//numbers := n
	//fmt.Println(numbers)
	for range numbers {
		for i := range numbers {
			if i == len(numbers)-1 {
				break
			}
			if numbers[i] <= numbers[i+1] {
				continue
			}
			numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
		}
	}
	return numbers
}
