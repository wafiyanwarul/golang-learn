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

	

}
