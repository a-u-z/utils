package main

func reverseStrRecursion(input string) string {
	runes := []rune(input)
	if len(runes) <= 1 {
		return string(runes)
	}
	return reverseStrRecursion(string(runes[1:])) + string(runes[0])
}
