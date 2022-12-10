package main

import "strings"

// test 可以寫各種字串拼接的比較
// 可參考 https://geektutu.com/post/hpg-string-concat.html
// len => 輸出字串的長度（先分配效能會提高）, s => 欲連結字串
func WriteStringByBuilder(len int, s ...string) string {
	var builder strings.Builder
	builder.Grow(len)
	for _, v := range s {
		builder.WriteString(v)
	}
	return builder.String()
}
