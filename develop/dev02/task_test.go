package main

import "testing"

func TestUnpackString(t *testing.T) {
	list_tests := []struct {
		input    string // Введенная строка
		expected string // Ожидаемный результат
		errTest  bool   // true если ошибка ожидается, false если не ожидается
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"qwe\\4\\5", "qwe45", false},
		{"qwe\\45", "qwe44444", false},
		{"qwe\\\\5", "qwe\\\\\\\\\\", false},
	}

	for _, test := range list_tests {
		test_result, err := unpackString(test.input)
		if err != nil && test.errTest == false {
			t.Error("Получили ошибку там где недолжны были. Ошибка: ", err, "Ожидал: ", test.expected, "Поличил: ", test_result)
		}
		if err == nil && test.errTest == true {
			t.Error("Не получили ошибку там где должны были ее получить.", "Ожидал: ", test.expected, "Поличил: ", test_result)
		}

		if test_result != test.expected {
			t.Error("Ожидаемый и полученный результат не совпали. Ожидал: ", test.expected, "Поличил: ", test_result)
		}
	}
}
