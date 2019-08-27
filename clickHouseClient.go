package common

import (
	"fmt"
	"net/url"
	"strings"
)

func ClickHouseQuery(conHttp, SQL string) [][]string {
	var resultSlice [][]string
	v := url.Values{}
	v.Add("query", SQL)
	body := v.Encode()
	baseUrl := conHttp
	queryUrl := body
	myUri := fmt.Sprintf("%s/?%s", baseUrl, queryUrl)
	result := Get(myUri)
	for _, i := range strings.Split(result, "\n") {
		if i != "" {
			resultSlice = append(resultSlice, strings.Split(i, "\t"))
		}
	}
	return resultSlice
}
