package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern


Суть паттерна:
фабричный метод это порождающий паттерн проектирования, который определяет интерфейс для создания объектов определённого типа.
Паттерн предлагает создавать объекты не напрямую, а через вызов специального фабричного метода. Все объекты должны иметь общий интерфейс,
который отвечает за их создание.

Применимость:
1. Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код.
Фабричный метод отделяет код производства продуктов от остального кода, который эти продукты использует.
2. Когда нужно дать возможность пользователям расширять части вашего фреймворка или библиотеки.
3. Когда нужно экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых.

Плюсы:
1. Упрощает добавление новых объектов в программу.
2. Реализует принцип открытости/закрытости.
3. Разделяет создание объектов от их использования

Минусы:
1. Может привести к созданию больших параллельных иерархий объектов.
2. Появляется божественный конструктор. Инициализация объетов всегда идет через него.

Примеры:
1. Многие ORM-библиотеки используют фабричный метод для создания подключений к различным базам данных.
2. При тестировании приложений часто требуется подменять реальные реализации зависимостей на моки или заглушки.
Фабричный метод позволяет гибко создавать эти моки в зависимости от типа теста.

Мой пример:
Есть интернет магазин который раньше продавал только серверы, но теперь решает расширить
каталог ПК и ноутбуками, для реализации этого используется фабричный метод
*/

const (
	ServerType            = "Server"
	PerosonalComputerType = "PersonalComputer"
	NotebookType          = "Notebook"
)

// Computer - интерфейс, который будут реализовывать различные типы компьютеров
type Computer interface {
	GetType() string
	PrintDetails()
}

// Server - структура для типа "Сервер"
type Server struct {
	Type   string
	Core   int
	Memory int
}

func (s Server) GetType() string {
	return s.Type
}

func (s Server) PrintDetails() {
	fmt.Printf("%s: Core: %d, Mem: %d\n", s.Type, s.Core, s.Memory)
}

// NewServer - функция для создания нового сервера
func NewServer() Computer {
	return Server{
		Type:   ServerType,
		Core:   16,
		Memory: 1024,
	}
}

// PersonalComputer - структура для типа "Персональный компьютер"
type PersonalComputer struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func (p PersonalComputer) GetType() string {
	return p.Type
}

func (p PersonalComputer) PrintDetails() {
	fmt.Printf("%s: Core: %d, Mem: %d, Monitor: %v\n", p.Type, p.Core, p.Memory, p.Monitor)
}

// NewPersonalComputer - функция для создания нового персонального компьютера
func NewPersonalComputer() Computer {
	return PersonalComputer{
		Type:    PerosonalComputerType,
		Core:    4,
		Memory:  512,
		Monitor: true,
	}
}

// Notebook - структура для типа "Ноутбук"
type Notebook struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func (n Notebook) GetType() string {
	return n.Type
}

func (n Notebook) PrintDetails() {
	fmt.Printf("%s: Core: %d, Mem: %d, Monitor: %v\n", n.Type, n.Core, n.Memory, n.Monitor)
}

// NewNotebook - функция для создания нового ноутбука
func NewNotebook() Computer {
	return Notebook{
		Type:    NotebookType,
		Core:    2,
		Memory:  128,
		Monitor: false,
	}
}

// New - фабричная функция для создания различных типов компьютеров
func New(typeName string) Computer {
	switch typeName {
	case ServerType:
		return NewServer()
	case PerosonalComputerType:
		return NewPersonalComputer()
	case NotebookType:
		return NewNotebook()
	default:
		fmt.Printf("%s: Несуществующий тип объекта!\n", typeName)
		return nil
	}
}

// Пример клиентского кода
func main() {
	// Создание сервера
	server := New(ServerType)
	if server != nil {
		server.PrintDetails()
	}

	// Создание персонального компьютера
	personalComputer := New(PerosonalComputerType)
	if personalComputer != nil {
		personalComputer.PrintDetails()
	}

	// Создание ноутбука
	notebook := New(NotebookType)
	if notebook != nil {
		notebook.PrintDetails()
	}

	// Попытка создания объекта с несуществующим типом
	unknown := New("UnknownType")
	if unknown != nil {
		unknown.PrintDetails()
	}
}
