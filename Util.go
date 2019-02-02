package main

import (
	"crypto/md5"
			"fmt"
	"io"
)

func GetYouDaoSign(queryString, salt string) string {
	w := md5.New()
	str := fmt.Sprintf(Youdao["sign_formate"], queryString, salt)
	io.WriteString(w, str)
	md5Str := fmt.Sprintf("%x", w.Sum(nil))
	return md5Str
}