package utils

import (
	"errors"
	"regexp"
)

var letter = regexp.MustCompile(`^[a-zA-Z]*$`).MatchString
var dateYYYYMMDDwithHyphen = regexp.MustCompile(`^[0-9]{4}-(0[0-9]|1[0-2])-([0-2][0-9]|3[0-1])$`).MatchString
var commaSeparated = regexp.MustCompile(`^[a-zA-Z, ]*$`).MatchString

func NullOrLetter(s string) bool {
	return letter(s)
}

func NotNullAndLetter(s string) bool {
	return len(s) != 0 && letter(s)
}

func NullOrCommaSeparatedLetter(s string) bool {
	return commaSeparated(s)
}

func NotNullAndDateFormat(s string) bool {
	return dateYYYYMMDDwithHyphen(s)
}

func Panic(err error) {
	if err != nil {
		panic(errors.New(err.Error()))
	}
}
