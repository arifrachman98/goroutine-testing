package belajargolanggoroutine

import (
	"fmt"
	"strconv"
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

func TestCreateChannel(t *testing.T) {
	channel := make(chan string) //Membuka channel baru
	defer close(channel)         //Menutup Channel dan dieksekusi terakhir kali dengan defer

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Arif Rachman Hakim" //Mengirim data ke channel
		fmt.Println("Pengiriman data ke channel selesai")
	}()

	data := <-channel //Menerima data dari channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveResponse(chann chan string) {
	time.Sleep(2 * time.Second)
	chann <- "You get the response"
}

func TestChannelAsParameter(t *testing.T) {
	chen := make(chan string)
	defer close(chen)

	go GiveResponse(chen)

	data := <-chen
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(chann chan<- string) {
	time.Sleep(1 * time.Second)
	chann <- "Arif Rachman Hakim"
}

func OnlyOut(chann <-chan string) {
	data := <-chann
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	chenn := make(chan string)
	defer close(chenn)

	go OnlyIn(chenn)
	go OnlyOut(chenn)

	time.Sleep(1 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	chann := make(chan string, 3)
	defer close(chann)

	for i := 1; i <= cap(chann); i++ {
		chann <- "Hahaha"
	}

	time.Sleep(2 * time.Second)
	for i := 1; i <= cap(chann); i++ {
		fmt.Println("Data ke-", i, <-chann)
	}

	fmt.Println("Selesai")
	time.Sleep(3 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	chann := make(chan string)

	go func() {
		for i := 1; i <= 10; i++ {
			chann <- "Perulangan ke-" + strconv.Itoa(i)
		}
		close(chann)
	}()

	for data := range chann {
		fmt.Println(data)
	}
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

/*
	Jika channel sudah mengirimkan data tapi tidak ada penerimanya, maka akan terjadi blocking pada program
	Jika ada penerima pada channel akan tetapi tidak terdapat pengirimnya, maka akan terjadi deadlock pada program
	Untuk mencegah hal tersebut, diperlukan langkah yang tepat sebagai berikut :
	1. Pastika kita membuat Channel
	2. Tentukan tipe data yang akan diterima oleh channel
	3. Pastikan channel yang kita buat memiliki pengirim dan penerima
	4. Setelah proses pengiriman dan penerimaan data pada channel, pastikan channel tertutup
*/

/*
	Channel dapat diatur sebagai IN(mengirim data) saja atau sebagai OUT(menerima data) saja dengan cara menandainya pada parameter.
	Example :
	Mengirim data :	func OnlyIn(chann chan<- string)
	Menerima data : func OnlyOut(chann <-chan string)
*/

/*
	Ada kondisi dimana channel dapat menerima data lebih banyak dibandingkan mengirim data yang ada pada channel itu sendiri,
	dengan demikian pengiriman data pada channel akan ikut melambat. Maka dari itu diperlukan Buffered Channel untuk meng antisipasi hal demikian.
	Buffered channel sendiri merupakan sebuah wadah yang digunakan untuk menampung pengiriman data yang menumpuk pada channel
*/

/*
	Ada kondisi dimana sebuah channel menerima data secara terus menerus dari pengirim dan kita tidak mengetahui kapan channel tersebut akan berhenti menerima data.
	Hal yang dapat kita lakukan untuk menanggulangi hal tersebut adalah dengan menggunakan perulangan range ketika menerima data dari channel.
	Ketika channel di close(), maka secara otomatis perulangan tersebut akan terhenti secara otomatis.
	Dengan cara ini lebih sederhana dari pada kita melakukan pengecekan channel secara manual. Cara ini disebut sebagai Range Channel
*/
