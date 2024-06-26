package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять
поведение в зависимости от своего состояния.

Применимость паттерна:
1. Когда есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния,
причём типов состояний много, и их код часто меняется.
2. Когда операции имеют разные реализации, зависящие от состояния(большие условные конструкции).

Плюсы:
1. Избавляет от множества больших условных операторов.
2. Концентрирует в одном месте код, связанный с определённым состоянием.

Минусы:
1. Может неоправданно усложнить код, если состояний мало и они редко меняются.

Примеры:
1. В играх состояние игрока (бег, прыжок, атака) может быть реализовано с помощью паттерна Состояние.
2. Пост в социальной сети может быть реализован с помощью паттерна Состояние
(в теневом бане, в бане, продвигается рекламой, удален, стандартное состояние и т.д.)

Мой пример:
Допустим, у нас есть кофемашина, которая может находиться в трех состояниях: ожидание, приготовление кофе и выдача кофе.
реализуем эти состояния с использованием паттерна состояние.
*/

// State интерфейс, определяющий методы для состояний
type State interface {
	InsertCoin()
	PressButton()
	DispenseCoffee()
}

// CoffeeMachine контекст, который хранит текущее состояние
type CoffeeMachine struct {
	state State
}

func (c *CoffeeMachine) SetState(state State) {
	c.state = state
}

func (c *CoffeeMachine) InsertCoin() {
	c.state.InsertCoin()
}

func (c *CoffeeMachine) PressButton() {
	c.state.PressButton()
}

func (c *CoffeeMachine) DispenseCoffee() {
	c.state.DispenseCoffee()
}

// WaitingState состояние ожидания
type WaitingState struct {
	machine *CoffeeMachine
}

func (s *WaitingState) InsertCoin() {
	fmt.Println("Монета вставлена. Пожалуйста, нажмите кнопку для приготовления кофе.")
	s.machine.SetState(&BrewingState{machine: s.machine})
}

func (s *WaitingState) PressButton() {
	fmt.Println("Пожалуйста, вставьте монету сначала.")
}

func (s *WaitingState) DispenseCoffee() {
	fmt.Println("Пожалуйста, вставьте монету сначала.")
}

// BrewingState состояние приготовления кофе
type BrewingState struct {
	machine *CoffeeMachine
}

func (s *BrewingState) InsertCoin() {
	fmt.Println("Монета уже вставлена. Пожалуйста, нажмите кнопку для приготовления кофе.")
}

func (s *BrewingState) PressButton() {
	fmt.Println("Приготовление кофе... Пожалуйста, подождите.")
	s.machine.SetState(&DispensingState{machine: s.machine})
}

func (s *BrewingState) DispenseCoffee() {
	fmt.Println("Пожалуйста, подождите. Кофе готовится.")
}

// DispensingState состояние выдачи кофе
type DispensingState struct {
	machine *CoffeeMachine
}

func (s *DispensingState) InsertCoin() {
	fmt.Println("Пожалуйста, подождите. Кофе выдаётся.")
}

func (s *DispensingState) PressButton() {
	fmt.Println("Пожалуйста, подождите. Кофе выдаётся.")
}

func (s *DispensingState) DispenseCoffee() {
	fmt.Println("Ваш кофе готов! Приятного дня.")
	s.machine.SetState(&WaitingState{machine: s.machine})
}

func main() {
	machine := &CoffeeMachine{}
	waitingState := &WaitingState{machine: machine}
	machine.SetState(waitingState)

	// Пример взаимодействия с кофемашиной
	machine.PressButton()    // Output: Пожалуйста, вставьте монету сначала.
	machine.DispenseCoffee() // Output: Пожалуйста, вставьте монету сначала.
	machine.InsertCoin()     // Output: Монета вставлена. Пожалуйста, нажмите кнопку для приготовления кофе.
	machine.PressButton()    // Output: Приготовление кофе... Пожалуйста, подождите.
	machine.DispenseCoffee() // Output: Пожалуйста, подождите. Кофе готовится.
	machine.DispenseCoffee() // Output: Ваш кофе готов! Приятного дня.
}
