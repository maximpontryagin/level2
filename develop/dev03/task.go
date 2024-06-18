package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Константа хранящая название файла для чтения
const inputFile = "input.txt"

// SortConfing хранит информацию о флагах
type SortConfing struct {
	sortColumnFlag    int
	sortByNumberFlag  bool
	sortReverseFlag   bool
	sortNotRepeatFlag bool
}

// Парсит флаге и заносит их в структуру SortCinfing
func (cfg *SortConfing) parseConfig() {
	flag.IntVar(&cfg.sortColumnFlag, "k", 0, "указание колонки для сортировки")
	flag.BoolVar(&cfg.sortByNumberFlag, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&cfg.sortReverseFlag, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&cfg.sortNotRepeatFlag, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()
}

func readLinesFromFile(fileName string) ([][]string, error) {
	var result [][]string
	file, err := os.Open(fileName)
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		resLine := strings.Split(scanner.Text(), " ")
		result = append(result, resLine)
	}
	return result, scanner.Err()
}

func sortByAlphabet(input [][]string, cfg *SortConfing) [][]string {
	// Функция сортировки по алфавиту

	// Проверка флага на требование к обратной сортировке. Сравниваются слова в указанном столбце (fg.sortColumnFlag)
	if cfg.sortReverseFlag {
		sort.Slice(input, func(i, j int) bool {
			return input[i][cfg.sortColumnFlag] > input[j][cfg.sortColumnFlag]
		})
	} else {
		sort.Slice(input, func(i, j int) bool {
			return input[i][cfg.sortColumnFlag] < input[j][cfg.sortColumnFlag]
		})
	}
	return input
}

func sortByNumber(input [][]string, cfg *SortConfing) ([][]string, error) {
	// Функция сортировки по числам

	// Проверка что все элементы в указанном в флаге столбце (fg.sortColumnFlag) - числа
	for i := 0; i < len(input); i++ {
		_, err := strconv.Atoi(input[i][cfg.sortColumnFlag])
		if err != nil {
			return input, errors.New("не все значения в колонке целые числа")
		}
	}
	// Проверка флага на требование к обратной сортировке. Сравниваются числа в указанном столбце (fg.sortColumnFlag)
	if !cfg.sortReverseFlag {
		sort.Slice(input, func(i, j int) bool {
			// Преобразование строки в целое число, проверять ошибку не нужно, т.к. проверка была выполнена выше
			elementOne, _ := strconv.Atoi(input[i][cfg.sortColumnFlag])
			elementTwo, _ := strconv.Atoi(input[j][cfg.sortColumnFlag])
			return elementOne < elementTwo
		})
	} else {
		sort.Slice(input, func(i, j int) bool {
			elementOne, _ := strconv.Atoi(input[i][cfg.sortColumnFlag])
			elementTwo, _ := strconv.Atoi(input[j][cfg.sortColumnFlag])
			return elementOne > elementTwo
		})
	}
	return input, nil
}

func sortInput(input [][]string, cfg *SortConfing) ([][]string, error) {
	// Если флаг на сортировку по числам false, то сортируем по алфавиту, иначе по числам
	if !cfg.sortByNumberFlag {
		result := sortByAlphabet(input, cfg)
		return result, nil
	}
	result, err := sortByNumber(input, cfg)
	if err != nil {
		return input, err
	}
	return result, nil
}

func deleteNonUniqueLines(input [][]string) [][]string {
	// Проверяет строки на уникальность
	seen := make(map[string]bool)
	var result [][]string
	for _, line := range input {
		strLine := strings.Join(line, " ")
		if !seen[strLine] {
			seen[strLine] = true
			result = append(result, line)
		}
	}
	return result
}

func main() {
	var cfg SortConfing
	cfg.parseConfig()
	// Считывание данных из файла inputFile в виде слайса слайсов строк var result [][]string
	resultReadLine, err := readLinesFromFile(inputFile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Вызов функции сортировки считанных данных
	res, err := sortInput(resultReadLine, &cfg)
	if err != nil {
		log.Println(err)
	}

	if cfg.sortNotRepeatFlag {
		res = deleteNonUniqueLines(res)
	}
	fmt.Println(res)
}
