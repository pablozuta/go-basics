// An array in Go is a fixed-length data type that contains a contiguous
// block of elements of the same type.
// This could be a built-in type such as integers and strings
// or it can be a struct type.

package main

import "fmt"

func main() {
	// declaring arrays
	var arrayNumeros [5]int
	arrayNumeros[3] = 2666

	primes := [6]int {2, 3, 5, 7, 11, 13}

	// array de strings
	var arrayStrings [3]string
	arrayStrings[0] = "Dave"
	arrayStrings[1] = "Pamela"
	arrayStrings[2] = "George"

	// array de booleans
	var arrayBooleans [2]bool
	arrayBooleans[0] = true
	arrayBooleans[1] = false

	fmt.Println(arrayStrings)
	fmt.Println(arrayNumeros)
	fmt.Println(arrayBooleans)
	fmt.Println(primes)

	// declaring arrays using array literals
	arrayNumerosLiteral := [4]int {10, 30, 32, 643}
	fmt.Println(arrayNumerosLiteral)

	arrayStringsLiteral := [5]string {"Rethoric of David", "Oath of the Horatti", "The Parallel Lives", "Historie Ancienne", "Brutus Condemming his Sons to Death"}
	fmt.Println(arrayStringsLiteral[0])
	fmt.Println(arrayStringsLiteral[1])
	fmt.Println(arrayStringsLiteral[2])
	fmt.Println(arrayStringsLiteral[3])
	fmt.Println(arrayStringsLiteral[4])

	// initialize specific index of an array
	x := [5]int {1:34, 3:60} 
	fmt.Println(x)

	// looping por un array
	for i:= 0; i < len(x); i++ {
		fmt.Println("looping por el array: ",x[i])
	}

	j := 0
	for range x {
		fmt.Println(x[j])
		j++
	}

	counter := 0
	for range arrayStrings {
		fmt.Println(arrayStrings[counter])
		counter ++
	}

	for i:= 0; i < len(arrayStrings); i++ {
		fmt.Println(arrayStrings[i])
	}
}

// Arrays are valuable data structures because the memory is allocated sequentially.
// Having memory in a contiguous form can help to keep the memory you use 
// stay loaded within CPU caches longer.
// Using index arithmetic, you can iterate through all the elements
// of an array quickly.
// The type information for the array provides the distance in memory 
// you have to move to find each element.
// Since each element is of the same type and follows each other sequentially
// moving through the array is consistent and fast.