package common

import (
	"io/ioutil"
	"net/http"
)

func GetGitignoreIo(ext string) string {
	url := "https://www.toptal.com/developers/gitignore/api/" + ext

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray)
}
