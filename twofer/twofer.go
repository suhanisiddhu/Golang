package twofer

import "strings"

func name(x string) bool {
	if strings.TrimSpace(x) == "" {
		return false
	} else {
		return true
	}

}
