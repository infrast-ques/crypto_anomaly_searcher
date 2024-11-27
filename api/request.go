package api

import (
	"strings"
)

const (
	binanceHost         = "api.binance.com"
	tickers24hrEndpoint = "/api/v3/ticker/24hr"
	tickersEndpoint     = "/api/v3/ticker"
)

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

func AsQueryParamList(list []string) string {
	return "[\"" + strings.Join(list, "\",\"") + "\"]"
}
