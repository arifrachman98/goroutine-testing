package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Test")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}

/*
	Go Routine tidak cocok digunakan pada function yang mengembalikan nilai (function), karna nilai tersebut tidak akan di tangkap oleh goroutine
	Akan tetapi, goroutine sangat cocok digunakan pada function yang tidak mengembalikan nilai (method)
	Untuk fuction yang mengembalikan nilai dengan goroutine, maka gunakan channel sebagai solusinya
*/

/*
	Channel sendiri merupakan tempat komunikasi secara synchronous yang bisa dilakukan oleh goroutine
	Pada channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine yang berbeda
	Saat melakukan pengiriman data ke channel, goroutine akan ter-block sampai ada yang menerima data tersebut
	Maka dari itu, channel disebut sebagai alat komunikasi synchronous (blocking) antar goroutine
	Channel dapat digunakan sebagai alternatif, seperti mekanisme async await yang terdapat pada bahasa pemrograman lain
*/

/*
	Karakteristik dari channel:
	- hanya menampung satu data dan satu tipe data
	- dapat diambil oleh lebih dari satu goroutine
	- jika tidak digunakan, channel sebaiknya di close agar tidak menyebabkan memory leak
*/

/*
	Channel pada golang direpresentasikan dengan tipe data chan
	ketika membuat Channel, kita dapat meng implementasikan nya dengan 'make()'
	Lalu, saat membuat channel, kita harus menentukan tipe data yang akan digunakan pada channel tersebut
*/
