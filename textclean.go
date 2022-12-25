package main

import "strings"

func cleanContentBytes(txt []byte) string {
	text := string(txt)
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "\n", "")
	return text
}

func CleanContentText(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "\n", "")
	return text
}
