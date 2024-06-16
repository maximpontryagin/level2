package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const escapeSymbolUnicode = 92 // номер обратного слеша в Unicode

func unpackString(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}
	if _, err := strconv.Atoi(string(str[0])); err == nil {
		return "", errors.New("incorrect string")
	}

	runes := []rune(str)
	var result strings.Builder

	for i := 0; i < len(runes); i++ {
		if runes[i] == escapeSymbolUnicode {
			i++ // Переходим на символ который нужно продублировать
		}
		toWrite := string(runes[i])

		if i+1 != len(runes) {
			if unicode.IsDigit(runes[i+1]) {
				number, err := strconv.Atoi(string(runes[i+1]))
				if err != nil {
					return "", errors.New("ошибка при конвертации строки в число")
				}
				for j := 0; j < number; j++ {
					result.WriteString(toWrite) // Запись символа указанного колличества раз
				}
				i++ // перешагиваем цифру
			} else {
				result.WriteString(toWrite) // Запись символа после которого нету цифры
			}
		} else {
			result.WriteString(toWrite) // Запись последнего символа
		}
	}
	return result.String(), nil
}

func main() {
	var str string
	fmt.Scan(&str)
	unpacked, err := unpackString(str)
	if err != nil {
		log.Println("Error:", err)
	}
	fmt.Println(unpacked)
}
