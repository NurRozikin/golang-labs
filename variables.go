package main

import "fmt"

func main() {

	/*/ ################# Deklarasi Variabel Dengan Tipe Data ################# /*/
	fmt.Println("\n # Deklarasi Variabel Dengan Tipe Data")
	var a string = "Kopral"
	var b string
	b = "Jono"

	fmt.Printf("hallo %s %s!\n", a, b)

	/*/	################# Deklarasi Variabel Tanpa Tipe Data ################# /*/
	fmt.Println("\n # Deklarasi Variabel Tanpa Tipe Data")
	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	var c = "john"

	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	d := "wick"

	fmt.Printf("hallo %s %s!\n", c, d)

	// Tanda := hanya digunakan sekali di awal pada saat deklarasi. Untuk assignment nilai selanjutnya harus menggunakan tanda = , contoh:
	e := "wick"
	e = "ethan"
	e = "bourne"

	fmt.Printf("hallo %s !\n", e)

	/*/ ################# Deklarasi Multi Variabel ################# /*/
	fmt.Println("\n # Deklarasi Multi Variabel")

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fmt.Println("first, second, third = ", first, second, third)

	var fourth, fifth, sixth string = "empat", "lima", "enam"
	fmt.Println("fourth, fifth, sixth = ", fourth, fifth, sixth)

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"
	fmt.Println("seventh, eight, ninth = ", seventh, eight, ninth)

	// deklarasi multi variabel bisa dilakukan untuk variabel-variabel yang tipe data satu sama lainnya berbeda
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Println("one, isFriday, twoPointTwo, say = ", one, isFriday, twoPointTwo, say)

	/*/ ################# Variabel Underscore _ ################# /*/
	fmt.Println("\n # Variabel Underscore _")
	/*
		Underscore ( _ ) adalah predefined variabel yang bisa dimanfaatkan untuk menampung nilai yang tidak dipakai.
		Golang memiliki aturan unik yang jarang dimiliki bahasa lain, yaitu tidak boleh ada satupun variabel yang menganggur. Artinya, semua variabel yang dideklarasikan harus digunakan.
		Jika ada variabel yang tidak digunakan tapi dideklarasikan, error akan muncul dan program tidak bisa di-run ataupun di-compile.
	*/
	_ = "belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "john", "wick"
	fmt.Printf("hallo %s !\n", name)

	/*/ ################# Deklarasi Variabel Menggunakan Keyword 'new' ################# /*/
	fmt.Println("\n # Deklarasi Variabel Menggunakan Keyword new")
	//Keyword new digunakan	untuk mencetak data pointer dengan tipe data tertentu.	Nilai data default-nya akan menyesuaikan tipe datanya.

	nama := new(string)
	fmt.Println(nama)  // 0x20818a220
	fmt.Println(*nama) // ""

	//fmt.Println("\n # Deklarasi Variabel Menggunakan Keyword make")
	//Keyword ini hanya bisa digunakan untuk pembuatan beberapa jenis variabel saja, yaitu: channel, slice, map

}
