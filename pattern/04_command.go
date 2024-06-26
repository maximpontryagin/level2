package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Суть паттерна:
Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты,
позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь,
логировать их, а также поддерживать отмену операций.

Применимость паттерна:
1. Когда нужно параметризовать объекты выполняемым действием. Команда превращает операции в объекты.
А объекты можно передавать, хранить и взаимозаменять внутри других объектов.
2. Когда нужно ставить операции в очередь или выполнять их по расписанию.
3. Когда нужна операция отмены.

Плюсы:
1. Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
2. Позволяет реализовать простую отмену и повтор операций.
3. Т.к. команда это отдельный объект, в который мы добавляем нужную нам логику, не изменяя оригинальной структуры,
мы следуем принципу открытости - закрытости.
4. Позволяет собирать сложные команды из простых.

Минусы:
1. Усложнение кода из-за вода множества дополнительных структур

Примеры:
1. Отмена/повтор операции: В текстовых редакторах или графических редакторах для реализации функций "Отмена" и "Повтор".
2. Очереди задач: В системах, где задачи нужно выполнять по расписанию или в порядке очереди. Например,
это может быть система управления печатью в офисе, где множество пользователей отправляют документы на печать.
Система должна обрабатывать эти задачи по мере их поступления, а также поддерживать возможность отмены и повторного выполнения задач.

Мой пример:
Есть методы включения и выключения света в умном доме и так вышло, что нам необходимо добавить возможность
отмены и повторного вызова команды.
Основные компоненты:
Интерфейс Command: Определяет методы Execute() и Undo().
Конкретные команды: Реализуют интерфейс Command и содержат ссылки на приемник (Receiver).
Приемник (Receiver): Класс, который выполняет действия команд.
Отправитель (Invoker): Класс, который хранит и вызывает команды.
*/

// Command интерфейс определяет методы Execute и Undo
type Command interface {
	Execute()
	Undo()
}

// LightOnCommand реализует Command для включения света
type LightOnCommand struct {
	receiver *Light
}

// NewLightOnCommand создает новый LightOnCommand
func NewLightOnCommand(receiver *Light) *LightOnCommand {
	return &LightOnCommand{receiver: receiver}
}

// Execute выполняет команду включения света
func (c *LightOnCommand) Execute() {
	c.receiver.On()
}

// Undo отменяет команду включения света
func (c *LightOnCommand) Undo() {
	c.receiver.Off()
}

// LightOffCommand реализует Command для выключения света
type LightOffCommand struct {
	receiver *Light
}

// NewLightOffCommand создает новый LightOffCommand
func NewLightOffCommand(receiver *Light) *LightOffCommand {
	return &LightOffCommand{receiver: receiver}
}

// Execute выполняет команду выключения света
func (c *LightOffCommand) Execute() {
	c.receiver.Off()
}

// Undo отменяет команду выключения света
func (c *LightOffCommand) Undo() {
	c.receiver.On()
}

// Light представляет приемник для управления светом
type Light struct{}

func (l *Light) On() {
	fmt.Println("Свет включен")
}

func (l *Light) Off() {
	fmt.Println("Свет выключен")
}

// Invoker хранит команды и выполняет их
type Invoker struct {
	command Command
}

// SetCommand устанавливает команду для Invoker
func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

// ExecuteCommand выполняет установленную команду
func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

// UndoCommand отменяет установленную команду
func (i *Invoker) UndoCommand() {
	i.command.Undo()
}

func main() {
	// Создаем приемник
	light := &Light{}

	// Создаем команды
	lightOnCommand := NewLightOnCommand(light)
	lightOffCommand := NewLightOffCommand(light)

	// Создаем invoker
	invoker := &Invoker{}

	// Выполнение и отмена команды включения света
	invoker.SetCommand(lightOnCommand)
	invoker.ExecuteCommand()
	invoker.UndoCommand()

	// Выполнение и отмена команды выключения света
	invoker.SetCommand(lightOffCommand)
	invoker.ExecuteCommand()
	invoker.UndoCommand()
}
