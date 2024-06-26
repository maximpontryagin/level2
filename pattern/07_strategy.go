package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Стратегия — это поведенческий паттерн проектирования

Применимость:
1. Когда нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
2. Когда есть множество похожих классов, отличающихся только некоторым поведением.
3. Когда не хочется обнажать детали реализации алгоритмов для других классов.
4. Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора.
Каждая ветка такого оператора представляет собой вариацию алгоритма.

Плюсы:
1. Горячая замена алгоритмов на лету.
2. Изолирует код и данные алгоритмов от остальных классов.
3. Уход от наследования к делегированию.
4. Реализует принцип открытости/закрытости.

Минусы:
1. Усложняет программу за счёт дополнительных классов.
2. Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

Примеры:
1. Поиск опитмального маршрута разными стратегиями пеший/авто/общественный транспорт.

Мой пример:
Есть сервис, который поддерживает оплату через разные платежные системы, такие как PayPal,
Mir и Bitcoin. Каждая платежная система будет представлена как отдельная стратегия.
*/

// PaymentStrategy определяет интерфейс стратегии оплаты
type PaymentStrategy interface {
	Pay(amount float64) string
}

// PayPalStrategy реализует стратегию оплаты через PayPal
type PayPalStrategy struct {
	email string
}

func (p *PayPalStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal account %s", amount, p.email)
}

// MirStrategy реализует стратегию оплаты через Mir
type MirStrategy struct {
	cardNumber string
}

func (s *MirStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Mir card %s", amount, s.cardNumber)
}

// BitcoinStrategy реализует стратегию оплаты через Bitcoin
type BitcoinStrategy struct {
	walletAddress string
}

func (b *BitcoinStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Bitcoin wallet %s", amount, b.walletAddress)
}

// Order представляет собой заказ, который использует стратегию оплаты
type Order struct {
	amount   float64
	strategy PaymentStrategy
}

// SetStrategy задает стратегию оплаты для заказа
func (o *Order) SetStrategy(strategy PaymentStrategy) {
	o.strategy = strategy
}

// ProcessPayment выполняет оплату с использованием текущей стратегии
func (o *Order) ProcessPayment() string {
	return o.strategy.Pay(o.amount)
}

func main() {
	order := &Order{amount: 100.0}

	// Оплата через PayPal
	order.SetStrategy(&PayPalStrategy{email: "user@example.com"})
	fmt.Println(order.ProcessPayment())

	// Оплата через Mir
	order.SetStrategy(&MirStrategy{cardNumber: "1234-5678-9101-1121"})
	fmt.Println(order.ProcessPayment())

	// Оплата через Bitcoin
	order.SetStrategy(&BitcoinStrategy{walletAddress: "1BitcoinAddress12345"})
	fmt.Println(order.ProcessPayment())
}
