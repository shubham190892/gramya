package utils

import "strings"

func QueryParams(uri string) map[string]string {
	queryMap := make(map[string]string)
	qs := strings.Split(uri, "?")
	if len(qs) != 2 {
		return queryMap
	}
	for _, e := range strings.Split(qs[1], "&") {
		param := strings.Split(e, "=")
		queryMap[param[0]] = param[1]
	}
	return queryMap
}
