package helper

import "github.com/fatih/color"

func Stringify(input string) (result string, err error) {
    color.Cyan(input)
    return input + "*", nil
}
