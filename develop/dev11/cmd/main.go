package main

import (
	"context"
	"dev11/internal/config"
	"dev11/internal/server"
	"dev11/internal/storage"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

func main() {
	cfg, err := config.NewConfigServer()
	if err != nil {
		log.Println("Ошибка парсинга конфига: ", err)
	}
	fmt.Println(cfg)
	// Инициализация кеша в памяти. Используется как основное хранилище данных.
	cache := storage.NewCache()

	// var event models.Event
	// err := json.Unmarshal(m.Data, &event)

	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	// Запуск сервера (Graceful shutdown внутри функции StartServer)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.StartServer(ctx, &cfg, cache); err != nil {
			log.Println("Ошибка при работе HTTP сервера:", err)
		} else {
			log.Println("Успешно закрыл HTTP сервер")
		}
	}()
	wg.Wait()
}
