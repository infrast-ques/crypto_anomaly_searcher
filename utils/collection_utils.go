package utils

import (
	"fmt"
	"strings"
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

// Flatten todo допилить
func Flatten[T [][]interface{}](lists T) []interface{} {
	var res []interface{}
	for _, list := range lists {
		for _, item := range list {
			res = append(res, item)
		}
	}
	return res
}

func SliceFilter(tickers []string, predicate func(string) bool) []string {
	var res []string
	for _, ticker := range tickers {
		if predicate(ticker) {
			res = append(res, ticker)
		}
	}
	return res
}

func FloatToStrFmt(number float64) string {
	return strings.Replace(fmt.Sprintf("%.2f", number), ".", ",", 1)
}

func StrListToStr(anyList []string) string {
	var sb strings.Builder
	for _, s := range anyList {
		sb.WriteString(fmt.Sprintf("%s\n", s))
	}
	return sb.String()
}

func ConvToISlice[T any](data []T) []interface{} {
	result := make([]interface{}, len(data))
	for i, v := range data {
		result[i] = v
	}
	return result
}
