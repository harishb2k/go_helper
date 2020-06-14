package helper

import "github.com/fatih/color"

func Stringify(input string) (result string, err error) {
	color.Cyan("v1.0.3", input)
	return input + "*-", nil
}
