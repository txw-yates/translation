package main

import (
	translateHttp "./Http"
		"crypto/md5"
	"io"
	"fmt"
	"strings"
	"time"
	"math/rand"
	"encoding/json"
	)

var transferStr string

func main() {
	fmt.Printf(" %c[0;%dm%s%c[0m ", 0x1B, 34, "请输入要翻译的内容:\n", 0x1B)
	fmt.Scanln(&transferStr)

	// 有道翻译 md5加密
	w := md5.New()
	i := string(time.Now().Unix()) + string(rand.Intn(10) + 1)
	str := "fanyideskweb" + transferStr + i + "ebSeFb%=XZ%T[KZ)c(sy!"
	io.WriteString(w, str)
	md5Str := fmt.Sprintf("%x", w.Sum(nil))

	res := translateHttp.Post("http://fanyi.youdao.com/translate?smartresult=dict&smartresult=rule",
		"application/x-www-form-urlencoded",
		strings.NewReader("i=" + transferStr + "&from=zh-CHS&to=ENsmartresult=dict&client=fanyideskweb&doctype=json&version=2.1&keyfrom=fanyi.web&action=FY_BY_CLICKBUTTION&typoResult=false&salt=" + i + "&sign=" + md5Str))

	var dat map[string] interface{}
	if err := json.Unmarshal([]byte(res), &dat); err == nil {
		fmt.Printf(" %c[%d;%dm%s%c[0m ", 0x1B, 0, 31, "翻译结果为:\n", 0x1B)
		fmt.Println(dat["translateResult"].([]interface{})[0].([]interface{})[0].(map[string]interface{})["tgt"])
	} else {
		fmt.Printf(" %c[0;%dm%s%c[0m ", 0x1B, 31, "翻译失败!!!", 0x1B)
	}
}