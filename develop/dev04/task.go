package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Сортирует буквы в словах в алфавитном порядке и возвращает отсортированную строку
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func searchAnagram(listWords []string) map[string][]string {
	anagramsSets := make(map[string][]string) // anagramsSets хранит множества анаграм
	wordMap := make(map[string]string)        // для отслеживания отсортированных версий слов

	for _, word := range listWords {
		lowerWord := strings.ToLower(word)  // Приведение слова к нижнему регистру
		sortedWord := sortString(lowerWord) // Сортировка символов в слове
		// Проверяем есть ли отсортированное слово в wordMap, если существует,
		// то добавляем слово в существующее множество анаграмм, если нет, создаем новое множество.
		if originalWord, exists := wordMap[sortedWord]; exists {
			anagramsSets[originalWord] = append(anagramsSets[originalWord], lowerWord)
		} else {
			wordMap[sortedWord] = lowerWord
			anagramsSets[lowerWord] = []string{lowerWord}
		}
	}

	// Удаление множеств состоящих из 1 элемента и сортировка массивов
	for key, listAnagrams := range anagramsSets {
		if len(listAnagrams) < 2 {
			delete(anagramsSets, key)
		} else {
			sort.Strings(listAnagrams)
		}
	}
	return anagramsSets
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кот", "ток", "кто"}
	anagramSets := searchAnagram(words)
	for key, group := range anagramSets {
		fmt.Printf("%s: %v\n", key, group)
	}
}
