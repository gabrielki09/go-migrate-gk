package scaffold

import (
	"log"
	"strings"
)

type Parts string

func buildPathNameByParser() string {
	return ""
}

// O que vier de -model com \
func parserName(fullName string) string {
	parts := strings.Split(fullName, "\\")

	log.Println(parts)

	return ""
}
