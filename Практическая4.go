package main
import (
	"fmt"
	"sync"
	"time"
	"net/http"
	"crypto/md5"
	"io"
	"os"
)

func main() {

// Задача 1

var w sync.WaitGroup
	w.Add(1)
	go func() {
		defer w.Done()
		for i := 1; i <= 5; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}()
	w.Wait()
	fmt.Println()

// Задача 2 

j := make(chan int, 10)
r := make(chan int, 10)

	var w2 sync.WaitGroup
	for i := 0; i < 3; i++ {
		w2.Add(1)
		go func() {
			defer w2.Done()
			for x := range j {
				r <- x * x
			}
		}()
	}
	for i := 1; i <= 10; i++ {
		j <- i
	}
	close(j)
	go func() {
		w2.Wait()
		close(r)
	}()
	for a := range r {
		fmt.Print(a, " ")
	}
	fmt.Println("\n")

//Задача 3

	tick := time.Tick(200 * time.Millisecond)
	for i := 1; i <= 15; i++ {
		<-tick
		fmt.Println("Запрос", i, "обработан")
	}
	fmt.Println()

//Задача 4

u := []string{
	"https://store.steampowered.com",
	"https://www.dotabuff.com",
	"https://github.com",
}
j1 := make(chan string, len(u))
r1 := make(chan string, len(u))
var w4 sync.WaitGroup
	for i := 0; i < 3; i++ {
		w4.Add(1)
		go func() {
			defer w4.Done()
			for url := range j1 {
				resp, err := http.Get(url)
				if err != nil {
					r1 <- url + " ошибка"
					continue
				}
				r1 <- url + " " + resp.Status
				resp.Body.Close()
			}
		}()
	}
	for _, x := range u {
		j1 <- x
	}
	close(j1)
	go func() {
		w4.Wait()
		close(r1)
	}()
	for v := range r1 {
		fmt.Println(v)
	}
	fmt.Println()

//Задача 5

file := []string{"Практическая4.go"}
j2 := make(chan string, len(file))
r2 := make(chan string, len(file))
var w5 sync.WaitGroup
for i := 0; i < 3; i++ {
	w5.Add(1)
	go func() {
		defer w5.Done()
		for f := range j2 {
			fh, err := os.Open(f)
			if err != nil {
				r2 <- f + ": ошибка"
				continue
			}
			h := md5.New()
			io.Copy(h, fh)
			fh.Close()
			r2 <- fmt.Sprintf("%s: %x", f, h.Sum(nil))
		}
	}()
}
for _, x := range file {
	j2 <- x
}
close(j2)
go func() {
	w5.Wait()
	close(r2)
}()
for v := range r2 {
	fmt.Println(v)
}
fmt.Println()
}