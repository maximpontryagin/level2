package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Суть паттерна:
Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

Применимость:
Применяется там где есть логика сборки какого либо объяекта постепенно, заказ в магазине, выбор комплектующих для ПК и т.д.

Плюсы:
1. Пошаговое создание объекта, что увеличивает наглядность.
2. Возможность переиспользовать код для создания разновидностей одного объекта(каменный или деревянный дом и т.д.)
3. Шаблон Builder отделяет конструкцию от представления.

Минусы:
1. Увеличивается сложность кода из-за дополнительных структур
2. Билдер и создаваемый им продукт жестко связаны между собой => при внесеннии изменений в структуру и методы продукта
скорее всего придется соотвествующим образом изменять и его билдера.
3. Алгоритм создания сложного объекта не должен зависеть от того, из каких частей состоит объект и как они стыкуются между собой;

Примеры:
Заказ в магазине, выбор комплектующих для ПК и т.д.

Мой пример: Есть hr агенство где необходимо пошагово собрать характеристики кондидата на должность(имя, возраст, профессия, грейд).
*/

// Candidate - структура, представляющая кандидата
type Candidate struct {
	name       string
	age        int
	profession string
	grade      string
}

// candidateBuilder - структура строителя кандидата
type candidateBuilder struct {
	name       string
	age        int
	profession string
	grade      string
}

// CandidateBuilderI - интерфейс строителя кандидата
type CandidateBuilderI interface {
	Name(val string) CandidateBuilderI
	Age(val int) CandidateBuilderI
	Profession(val string) CandidateBuilderI
	Grade(val string) CandidateBuilderI
	Build() Candidate
}

// NewCandidateBuilder - функция для создания нового строителя кандидата
func NewCandidateBuilder() CandidateBuilderI {
	return &candidateBuilder{}
}

// Методы для задания характеристик кандидата
func (c *candidateBuilder) Name(val string) CandidateBuilderI {
	c.name = val
	return c
}

func (c *candidateBuilder) Age(val int) CandidateBuilderI {
	c.age = val
	return c
}

func (c *candidateBuilder) Profession(val string) CandidateBuilderI {
	c.profession = val
	return c
}

func (c *candidateBuilder) Grade(val string) CandidateBuilderI {
	c.grade = val
	return c
}

// Build - метод для создания кандидата
func (c *candidateBuilder) Build() Candidate {
	return Candidate{
		name:       c.name,
		age:        c.age,
		profession: c.profession,
		grade:      c.grade,
	}
}

func main() {
	// Объявление нового строителя кандидата
	candidateBuilder := NewCandidateBuilder()
	// Сбор кандидата по частям по паттерну строитель
	candidate := candidateBuilder.Name("Max").
		Age(23).
		Profession("Golang developer").
		Grade("Intern").
		Build()
	fmt.Println(candidate)
}
