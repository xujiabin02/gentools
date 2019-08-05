package common

import (
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile(fileName string) string {
	var resultRturn string
	//fileName := "./logconfig.json"
	if fileObj, err := os.Open(fileName); err == nil {
		defer fileObj.Close()
		if contents, err := ioutil.ReadAll(fileObj); err == nil {
			resultRturn = strings.Replace(string(contents), "\n", "", 1)
			//fmt.Println("Use os.Open family functions and ioutil.ReadAll to read a file contents:",result)
		}

	}
	return resultRturn
}
