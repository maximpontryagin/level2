package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

// inputChannels - слайс каналов
func or(inputChannels ...<-chan interface{}) <-chan interface{} {
	// or позволяет блокировать программу возвращаемым каналом chSignal до получения сигнала из всех
	// переданных в функцию каналов. Может быть полезна для синхронизации горутин.
	chSignal := make(chan interface{})
	defer close(chSignal)

	var wg sync.WaitGroup
	for _, channel := range inputChannels {
		wg.Add(1)
		go func(ch <-chan interface{}) {
			for i := range ch {
				chSignal <- i
			}
			wg.Done()
		}(channel)
	}
	wg.Wait()
	return chSignal
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		// sig(2*time.Hour),
		// sig(5*time.Minute),
		sig(1*time.Second),
		sig(4*time.Second),
		sig(9*time.Second),
		// sig(1*time.Hour),
		// sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
