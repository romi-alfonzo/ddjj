package extract

import (
	"bufio"
	"regexp"
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

func stringToInt64(line string) int64 {
	value := strings.ReplaceAll(line, ".", "")
	i, _ := strconv.ParseInt(value, 10, 64)

	return i
}

func stringToInt(line string) int {
	i, _ := strconv.Atoi(line)

	return i
}

func stringToYear(line string) int {
	year := stringToInt(line)

	if year == 0 {
		return 0
	}

	if year < 100 {
		return 2000 + year
	}

	return year
}

func isDate(line string) bool {
	matched, _ := regexp.MatchString(`[0-9]{2}\/[0-9]{2}\/[0-9]{4}`, line)
	return matched
}

func isBarCode(line string) bool {
	matched, _ := regexp.MatchString(`[0-9]{5,6}-[0-9]{5,6}-[0-9]{2,3}`, line)
	return matched
}

func isNumber(line string) bool {
	line = strings.ReplaceAll(line, ".", "")
	_, err := strconv.ParseInt(line, 10, 64)

	return err == nil
}
