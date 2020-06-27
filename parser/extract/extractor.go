package extract

import (
	"bufio"
	"strconv"
	"strings"
)

func moveUntil(scanner *bufio.Scanner, search string, exact bool) *bufio.Scanner {
	for scanner.Scan() {
		line := scanner.Text()

		if exact {
			if line == search {
				break
			}
		} else {
			if strings.Contains(line, search) {
				break
			}
		}

	}

	return scanner
}

func getInt(scanner *bufio.Scanner, precedence string, skip int) int {
	value := getString(scanner, precedence, skip)

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return valueInt
}

func getString(scanner *bufio.Scanner, precedence string, skip int) string {
	var value string
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, precedence) {
			for i := 0; i < skip; i++ {
				scanner.Scan()
			}
			value = scanner.Text()

			break
		}
	}

	return strings.TrimSpace(value)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getTotalInCategory(scanner *bufio.Scanner) int64 {
	scanner.Scan()
	scanner.Scan()
	line := strings.ReplaceAll(scanner.Text(), ".", "")
	i, _ := strconv.ParseInt(line, 10, 64)

	return i
}
