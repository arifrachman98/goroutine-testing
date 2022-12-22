package belajargolanggoroutine

import (
	"fmt"
	"strconv"
	"sync"
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
		fmt.Println("Berhasi melakukan", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	chann1 := make(chan string)
	chann2 := make(chan string)
	defer close(chann1)
	defer close(chann2)

	go OnlyIn(chann1)
	go GiveResponse(chann2)

	count := 0
	for {
		select {
		case data := <-chann1:
			fmt.Println(data)
			count++
		case data := <-chann2:
			fmt.Println(data)
			count++
		default:
			println("Menunggu Data")
		}

		if count == 2 {
			break
		}
	}

}

func TestRaceCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter : ", x)
}

type BankAccount struct {
	RWmutex sync.RWMutex
	saldo   int
}

func (akun *BankAccount) AddBalance(jumlah int) {
	akun.RWmutex.Lock()
	akun.saldo = akun.saldo + jumlah
	akun.RWmutex.Unlock()
}

func (akun *BankAccount) GetBalance() int {
	akun.RWmutex.RLock()
	saldo := akun.saldo
	akun.RWmutex.RUnlock()
	return saldo
}

func TestRWmutex(t *testing.T) {
	akun := BankAccount{}

	for i := 1; i < 1000; i++ {
		go func() {
			for j := 1; j < 100; j++ {
				akun.AddBalance(2)
				fmt.Println(akun.GetBalance())
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Saldo Total :", akun.GetBalance())
}
