package main

import (
	translateHttp "./Http"
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

	salt := string(time.Now().Unix()) + string(rand.Intn(10) + 1)
	res := translateHttp.Post(
		Youdao["url"],
		Youdao["content_type"],
		strings.NewReader(fmt.Sprintf(Youdao["query_formate"], transferStr, salt, GetYouDaoSign(transferStr, salt))))

	var dat map[string] interface{}
	if err := json.Unmarshal([]byte(res), &dat); err == nil {
		fmt.Printf(" %c[%d;%dm%s%c[0m ", 0x1B, 0, 31, "翻译结果为:\n", 0x1B)
		fmt.Println(dat["translateResult"].([]interface{})[0].([]interface{})[0].(map[string]interface{})["tgt"])
	} else {
		fmt.Printf(" %c[0;%dm%s%c[0m ", 0x1B, 31, "翻译失败!!!", 0x1B)
	}
}