package piscine

import (
	"regexp"
	"strings"
	"unicode"
)

func Rot14(str string) string {
	fact := 14
	res := ""
	for _,v := range str{
		re := regexp.MustCompile("^[0-9_]*$")

		if string(v) == " "{
			res += " "
		}else if(re.MatchString(string(v))){
				
		} else{
			if unicode.IsLower(v) {
				codeA := 'a'
				codeL := 'l'
				codeZ := 'z'
				if v >= codeA && v <= codeZ {
					if v > codeL {
						nb1 := v + int32(fact)
						if nb1 == 123{
							res += strings.ToLower(string(codeA))
						}else{
							nb2 := nb1 - codeZ
							nb2 = (codeA + nb2) - 1
							res += strings.ToLower(string(nb2))
						}
					}else{
						res += strings.ToLower(string(v + 14))
					}
				}
			}else{
				codeA := 'A'
				codeL := 'L'
				codeZ := 'Z'
				if v >= codeA && v <= codeZ {
					if v > codeL {
						nb1 := v + int32(fact)
						if nb1 == 123{
							res += strings.ToUpper(string(codeA))
						}else{
							nb2 := nb1 - codeZ
							nb2 = (codeA + nb2) - 1
							res += strings.ToUpper(string(nb2))
						}
					}else{
						res += strings.ToUpper(string(v + 14))
					}
				}
			}

		}

	}
return res
}

