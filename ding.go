package common

import (
	"fmt"
	"github.com/tidwall/sjson"
)

func Ding(apiUrl, title, text string, mobile []string) string {
	data := `{
     "msgtype": "markdown",
     "markdown": {
         "title":"",
         "text":"",
     },
     "at": {
         "atMobiles": [
         ], 
         "isAtAll": false
     }
}`
	data, _ = sjson.Set(data, "markdown.title", title)
	if len(mobile) > 0 {
		for _, v := range mobile {
			text += fmt.Sprintf("@%s", v)
		}
	}
	data, _ = sjson.Set(data, "markdown.text", text)
	data, _ = sjson.Set(data, "at.atMobiles", mobile)
	repoB := Post(apiUrl, data)
	return repoB

}
