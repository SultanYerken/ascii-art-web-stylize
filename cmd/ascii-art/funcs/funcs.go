package funcs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	NumberOfStringInFile        = 855
	CountSymbolInFileStandard   = 6623
	CountSymbolInFileShadow     = 7463
	CountSymbolInFileThinkertoy = 4703
)

var (
	ErrorBad            = errors.New("error: BadRequset")
	ErrorInternal       = errors.New("error: InternalServer")
	ErrorContentChanged = errors.New("error: readable file.txt changed")
)

func ReadArgs(s string) ([]string, error) {
	s = strings.Replace(s, "\r\n", "\n", -1)

	for _, w := range s {
		if (w < ' ' || w > '~') && w != 10 {
			return nil, ErrorBad
		}
	}

	// newline := regexp.MustCompile(`(\\n)`)
	// res := newline.ReplaceAllString(s, "\n")

	// Splits a string into an array of strings, where each string is delimited by '\n
	// Разделяет строку на массив строк, где каждая строка разделена '\n'
	resarray := strings.Split(s, "\n")

	return resarray, nil
}

func Readfile(s string) (string, error) {
	if s == "" {
		s = "standard"
	}

	font, err := errchangeFont(s)
	if err != nil {
		log.Println(err)
		return "", ErrorBad
	}

	content, err := ioutil.ReadFile(font)
	if err != nil {
		log.Println(err)
		return "", ErrorInternal
	}
	if len(content) != CountSymbolInFileStandard && len(content) != CountSymbolInFileShadow && len(content) != CountSymbolInFileThinkertoy {
		fmt.Println("len(content)", len(content))
		return "", ErrorContentChanged
	}

	return string(content), nil
}

// splits file.txt into a double array of strings by Characters and Lines within a Character
// разбивает file.txt на двойной массив строк по Знакам и Строкам внутри Знака
func Arrayart(s string) ([][]string, error) {
	res := [][]string{}
	arrart := []string{}
	str := ""
	count := 0
	countstrings := 0

	for _, w := range s {
		str += string(w)
		if w == '\n' {
			count++
			countstrings++
			// adds the string "str" ​​before '\n' to the array of strings "arrart" and removes the '\n' in the string (this is necessary to print art with a line and not in a column)
			// добавляет в массив строк "arrart" строку "str" до '\n'  и  Удаляет '\n' в строке (это нужно для печати арта с трочку а не в столбик)
			arrart = append(arrart, str[:len(str)-1])
			// updates the string "str" ​​to include the next string
			// обновляет строку "str" чтобы записать в нее следующую строку
			str = ""
			if count == 9 {
				// when the loop goes through 9 lines adds to the double array "res" the array "arrart", which contains 9 lines, i.e. 1 character from file.txt
				// когда цикл проходит 9 строк добавляет в двойной массив "res" массив "arrart", который содержит 9 строкб т.е 1 символ из file.txt
				res = append(res, arrart)
				// updates the array to store the next character from file.txt
				// обновляет массив чтобы записать в него следующий символ из file.txt
				arrart = []string{}
				// updates the string "str" ​​to include the next string
				// обновляет строку "str" чтобы записать в нее следующую строку
				str = ""
				count = 0
			}
		}
	}
	if countstrings != NumberOfStringInFile {
		return nil, ErrorContentChanged
	}
	return res, nil
}

// Output arguments in the desired art font
// Выводит аргументы нужным арт-шрифтом
func Compare(array [][]string, s []string) string {
	result := ""
	res := ""
	num := 0
	// count arguments that are separated by '\n'
	// подсчет аргументов которые раделенны '\n'
	j := 0

	if isOnlyNewline(s) {
		// len(s)-1 - removes the extra last '\n'
		// len(s)-1 - убирает лишний последний '\n'
		for i := 0; i < len(s)-1; i++ {
			// fmt.Println()
			result += "\n"
		}
		if result == "" {
			return result
		}

		return result[:len(result)-1]
	}

	for j < len(s) {
		// 8 character lines file.txt
		// 8 строк символа file.txt
		for i := 1; i <= 8; i++ {
			// "" is '\n' in the array, if you put '\n' between words, there will be one newline and not 8
			// ""- это '\n' в массиве , если между словами ставить '\n', будет один newline а не 8
			if s[j] == "" {
				// fmt.Println()
				result += "\n"
				break
			}
			for _, w := range s[j] {
				// equates the ascii character index of the argument to the file.txt character index
				// приравнивает индекс символа аргумента по ascii индексу символа file.txt
				num = int(w) - 32
				// writes to "res" the First Line of each Character in file.txt
				//  записывает в "res"  Первую Строчку каждого Символа file.txt
				res = array[num][i]
				// fmt.Print(res)
				result += res
			}
			// puts a newline after writing a line to start printing the second line, etc.
			//  после написания строчки ставит newline, чтобы начать печетать вторую строчку и т.д
			// fmt.Println()
			result += "\n"
		}
		j++
	}
	result = result[:len(result)-1]
	return result
}

// checks arguments if only '\n' are written to it
// проверяет аргументы если в него запиcаны только '\n'
func isOnlyNewline(s []string) bool {
	count := 0
	for _, v := range s {
		if v == "" {
			count++
		}
	}
	// if count == len(s) - 'true', if not 'false'
	// если count == len(s) - 'true', если нет 'false'
	return count == len(s)
}

func errchangeFont(s string) (string, error) {
	standartFont := "cmd/ascii-art/standard.txt"
	shadowFont := "cmd/ascii-art/shadow.txt"
	thinkertoyFont := "cmd/ascii-art/thinkertoy.txt"
	// @TODO: use switch and case

	switch s {
	case "standard":
		return standartFont, nil
	case "shadow":
		return shadowFont, nil
	case "thinkertoy":
		return thinkertoyFont, nil
	default:
		return "", fmt.Errorf("error: this %s font doesn't exist.\nUse: thinkertoy, shadow, standard", s)
	}
}
