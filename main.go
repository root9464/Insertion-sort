// Сортировка вставками (Insertion sort) + метод рандомного массива
package main

import (
	"fmt"
	"math/rand"
)

func insertionSorted(array []int) []int {
	length := len(array)

	for i := 1; i < length; i++ {
		j := i //переопределние для обнуления счетчика
		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]

			}
			j = j - 1
		}

	}
	return array
}

func main() {
	len := int(5) // обьявление числа которое будет будущей длиной рандомного (неотсортированного) массива
	// создание рандомного (неотсортированного) массива
	array := rand.Perm(len)
	fmt.Println(array)

	// сортировка неотсортированного массива
	sort := insertionSorted(array)
	fmt.Println(sort)
}
