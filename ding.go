package per

import (
	"fmt"
	"github.com/tidwall/sjson"
)

func Ding(title, text, apiUrl string, mobile []string) string {
	data := `{
     "msgtype": "markdown",
     "markdown": {
         "title":"",
         "text":"",
     },
     "at": {
         "atMobiles": [
             "18312561610",
             "18702861755",
             "18758879303"
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
	//apiUrl:= `https://oapi.dingtalk.com/robot/send?access_token=429597fac225a147302a50073a2607bced30b8eaf5b525725facfd637662049a`
	//apiUrl = `https://oapi.dingtalk.com/robot/send?access_token=2b86a92bc850ffa298820208824fa0dde5f90ea3d8780b57554c3e5e0f57a9a2`
	repoB := Post(apiUrl, data)
	return repoB

}
