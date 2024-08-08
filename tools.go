package tools

import (
	"fmt"
	"strings"
)

func init() {
	fmt.Println("require tools")
}

// StrTrim 去除空白
func StrTrim(str string) string {
	return strings.Trim(str, " ")
}
