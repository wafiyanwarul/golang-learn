package main

import "fmt"

func main() {
	/*
	   ini adalah komentar di GO language
	   ini adalah tipe komentar multiline
	*/
	fmt.Println("Hello, world!")
	fmt.Println("This is GO Language!")

	// fmt.Println("Pesan ini tidak akan dieksekusi")
	// fmt.Println("ini adalah tipe komentar inline")

	// VARIABLES
	/*
		var <nama-variabel> <tipe-data> = <nilai>
	*/
	var firstname string = "Wafiy Anwarul"
	var lastname string = "Hikam"

	// PRINTF
	fmt.Printf("Hello, Joe Willy!\n")
	fmt.Printf("Hello, %s %s!\n", firstname, lastname)
	fmt.Println("Hello, my name is", firstname, lastname+"!")

	// VAR WITHOUT DATA TYPE
	var first string = "Carol" // dengan tipe data menggunakan perantara (=)
	last := "Dweck" // tanpa tipe data menggunakan perantara (:=)

	fmt.Println("Welcome, %s %s!\n", first, last)

	// MULTI VARIABLE DECCLARATIONS
	var first1, second1, third1 string
	first1, second1, third1 = "one", "two", "three"

	fmt.Println(first1, second1, third1)

	var a, b, c string = "1", "2", "3"
	fmt.Println(a, b, c)

	//lebih ringkas 
	a, b, c, d, e := "1", "2", "3", "4", "5"
	fmt.Println("This is", a, "and", b, "is", c, d, e)

	
}
