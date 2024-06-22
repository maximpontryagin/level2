package server

import (
	"context"
	"dev11/internal/config"
	"dev11/internal/models"
	"dev11/internal/storage"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type HandlersWithCache struct {
	cache *storage.Cache
}

func NewHandlersWithCache(cache *storage.Cache) *HandlersWithCache {
	return &HandlersWithCache{cache: cache}
}

func (handler HandlersWithCache) HandlerCreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	var event models.Event
	// Читаем тело запроса. ReadAll выполняет чтение из r до тех пор, пока не возникнет ошибка или EOF
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка в чтении запроса", http.StatusBadRequest)
		return
	} else {
		// Если все нормально - пишем по указателю в структуру event
		err = json.Unmarshal(body, &event)
		if err != nil {
			http.Error(w, "Ошибка в чтении запроса", http.StatusBadRequest)
			return
		}
	}
	// Записы полученного ивента в память
	err = handler.cache.SaveEventInCache(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	err = json.NewEncoder(w).Encode(models.SuccessfulRequest{Result: event})
	if err != nil {
		http.Error(w, "Ошибка в отправке Json", http.StatusInternalServerError)
		return
	}
}

func (handler HandlersWithCache) HandlerUpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	var event models.Event
	// Читаем тело запроса. ReadAll выполняет чтение из r до тех пор, пока не возникнет ошибка или EOF
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка в чтении запроса", http.StatusBadRequest)
		return
	} else {
		// Если все нормально - пишем по указателю в структуру event
		err = json.Unmarshal(body, &event)
		if err != nil {
			http.Error(w, "Ошибка в чтении запроса", http.StatusBadRequest)
			return
		}
	}
	// Обновление полученного эвента
	err = handler.cache.UpdateEventInCache(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	err = json.NewEncoder(w).Encode(models.SuccessfulRequest{Result: event})
	if err != nil {
		http.Error(w, "Ошибка в отправке Json", http.StatusInternalServerError)
		return
	}
}

func (handler HandlersWithCache) HandlerDeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	var event models.Event
	// Читаем тело запроса. ReadAll выполняет чтение из r до тех пор, пока не возникнет ошибка или EOF
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка в чтении запроса", http.StatusBadRequest)
		return
	} else {
		// Если все нормально - пишем по указателю в структуру event
		err = json.Unmarshal(body, &event)
		if err != nil {
			http.Error(w, "Ошибка в чтении запроса", http.StatusBadRequest)
			return
		}
	}
	// Обновление полученного эвента
	err = handler.cache.DeleteEventCache(event.ID)
	if err != nil {
		// err = json.NewEncoder(w).Encode(models.EventExists{Error: err.Error()})
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	err = json.NewEncoder(w).Encode("Эвент с ID " + strconv.Itoa(event.ID) + " успешно удален")
	if err != nil {
		http.Error(w, "Ошибка в отправке Json", http.StatusInternalServerError)
		return
	}
}

func (handler HandlersWithCache) HandlerEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	today := time.Now()
	events := handler.cache.GetEventsForDay(today)
	err := json.NewEncoder(w).Encode(models.SuccessfulRequestGet{Results: events})
	if err != nil {
		http.Error(w, "Ошибка в отправке Json", http.StatusInternalServerError)
		return
	}
}

func (handler HandlersWithCache) HandlerEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	startOfWeek := time.Now().AddDate(0, 0, -int(time.Now().Weekday()))
	events := handler.cache.GetEventsForWeek(startOfWeek)
	err := json.NewEncoder(w).Encode(models.SuccessfulRequestGet{Results: events})
	if err != nil {
		http.Error(w, "Ошибка в отправке Json", http.StatusInternalServerError)
		return
	}
}

func (handler HandlersWithCache) HandlerEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	startOfMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Local)
	events := handler.cache.GetEventsForMonth(startOfMonth)
	err := json.NewEncoder(w).Encode(models.SuccessfulRequestGet{Results: events})
	if err != nil {
		http.Error(w, "Ошибка в отправке Json", http.StatusInternalServerError)
		return
	}
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		inputTime := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("Пришел %s запрос | URL %s | Дата и время %s | Обработался за %s", req.Method, req.RequestURI, inputTime, time.Since(inputTime))
	})
}

func StartServer(ctx context.Context, cfg *config.ConfigServer, cache *storage.Cache) error {
	httpServer := &http.Server{
		Addr: cfg.Host + ":" + cfg.Port,
	}
	handlerWithCache := NewHandlersWithCache(cache)
	http.HandleFunc("/create_event/", loggingMiddleware(handlerWithCache.HandlerCreateEvent))
	http.HandleFunc("/update_event/", loggingMiddleware(handlerWithCache.HandlerUpdateEvent))
	http.HandleFunc("/delete_event/", loggingMiddleware(handlerWithCache.HandlerDeleteEvent))
	http.HandleFunc("/events_for_day/", loggingMiddleware(handlerWithCache.HandlerEventsForDay))
	http.HandleFunc("/events_for_week/", loggingMiddleware(handlerWithCache.HandlerEventsForWeek))
	http.HandleFunc("/events_for_month/", loggingMiddleware(handlerWithCache.HandlerEventsForMonth))

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка при запуске сервера: %v", err)
		}
	}()
	log.Println("Сервер запустился и слушает порт", httpServer.Addr)
	// Ждём сигнал о завершении работы сервиса
	<-ctx.Done()
	// Завершаем работу http сервера
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second)
	defer shutdownCancel()
	return httpServer.Shutdown(shutdownCtx)

}
