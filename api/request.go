package api

import (
	"net/http"
	"strings"
)

type Request struct {
	r     http.Request
	model *any
}

func MapToQueryParams(paramsMap map[string]string) string {
	if len(paramsMap) == 0 {
		return ""
	}

	var sBuilder = strings.Builder{}

	for p, pVal := range paramsMap {
		sBuilder.WriteString(p + "=" + pVal + "&")
	}
	res := sBuilder.String()
	return res[:len(res)-1]
}
