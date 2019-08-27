package common

import "strings"

func expendSlice(someStringSlice []string, largeThan int) []string {
	e := make([]string, largeThan)
	d := append(someStringSlice, e...)
	return d
}

func transposeTwoDimensionArray(inputSlice [][]string) [][]string {
	var b int
	var c, d [][]string
	for _, i := range inputSlice {
		//fmt.Println(len(i))
		if len(i) > b {
			*&b = len(i)
		}
	}
	for _, i := range inputSlice {
		c = append(c, expendSlice(i, b-len(i)))
	}
	for i := 0; i < b; i++ {
		var e []string
		for j := 0; j < len(c); j++ {
			e = append(e, c[j][i])
		}
		d = append(d, e)
	}
	return d

}
func maxLength(inputSlice []string) int {
	var maxLen int
	for _, i := range inputSlice {
		iRune := []rune(i)
		currentLen := len(iRune) + extractChineseNum(i)
		if currentLen > maxLen {
			*&maxLen = currentLen
		}
	}
	return maxLen
}
func addNumStr(someString string, someNumber int) string {
	for i := 0; i < someNumber; i++ {
		someString = someString + " "
	}
	return someString
}

func TurnSliceToStringAddFrame(inputSlice [][]string) string {
	tranSlice := transposeTwoDimensionArray(inputSlice)
	//fmt.Println(a)
	for _, i := range tranSlice {
		max := maxLength(i)
		for j := 0; j < len(i); j++ {
			if len(i[j]) < max {
				i[j] = addNumStr(i[j], max-len(i[j]))
			}
		}
	}
	//fmt.Println(tranSlice)
	d := transposeTwoDimensionArray(tranSlice)
	var e string
	for _, i := range d {
		e = e + strings.Join(i, " | ") + "\n"
	}
	return e
}

func extractChineseNum(someStr string) int {
	runeStr := []rune(someStr)
	num := 0
	for i := 0; i < len(runeStr); i++ {
		charNum := runeStr[i]
		if charNum <= 40869 && charNum >= 19968 {
			num++
		}
	}
	return num
}
