package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

const protocol = "tcp"

type telnetConfig struct {
	timeout     string
	host        string
	port        string
	timeoutTime time.Duration
}

func (cfg *telnetConfig) parse() error {
	flag.StringVar(&cfg.timeout, "timeout", "10s", "connection timeout")
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		return fmt.Errorf("необходимо указать хост и порт")
	}
	cfg.host = args[0]
	cfg.port = args[1]

	timeoutTime, err := time.ParseDuration(cfg.timeout)
	if err != nil {
		return err
	}
	cfg.timeoutTime = timeoutTime
	return nil
}

func telnet(ctx context.Context, timeout time.Duration, address string) error {
	conn, err := net.DialTimeout(protocol, address, timeout)
	if err != nil {
		return fmt.Errorf("ошибка подключения по tcp: %v", err)
	}
	// Закрываем соединение
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Ошибка при закрытии соединения: %v", err)
		}
	}()

	done := make(chan struct{})

	go func() {
		copyTo(os.Stdout, conn)
		done <- struct{}{}
	}()
	go func() {
		copyTo(conn, os.Stdin)
		done <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		log.Println("Получен сигнал прерывания, соединение закрывается...")
	case <-done:
		log.Println("Соединение закрыто сервером или ввод завершен.")
	}
	return nil
}

func copyTo(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Printf("ошибка в копировании данных: %v", err)
	}
}

func main() {
	var cfg telnetConfig
	if err := cfg.parse(); err != nil {
		log.Fatalf("Ошибка при парсинге аргументов: %v", err)
	}

	address := fmt.Sprintf("%s:%s", cfg.host, cfg.port)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer stop()

	if err := telnet(ctx, cfg.timeoutTime, address); err != nil {
		log.Fatalf("Ошибка в работе telnet клиента: %v", err)
	}

	<-ctx.Done()
}
