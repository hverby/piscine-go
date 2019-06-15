package piscine

import "strings"

func Capitalize(word string) string {
	res := strings.Title(ToLower(word))
	i := Index(word, "_")
	if i == -1 {
		return res
	} else {
		res = strings.Replace(res, string(res[i+1]), ToUpper(string(res[i+1])), i+1)
		return res
	}
}
