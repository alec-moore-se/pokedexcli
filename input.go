package main

import "strings"

func cleanInput(text string) []string {
	temp := strings.ToLower(text)
	return strings.Fields(temp)
}
