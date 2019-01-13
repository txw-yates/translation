package Http

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
)

func Get(url string) {
	res, err := http.Get(url)
	if err != nil {
		// todo 错误捕获
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// todo 错误捕获
	}

	fmt.Println(string(body))
}

func Post(url string, contentType string, params io.Reader) []byte {
	res, err := http.Post(url,
		contentType,
		params)
	if err != nil {

	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// todo 错误捕获
	}

	return body
}