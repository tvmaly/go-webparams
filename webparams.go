package webparams

import (
	"fmt"
	"regexp"
	"strconv"
)

var int_regex *regexp.Regexp

func init() {
	int_regex, _ = regexp.Compile(`^\d+$`)
}

// ExtractInt64 takes a string that looks like an interger and returns an int64
// if it cannot convert the value to an int64 it returns 0 and an error
func ExtractInt64(param string) (int64, error) {

	if param == "" {
		return 0, fmt.Errorf("Error parameter is an empty string, does not look like an integer")
	}

	if int_regex.MatchString(param) == false {
		return 0, fmt.Errorf("Error parameter does not look like an integer in its string form")
	}

	val, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		return 0, fmt.Errorf("Error parameter could not be converted with strconv.ParseInt : %s", err)
	}

	return val, nil

}
