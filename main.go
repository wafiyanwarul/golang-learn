package main

import "fmt"

func main() {

	fmt.Println("Kelebihan Golang : \n1. Memiliki Garbage Collector \n2. Memiliki Concurrency \n3. Lebih Cepat \n4. Mudah Dipelajari \n5. Didukung Banyak Framework")
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
	fmt.Println("This is", a, b, c, d, "and", e)

	//dapat berbeda tipe data
	one, isMonday, twoPointTwo, say := "1", true, 2.2, "Hello" //beda tipe data dalam satu deklarasi
	fmt.Println(one, isMonday, twoPointTwo, say)

	/*variable underscore
	  (_) underscore adalah reserved variable 
	  untuk menampung variabel yang tidak digunakana agar tidak terjadi error
	*/

	_ = "Learn Golang"
	_ = "Golang is Easy"
	name, _ := "John", "Dalton"

	fmt.Println(name) // variabel underscore tidak dipanggil disini

	//variable declaration using keyword NEW
	// name1 = new(string) // variabel

	// fmt.Println(name1) // 0x20818a220
    // fmt.Println(*name1) // ""

	

}
