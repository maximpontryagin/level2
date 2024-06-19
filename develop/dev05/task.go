package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type grepConfig struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

func (cfg *grepConfig) parseConfig() {
	flag.IntVar(&cfg.after, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&cfg.before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&cfg.context, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&cfg.count, "c", false, "количество строк")
	flag.BoolVar(&cfg.ignoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&cfg.invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&cfg.fixed, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&cfg.lineNum, "n", false, "печатать номер строки")
	flag.Parse()
}

func readLinesFromFile(fileName string) ([]string, error) {
	var result []string
	file, err := os.Open(fileName)
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}

func createRegexp(cfg *grepConfig, pattern string) *regexp.Regexp {
	var reg *regexp.Regexp
	if cfg.fixed {
		// Если флаг fixed, используем точное совпадение строки
		// QuoteMeta возвращает строку, которая экранирует все метасимволы регулярных выражений внутри
		// текста аргумента; возвращаемая строка представляет собой регулярное выражение, соответствующее буквальному тексту
		reg = regexp.MustCompile(regexp.QuoteMeta(pattern))
	} else {
		if cfg.ignoreCase {
			// Если флаг ignore-case, добавляем (?i) для игнорирования регистра
			reg = regexp.MustCompile("(?i)" + pattern)
		} else {
			// Компиляция регулярного выражения без игнорирования регистра для сопоставления с текстом
			reg = regexp.MustCompile(pattern)
		}
	}
	return reg
}

func searchMatchLines(resultReadLines []string, cfg *grepConfig, reg *regexp.Regexp) []int {
	// Возвращает массив для хранения индексов строк, соответствующих шаблону
	matchedLines := make([]int, 0)
	for idx, line := range resultReadLines {
		// MatchString сообщает, содержит ли строка какое-либо совпадение с регулярным выражением reg
		matched := reg.MatchString(line)
		// Если инвертирование включено и строка не соответствует шаблону,
		// или если инвертирование выключено и строка соответствует шаблону
		if (cfg.invert && !matched) || (!cfg.invert && matched) {
			matchedLines = append(matchedLines, idx)
		}
	}
	return matchedLines
}

func countLines(matchedLines []int) int {
	return len(matchedLines)
}

func findIdxLines(cfg *grepConfig, matchedLines []int, resultReadLines []string) map[int]bool {
	// Собираем уникальные индексы строк для вывода, учитывая флаги after, before и context
	linesToPrint := make(map[int]bool)
	for _, idx := range matchedLines {
		// Определение начало и конец вывода. если флаги у before и after не установлены, то они = 0
		start := max(0, idx-cfg.before)
		end := min(len(resultReadLines)-1, idx+cfg.after)

		// Если установлен флаг context, переопределяем start и end
		if cfg.context > 0 {
			start = max(0, idx-cfg.context)
			end = min(len(resultReadLines)-1, idx+cfg.context)
		}
		// Добавляем строки в диапазоне [start, end] в результат
		for j := start; j <= end; j++ {
			linesToPrint[j] = true
		}
	}
	return linesToPrint
}

func main() {
	var cfg grepConfig
	cfg.parseConfig()

	// Args - элементы которые идут после вызова программы. Пример: go run task.go Args
	args := flag.Args()
	// Проверка что ввели паттерн и название файла
	if len(args) < 2 {
		log.Println("Пропущен паттерн и/или название файла")
		os.Exit(1)
	}
	// Паттерн идет первым элементов в args
	pattern := args[0]
	// Название идет вторым элементво в args
	inputFile := args[1]
	// Считываем указанный в args файл
	resultReadLines, err := readLinesFromFile(inputFile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Компиляция регулярного выражения для сопоставления с текстом
	reg := createRegexp(&cfg, pattern)

	// Создание массива для хранения индексов строк, соответствующих шаблону []int
	matchedLines := searchMatchLines(resultReadLines, &cfg, reg)

	// Если флаг count установлен, выводим количество строк и завершаем выполнение
	if cfg.count {
		result := countLines(matchedLines)
		log.Println("количество совпадающих строк:", result)
		return
	}

	// Собираем уникальные индексы строк для вывода, учитывая флаги after, before и context
	linesToPrint := findIdxLines(&cfg, matchedLines, resultReadLines)

	// Печать строк
	for i := range linesToPrint {
		// Если флаг lineNum установлен, выводим номер строки
		if cfg.lineNum {
			fmt.Printf("%d:", i+1)
		}
		// Выводим строку
		fmt.Println(resultReadLines[i])
	}
}
