package generatestring

import (
	"strings"
)

var param = "USOMAANAPAIUMASYDNIP"

func GenerateWord(to string) []string {
	j, i := 0, 0
	newString := ""
	result := []string{}
	param = strings.ToLower(param)
	for len(param) != i {

		if param[i] == to[j] {
			newString += string(param[i])
			param = param[:i] + param[i+1:]
			j++
			i = 0
			if newString == to {
				result = append(result, newString)
				newString = ""
				i, j = 0, 0
			}
		}
		i++

	}

	return result
}
