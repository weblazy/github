package controller

import (
	"encoding/json"
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
type EditReq struct {
	Id     string `json:"id"`
	Tag    string `json:"tag_name"`
	Body   string `json:"body"`
	Branch string `json:"target_commitish"`
}

type DeleteReq struct {
	Id string `json:"id"`
}

func List(token string) error {
	path := getPath()
	if path == "" {
		return fmt.Errorf("不在git项目的根目录下")
	}
	resp := make([]map[string]interface{}, 0)
	_, err := httpx.SendWithHeaders("GET", GetRelease(path), nil, &resp, GetHeader(token))
	if err != nil {
		return err
	}
	if len(resp) > 0 {
		for key := range resp {
			fmt.Printf("id:%.0f,tag_name:%s\n", resp[key]["id"], resp[key]["tag_name"])
		}
	}
	return nil
}

func Add(req *AddReq, token string) error {
	path := getPath()
	if path == "" {
		return fmt.Errorf("不在git项目的根目录下")
	}
	resp := make(map[string]interface{})
	_, err := httpx.SendWithHeaders("POST", GetRelease(path), req, &resp, GetHeader(token))
	if err != nil {
		return err
	}
	fmt.Printf("id:%.0f,tag_name:%s\n", resp["id"], resp["tag_name"])
	return nil
}

func Edit(req *EditReq, token string) error {
	path := getPath()
	if path == "" {
		return fmt.Errorf("不在git项目的根目录下")
	}
	resp := make(map[string]interface{})
	_, err := httpx.SendWithHeaders("PATCH", GetRelease(path)+"/"+req.Id, req, &resp, GetHeader(token))
	if err != nil {
		return err
	}
	fmt.Printf("id:%.0f,tag_name:%s\n", resp["id"], resp["tag_name"])
	return nil
}

func Delete(req *DeleteReq, token string) error {
	path := getPath()
	if path == "" {
		return fmt.Errorf("不在git项目的根目录下")
	}
	resp := make(map[string]interface{})
	_, err := httpx.SendWithHeaders("DELETE", GetRelease(path)+"/"+req.Id, nil, &resp, GetHeader(token))
	if err != nil {
		return nil
	}
	s, _ := json.Marshal(resp)
	fmt.Println(string(s))
	return nil
}

func getPath() string {
	configByte, err := readFile(".git/config")
	if err != nil {
		return ""
	}
	r1 := regexp.MustCompile("url = https://github.com(.*)\\.git")
	result := r1.FindStringSubmatch(string(configByte))
	if len(result) > 0 {
		return result[1]
	} else {
		return ""
	}
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
