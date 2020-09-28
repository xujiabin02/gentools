package common

import "time"

func GetNowStr() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

}

type StrNow struct {
	name    string
	Str     string
	tFormat string
	Unix    int64
	OneDay  int64
	proper  []strTime
}
type strTime struct {
	name  string
	value int64
}

func (sn *StrNow) init() {
	sn.OneDay = 86400
	sn.tFormat = "2006-01-02 15:04:05"
	sn.Unix = time.Now().Unix()
	sn.Str = time.Unix(sn.Unix, 0).Format(sn.tFormat)
	sn.proper = []strTime{
		{"S", 1},
		{"M", 60},
		{"H", 3600},
		{"d", 86400},
		{"w", 604800},
		{"m", 18144000},
		{"Y", 31536000},
	}
}
func (sn *StrNow) past(num int64, strTime string) string {
	result := ""
	for _, k := range sn.proper {
		if k.name == strTime {
			result = time.Unix(sn.Unix+(num*k.value), 0).Format(sn.tFormat)
		}
	}
	return result
}
