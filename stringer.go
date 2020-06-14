package helper

import "github.com/fatih/color"

func Stringify(input string) (result string, err error) {
	color.Cyan("1.0.2", input)
	return input + "*-", nil
}
