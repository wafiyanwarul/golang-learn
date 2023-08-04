package main

import "fmt"

func main() {

	fmt.Println("Kelebihan Golang : \n1. Memiliki Garbage Collector \n2. Memiliki Concurrency \n3. Lebih Cepat \n4. Mudah Dipelajari \n5. Didukung Banyak Framework")
	/*
	   ini adalah komentar di GO language
	   ini adalah tipe komentar multiline
	*/
	fmt.Println("      \nHello, world!")
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
	last := "Dweck"            // tanpa tipe data menggunakan perantara (:=)

	fmt.Printf("Welcome, %s %s!\n", first, last)

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
	fmt.Println("\nTipe data berbeda dalam satu deklarasi : \n", one, isMonday, twoPointTwo, say)

	/*variable underscore
	  (_) underscore adalah reserved variable
	  untuk menampung variabel yang tidak digunakana agar tidak terjadi error
	*/

	_ = "Learn Golang"
	_ = "Golang is Easy"
	name, _ := "John", "Dalton"

	fmt.Println(name) // variabel underscore tidak dipanggil disini

	// variable declaration using keyword NEW
	var name1 = new(string) // variabel

	fmt.Println(name1)  // 0x20818a220
	fmt.Println(*name1) // "angka di atas adalah alamat memori variabel name1 disimpan"

	//VARIABLE DECLARE USING "MAKE" KEYWORD
	fmt.Print()
	fmt.Println("'int' & 'uint'")
	fmt.Println("'int'  :: bilangan bulat (-) & (+)")
	fmt.Println("'uint' :: bilangan cacah positif saja (+)")

	//NUMERIC NON-DECIMAL
	//UINT //semua nilai variabel bertipe data uint dimulai dari "0" ke positif
	var positif uint = 18446744073709551615 //bisa sama dengan uint32 maupun uint64
	var positif0 uint8 = 255
	var positif1 uint16 = 65535
	var positif2 uint32 = 4294967295
	var positif3 uint64 = 18446744073709551615
	var positif4 byte = 255 // tipe data byte sama dengan uint8
	fmt.Println("\nuint ", positif, "\nuint8 ", positif0, "\nuint16 ", positif1, "\nuint32 ", positif2, "\nuint64 ", positif3, "\nbyte ", positif4)

	//INT //semua nilai variabel bertipe data int menampung nilai dari negatif sampai dengan positif(akhiran 7)
	var ca int = -9223372036854775808    //sama dengan int32 maupun int64
	var ca0 int8 = -128                  //sampai dengan 127
	var ca1 int16 = -32768               //sampai 32767
	var ca2 int32 = -2147483648          //sampai dengan 2147483647
	var ca3 int64 = -9223372036854775808 //sampai dengan 9223372036854775807
	var ca4 rune = -2147483648           //sama dengan int32
	fmt.Println("\nint ", ca, "\nint8 ", ca0, "\nint16 ", ca1, "\nint32 ", ca2, "\nint64 ", ca3, "\nrune ", ca4)

	var neg = -1243423644                        //tanpa deklarasi tipe data
	fmt.Printf("\nBilangan negatif : %d\n", neg) //compiler dengan cerdas akan menentukan tipe data int32
	//karena masuk ke cakupan nilai dari int32
	//template "%d" padan "fmt.Printf" digunakan untuk memformat data numerik non-desimal.

	//NUMERIC DECIMAL
	//tipe data numerik decimal hanya ada float32 dan float64
	var decNum = 2.62 //otomatis bertipe data float32 karena sesuai dengan cakupan
	fmt.Printf("\nBilangan decimal 1 : %f\n", decNum)

	// template %f digunakan untuk memformat data numerik desicimal menjadi string
	// akan menghasilkan 6 digit, seperti pada hasil pertama 2.620000
	// jumlah digit yang muncul dapat dikontrol dengan "%.nf" (ganti n dengan nilai yang diinginkan)
	// contohnya pada kode di bawah ini : yaitu dengan "%.9f" maka muncul 9 digit dibelakang titik

	fmt.Printf("Bilangan decimal 2 : %.9f\n", decNum) 

	//BOOLEAN ( TRUE / FALSE )
	var bb bool = false
	fmt.Printf("\nIni boolean : %t\n", bb)	
	// template "%t" digunakan untuk memformat tipe data boolean pada "fmt.Printf"

	//STRING ( ditandai ",',`,baris baru")
	var pesan string = `Ini adalah "pesan" 
	pesan ini memiliki
	tipe data yaitu
	tipe data string` // pada variable tipe string ini ditandai dengan tana (`) backticks yang berguna untuk membuat kalimat pada baris berikutnya yang masih di dalam tanda backtick terdeteksi sebagai tipe data string

	fmt.Printf("\nPesan %s\n", pesan) 
	// template "%s" untuk memformat tipe data string

	//NIL & ZERO VALUE (nilai kosong = nil)
	var qA string
	var qB bool 
	var qC int
	var qD = 0.
	fmt.Printf("\nZero Value String %s\n", qA) //""
	fmt.Printf("Zero Value Boolean %t\n", qB) //false
	fmt.Printf("Zero Value Numerik Non-Decimal %d\n", qC) //0
	fmt.Printf("Zero Value Numerik Decimal %.1f\n", qD) //0.0

	//nil tidak dapat digunakan pada tipe data pointer, tipe data fungsi, slice, map, channel, interface kosong atau any (yang merupakan alias dari interface{}


	// CONSTANTS

	const angka uint8 = 10 //constants hanya bisa dimasuki nilai sekali saja
	//angka := 12 //jika kode disamping di uncomment
	//maka akan muncul peringatan

	fmt.Print(angka) //cetak nilai constant

	var number int = 10 //nilai variabel dapat diubah
	number = 15         // di sini nilainya dapat diubah karena dia bukan constants melainkan variabel

	fmt.Print(number) //cetak nilai variabel

	const First string = "Wafiy"
	fmt.Println("\nHi ", First, "!") //cetak

	//multi constant declaration
	const (
		triangle = "segitiga"
		isSunday bool = true
		numeric uint8 = 7
		float1 = 3.7
	)

	//contoh lagi
	const (
		yaa = "ini kata 'ya'"
		yee //pada konteks ini, konstanta yang tidak diisi nilainya, maka tipe data dan nilainya akan disamakan dengan konstanta yang dideklarasikan sebelumnya
	)

	const (
		Today string = "Senin"
		hariIni
		Today1 = true
	)

	fmt.Println("Apakah hari ini hari", Today, "?", Today1)

	//deklarasi constant dalam satu baris
	
}
