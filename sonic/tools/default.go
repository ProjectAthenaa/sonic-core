package tools

import "strings"

func DefaultStr(val string, d string) string {
	if strings.EqualFold(val, "") {
		return d
	}
	return val
}
