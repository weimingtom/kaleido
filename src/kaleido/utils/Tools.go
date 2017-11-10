package Tools

import (
	"fmt"
)

func Sprintf(format string, parameters ...interface{}) string {
	return fmt.Sprintf(format, parameters...)
}

func Printf(format string, parameters ...interface{}) {
	fmt.Printf(format, parameters...)
}
