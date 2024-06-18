package main

import (
	"reflect"
	"testing"
)

func TestReadLinesFromFile(t *testing.T) {
	lines, err := readLinesFromFile("test_input.txt")
	if err != nil {
		t.Fatalf("Ошибка в чтении тестового файла: %v", err)
	}

	expected := [][]string{
		{"apple", "10"},
		{"banana", "2"},
		{"cherry", "30"},
		{"banana", "3"},
		{"apple", "25"},
	}

	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("Ожидалось %v, но получилось %v", expected, lines)
	}
}

func TestSortByAlphabet(t *testing.T) {
	cfg := &SortConfing{sortColumnFlag: 0, sortReverseFlag: false}
	input := [][]string{
		{"banana", "2"},
		{"apple", "10"},
		{"cherry", "30"},
		{"banana", "3"},
		{"apple", "25"},
	}

	expected := [][]string{
		{"apple", "10"},
		{"apple", "25"},
		{"banana", "2"},
		{"banana", "3"},
		{"cherry", "30"},
	}

	result := sortByAlphabet(input, cfg)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, но получилось %v", expected, result)
	}
}

func TestSortByNumber(t *testing.T) {
	cfg := &SortConfing{sortColumnFlag: 1, sortByNumberFlag: true, sortReverseFlag: false}
	input := [][]string{
		{"banana", "2"},
		{"apple", "10"},
		{"cherry", "30"},
		{"banana", "3"},
		{"apple", "25"},
	}

	expected := [][]string{
		{"banana", "2"},
		{"banana", "3"},
		{"apple", "10"},
		{"apple", "25"},
		{"cherry", "30"},
	}

	result, err := sortByNumber(input, cfg)
	if err != nil {
		t.Fatalf("Ошибка в сортировке по числу: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, но получилось %v", expected, result)
	}
}

func TestSortInputReverse(t *testing.T) {
	cfg := &SortConfing{sortColumnFlag: 0, sortReverseFlag: true}
	input := [][]string{
		{"banana", "2"},
		{"apple", "10"},
		{"cherry", "30"},
		{"banana", "3"},
		{"apple", "25"},
	}

	expected := [][]string{
		{"cherry", "30"},
		{"banana", "2"},
		{"banana", "3"},
		{"apple", "10"},
		{"apple", "25"},
	}

	result := sortByAlphabet(input, cfg)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, но получилось %v", expected, result)
	}
}

func TestDeleteNonUniqueLines(t *testing.T) {
	input := [][]string{
		{"banana", "2"},
		{"apple", "10"},
		{"cherry", "30"},
		{"banana", "2"},
		{"apple", "25"},
	}

	expected := [][]string{
		{"banana", "2"},
		{"apple", "10"},
		{"cherry", "30"},
		{"apple", "25"},
	}

	result := deleteNonUniqueLines(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, но получилось %v", expected, result)
	}
}
