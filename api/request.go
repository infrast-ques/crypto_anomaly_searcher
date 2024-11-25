package api

import (
	"net/http"
	"strings"
)

type Request[T any] struct {
	Req   http.Request
	Model T
}

func MapToQueryParams(paramsMap map[string]string) string {
	if len(paramsMap) == 0 {
		return ""
	}

	var sBuilder = strings.Builder{}
	for p, pVal := range paramsMap {
		sBuilder.WriteString(p + "=" + pVal + "&")
	}

	return sBuilder.String()
}
