package common

import (
	"strings"
)

func ClickHouseQuery(conHttp, SQL string) [][]string {
	var resultSlice [][]string
	result := Post(conHttp, SQL)
	for _, i := range strings.Split(result, "\n") {
		if i != "" {
			resultSlice = append(resultSlice, strings.Split(i, "\t"))
		}
	}
	return resultSlice
}
