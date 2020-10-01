package main

import (
	"os"
	"strings"
)

var (
	badWords []string
)

func loadBadWords() {
	envWords := os.Getenv("BADWORDS")
	badWords = strings.Split(envWords, ",")
}
