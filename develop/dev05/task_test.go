package main

import (
	"os"
	"reflect"
	"regexp"
	"testing"
)

func TestReadLinesFromFile(t *testing.T) {
	content := `Это тестовый файл.
Он содержит несколько строк.
Некоторые из них совпадают с шаблоном.
Шаблон - это слово, которое мы ищем.
Это конец файла.`

	fileName, err := createTestFile(content)
	if err != nil {
		t.Fatalf("Не удалось создать временный файл: %v", err)
	}
	defer os.Remove(fileName)

	lines, err := readLinesFromFile(fileName)
	if err != nil {
		t.Fatalf("Не удалось прочитать строки из файла: %v", err)
	}

	expectedLines := []string{
		"Это тестовый файл.",
		"Он содержит несколько строк.",
		"Некоторые из них совпадают с шаблоном.",
		"Шаблон - это слово, которое мы ищем.",
		"Это конец файла.",
	}

	if !reflect.DeepEqual(lines, expectedLines) {
		t.Errorf("Ожидалось: %v, получено: %v", expectedLines, lines)
	}
}

func TestCreateRegexp(t *testing.T) {
	cfg := &grepConfig{ignoreCase: true}
	pattern := "шаблон"
	reg := createRegexp(cfg, pattern)
	expectedPattern := "(?i)шаблон"

	if reg.String() != expectedPattern {
		t.Errorf("Ожидалось: %v, получено: %v", expectedPattern, reg.String())
	}

	cfg.fixed = true
	cfg.ignoreCase = false
	reg = createRegexp(cfg, pattern)
	expectedPattern = regexp.QuoteMeta(pattern)

	if reg.String() != expectedPattern {
		t.Errorf("Ожидалось: %v, получено: %v", expectedPattern, reg.String())
	}
}

func TestSearchMatchLines(t *testing.T) {
	lines := []string{
		"Это тестовый файл.",
		"Он содержит несколько строк.",
		"Некоторые из них совпадают с шаблоном.",
		"Шаблон - это слово, которое мы ищем.",
		"Это конец файла.",
	}

	cfg := &grepConfig{}
	reg := regexp.MustCompile("Шаблон")

	matchedLines := searchMatchLines(lines, cfg, reg)
	expectedMatchedLines := []int{3}

	if !reflect.DeepEqual(matchedLines, expectedMatchedLines) {
		t.Errorf("Ожидалось: %v, получено: %v", expectedMatchedLines, matchedLines)
	}
}

func TestCountLines(t *testing.T) {
	matchedLines := []int{1, 2, 3}
	count := countLines(matchedLines)
	expectedCount := 3

	if count != expectedCount {
		t.Errorf("Ожидалось: %v, получено: %v", expectedCount, count)
	}
}

func TestFindIdxLines(t *testing.T) {
	lines := []string{
		"Это тестовый файл.",
		"Он содержит несколько строк.",
		"Некоторые из них совпадают с шаблоном.",
		"Шаблон - это слово, которое мы ищем.",
		"Это конец файла.",
	}

	cfg := &grepConfig{before: 1, after: 1}
	matchedLines := []int{3}
	linesToPrint := findIdxLines(cfg, matchedLines, lines)
	expectedLinesToPrint := map[int]bool{2: true, 3: true, 4: true}

	if !reflect.DeepEqual(linesToPrint, expectedLinesToPrint) {
		t.Errorf("Ожидалось: %v, получено: %v", expectedLinesToPrint, linesToPrint)
	}
}

// Вспомогательная функция для создания временного файла
func createTestFile(content string) (string, error) {
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}
