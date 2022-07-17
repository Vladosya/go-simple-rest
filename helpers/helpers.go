package helpers

import (
	"strconv"
	"strings"
)

func StrToInt(str string) (int, error) { // функция приведения типа из string в int
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}
