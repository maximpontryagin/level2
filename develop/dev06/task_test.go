package main

import (
	"os"
	"reflect"
	"testing"
)

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

func TestReadLinesFromFile(t *testing.T) {
	tests := []struct {
		name             string
		content          string
		delimiter        string
		expectedReadLine [][]string
		separated        bool
		numberField      int
		expectedResult   [][]string
	}{
		{
			name:      "базовый случай",
			content:   "a\tb\tc\nd\te\tf\n",
			delimiter: "\t",
			expectedReadLine: [][]string{
				{"a", "b", "c"},
				{"d", "e", "f"},
			},
			separated:   false,
			numberField: 0,
			expectedResult: [][]string{
				{"a"},
				{"d"},
			},
		},
		{
			name:      "другой разделитель",
			content:   "a,b,c\nd,e,f\n",
			delimiter: ",",
			expectedReadLine: [][]string{
				{"a", "b", "c"},
				{"d", "e", "f"},
			},
			separated:   false,
			numberField: 0,
			expectedResult: [][]string{
				{"a"},
				{"d"},
			},
		},
		{
			name:      "Проверка с флагом separated",
			content:   "a,b,c\nd,e,f\ndef",
			delimiter: ",",
			expectedReadLine: [][]string{
				{"a", "b", "c"},
				{"d", "e", "f"},
				{"def"},
			},
			separated:   true,
			numberField: 0,
			expectedResult: [][]string{
				{"a"},
				{"d"},
			},
		},
		{
			name:      "Проверка без флага separated",
			content:   "a,b,c\nd,e,f\ndef",
			delimiter: ",",
			expectedReadLine: [][]string{
				{"a", "b", "c"},
				{"d", "e", "f"},
				{"def"},
			},
			separated:   false,
			numberField: 0,
			expectedResult: [][]string{
				{"a"},
				{"d"},
				{"def"},
			},
		},
		{
			name:      "Проверка второго столбца",
			content:   "a,b,c\nd,e,f\n",
			delimiter: ",",
			expectedReadLine: [][]string{
				{"a", "b", "c"},
				{"d", "e", "f"},
			},
			separated:   true,
			numberField: 1,
			expectedResult: [][]string{
				{"b"},
				{"e"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fileName, _ := createTestFile(test.content)
			resultReadLine, _ := readLinesFromFile(fileName, test.delimiter)
			if !reflect.DeepEqual(test.expectedReadLine, resultReadLine) {
				t.Errorf("1. Ожидалось: %v, получено: %v", test.expectedReadLine, resultReadLine)
			}
			result, _ := getResultByField(resultReadLine, test.numberField, test.separated, test.delimiter)
			if !reflect.DeepEqual(test.expectedResult, result) {
				t.Errorf("2. Ожидалось: %v, получено: %v", test.expectedResult, result)
			}
		})
	}
}
