package main

import "fmt"

func main() {
	// declaring variables without type using the 'var' keyword

	var numero = 100 // integer (there are many types of integer types in GO)
	var nombre = "Pedro Paramo" // para strings se usan double quotes
	var george = true // boolean
	var floating = 100.5 // float (there are float32 and float64)
	var myRune = 'A' // this is a rune (Unicode character, returns a UTF-( number))
	fmt.Println(numero)
	fmt.Println(nombre)
	fmt.Println(george)
	fmt.Println(floating)
	fmt.Println(myRune)
	fmt.Println("")

	// there are many types of integers in GO divided into Signed / Unsigned

	// SIGNED (4 categories)
	var numeroDos int8 = -109 // range -128 to 127
	var numeroTres int16 = 5637 // range -32768 to 32767
	var numeroCuatro int32 = -281738464 // range -2147483648 to 2147483647
	var numeroCinco int64 = 283748491375 // range -9223372036854775808 to 9223372036854775807
	fmt.Println("----SIGNED-----")
	fmt.Println("este es un int8: ", numeroDos)
	fmt.Println("este es un int16: ", numeroTres)
	fmt.Println("este es un int32: ", numeroCuatro)
	fmt.Println("este es un int64: ", numeroCinco)
	fmt.Println("")

	// UNSIGNED (4 categories) solo numeros positivos (y 0)
	var numeroSeis uint8 = 233 // range  0 to 255
	var numeroSiete uint16 = 4555 // range 0 to 65535
	var numeroOcho uint32 = 159384732 // range 0 to 4294967295
	var numeroNueve uint64 = 3574637281634 // range 0 to 18446744073709551615
	fmt.Println("----UNSIGNED-----")
	fmt.Println("este es un uint8: ", numeroSeis)
	fmt.Println("este es un int16: ", numeroSiete)
	fmt.Println("este es un int32: ", numeroOcho)
	fmt.Println("este es un int64: ", numeroNueve)





	
}