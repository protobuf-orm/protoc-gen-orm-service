package ast

import "strings"

func getName(fullname string) string {
	i := strings.LastIndex(fullname, ".")
	if i < 0 {
		return fullname
	}
	return fullname[i+1:]
}
