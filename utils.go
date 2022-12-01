package main

import (
	"strings"
)

func GetValue(src, keyword string) string {
	if strings.HasPrefix(src, keyword) {
		return strings.Trim(src, keyword)
	}
	return ""
}
