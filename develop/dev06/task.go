package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type cutConfig struct {
	fields    int
	delimiter string
	separated bool
}

func (cfg *cutConfig) parseConfig() {
	flag.IntVar(&cfg.fields, "f", 0, "выбрать поля (колонки)")
	flag.StringVar(&cfg.delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&cfg.separated, "s", false, "только строки с разделителем")

	flag.Parse()
}

func readLinesFromFile(fileName, delimiter string) ([][]string, error) {
	var result [][]string
	file, err := os.Open(fileName)
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		resLine := strings.Split(scanner.Text(), delimiter)
		result = append(result, resLine)
	}
	return result, scanner.Err()
}

func getResultByField(columns [][]string, numberField int, separated bool, delimiter string) ([][]string, error) {
	// getResultByField возвращает [][]string по указанному номеру колонки
	var result [][]string
	for _, line := range columns {
		// Если не нужно выводить колонки без разделителя, то проверяем входит ли delimiter(разделитель) в строку
		if separated && !strings.Contains(strings.Join(line, delimiter), delimiter) {
			continue
		}
		// Проверка что номер колонки из флага не выходит за пределы слайса
		if numberField < len(line) {
			result = append(result, []string{line[numberField]})
		} else {
			return nil, errors.New("номер колонки вышел за пределы")
		}
	}
	return result, nil
}

func main() {
	var cfg cutConfig
	cfg.parseConfig()
	args := flag.Args()
	if len(args) < 1 {
		log.Println("Пропущено название файла")
		os.Exit(1)
	}
	// Название файла из которого необходимо взять данные
	inputFile := args[len(args)-1]

	resultReadLine, err := readLinesFromFile(inputFile, cfg.delimiter)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// Получаем конечный результат
	result, err := getResultByField(resultReadLine, cfg.fields, cfg.separated, cfg.delimiter)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// Вывод результата
	for _, line := range result {
		fmt.Println(line)
	}
}
