package main

var Youdao = map[string]string{
	"url": "http://fanyi.youdao.com/translate?smartresult=dict&smartresult=rule",
	"content_type": "application/x-www-form-urlencoded",
	// 查询的语句 时间戳
	"sign_formate": "fanyideskweb" + "%s" + "%s" + "ebSeFb%=XZ%T[KZ)c(sy!",
	// 查询的语句 输入语种  翻译语种  时间戳   签名
	"query_formate": "i=" + "%s" + "&client=fanyideskweb&doctype=json&version=2.1&keyfrom=fanyi.web&action=FY_BY_CLICKBUTTION&typoResult=false&salt=" + "%s" + "&sign=" + "%s",
}