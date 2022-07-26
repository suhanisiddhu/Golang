package Driven

import (
	"strings"
)

func drive(str string) string {
	res := strings.ToUpper(str)
	if str != "" && str[len(str)-1:] == "?" && str != res {
		return "Sure"
	} else if str != "" && res == str && str[len(str)-1:] == "?" {
		return "Calm down, I know what I'm doing!"
	} else if strings.TrimSpace(str) == "" {
		return "Fine. Be that way!"
	} else if str == res {
		return "Whoa, chill out!"
	} else {
		return "Whatever"
	}

}
