package tool

import (
	"github.com/axgle/mahonia"
	"regexp"
	"strings"
)

func Decode2str(b []byte, dec mahonia.Decoder)string{
	_, temp, _ := dec.Translate(b, true)
	return string(temp)
}

// ClearXmlinRun 将一个xml内容，清洗成一行一行的，这里可能会去掉一行最开始和最后的空格; 空号交给run cmd去解决吧
func ClearXmlinRun(s string) []string {
	ss := strings.Split(s, "\n")
	for i:=0;i<len(ss);i++ {
		reg := regexp.MustCompile("^[ \u3000\t\r\n]*")
		ss[i] = reg.ReplaceAllString(ss[i], "")

		reg = regexp.MustCompile("[ \u3000\t\r\n]*$")
		ss[i] = reg.ReplaceAllString(ss[i], "")
	}
	return ss
}