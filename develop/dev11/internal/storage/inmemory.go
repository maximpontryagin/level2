package storage

import (
	"dev11/internal/models"
	"errors"
	"sync"
	"time"
)

// Cache - структура кэша.

type Cache struct {
	sync.RWMutex
	storage map[int]models.Event
}

// NewCache - функция инициализации хранилища
func NewCache() *Cache {
	return &Cache{storage: make(map[int]models.Event)}
}

func (cache *Cache) SaveEventInCache(event models.Event) error {
	cache.Lock()
	defer cache.Unlock()
	_, found := cache.storage[event.ID]
	if found {
		return errors.New("эвент с указанным ID уже существует")
	}
	cache.storage[event.ID] = event
	return nil
}

func (cache *Cache) UpdateEventInCache(event models.Event) error {
	cache.Lock()
	defer cache.Unlock()
	_, found := cache.storage[event.ID]
	if !found {
		return errors.New("эвент с указанным ID не найден")
	}
	cache.storage[event.ID] = event
	return nil
}

func (cache *Cache) GetEventCache(IdEvent int) (models.Event, bool) {
	cache.RLock()
	defer cache.RUnlock()
	value, found := cache.storage[IdEvent]
	return value, found
}

func (cache *Cache) DeleteEventCache(IdEvent int) error {
	cache.Lock()
	defer cache.Unlock()
	if _, found := cache.storage[IdEvent]; !found {
		return errors.New("эвент с указанным ID отсутствует в базе")
	}
	delete(cache.storage, IdEvent)
	return nil
}

func (cache *Cache) GetEventsForDay(day time.Time) []models.Event {
	cache.RLock()
	defer cache.RUnlock()
	var events []models.Event
	for _, event := range cache.storage {
		if event.Date.Year() == day.Year() && event.Date.YearDay() == day.YearDay() {
			events = append(events, event)
		}
	}
	return events
}

func (cache *Cache) GetEventsForWeek(weekStart time.Time) []models.Event {
	cache.RLock()
	defer cache.RUnlock()
	var events []models.Event
	weekEnd := weekStart.AddDate(0, 0, 7)
	for _, event := range cache.storage {
		if event.Date.After(weekStart) && event.Date.Before(weekEnd) {
			events = append(events, event)
		}
	}
	return events
}

func (cache *Cache) GetEventsForMonth(monthStart time.Time) []models.Event {
	cache.RLock()
	defer cache.RUnlock()
	var events []models.Event
	monthEnd := monthStart.AddDate(0, 1, 0)
	for _, event := range cache.storage {
		if event.Date.After(monthStart) && event.Date.Before(monthEnd) {
			events = append(events, event)
		}
	}
	return events
}
