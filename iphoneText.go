package common

import (
	"fmt"
	"github.com/tidwall/sjson"
)

func IphoneText(apiUrl string, personalMobile []string, iText string) {
	iData := `{
    "mobiles":[
        13683368766,
        13011001183
    ],
    "context":"test for alarm"
}`
	iData, _ = sjson.Set(iData, "mobiles", personalMobile)
	iData, _ = sjson.Set(iData, "context", iText)
	result := Post(apiUrl, iData)
	fmt.Println(result)
}
