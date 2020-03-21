package controller

import (
	"fmt"
	"github/httpx"
	"io/ioutil"
	"os"
	"regexp"
)

type AddReq struct {
	Tag    string `json:"tag_name"`
	Body   string `json:"body"`
	Branch string `json:"target_commitish"`
}

func List(token string) error {
	configByte, err := readFile(".git/config")
	if err != nil {
		return fmt.Errorf("不在git项目的根目录下")
	}
	r1 := regexp.MustCompile("url = https://github.com(.*)\\.git")
	result := r1.FindStringSubmatch(string(configByte))
	if len(result) > 0 {
		resp := make([]map[string]interface{}, 0)
		httpx.SendWithHeaders("GET", GetRelease(result[1]), nil, &resp, GetHeader(token))
		if len(resp) > 0 {
			for key := range resp {
				fmt.Printf("id:%d,tag_name:%s\n", resp[key]["url"])
			}
		}

	} else {
		return fmt.Errorf("不在git项目的根目录下")
	}
	return nil
}

func Add(req *AddReq, token string) error {
	configByte, err := readFile(".git/config")
	if err != nil {
		return fmt.Errorf("不在git项目的根目录下")
	}
	r1 := regexp.MustCompile("url = https://github.com(.*)\\.git")
	result := r1.FindStringSubmatch(string(configByte))
	if len(result) > 0 {
		resp := make(map[string]interface{})
		httpx.SendWithHeaders("POST", GetRelease(result[1]), req, &resp, GetHeader(token))
		fmt.Println(resp["url"])
	} else {
		return fmt.Errorf("不在git项目的根目录下")
	}
	return nil
}

func readFile(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}
	return fd, nil
}
