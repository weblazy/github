package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/pkg/errors"
)

const (
	applicationJson = "application/json;charset=UTF-8"
	applicationXml  = "application/xml"
)

func SendWithHeaders(method string, path string, body interface{}, v interface{}, headers ...map[string]string) (int, error) {
	if v != nil {
		rv := reflect.ValueOf(v)
		if err := validatePtr(&rv); err != nil {
			return 0, err
		}
	}
	var reader io.Reader
	if body != nil {
		byteData, err := json.Marshal(body)
		if err != nil {
			return 0, errors.WithMessagef(err, "marshal body failed: %+v", body)
		}
		reader = bytes.NewReader(byteData)
	}
	request, err := http.NewRequest(method, path, reader)
	if err != nil {
		return 0, err
	}
	if len(headers) > 0 {
		headerMap := headers[0]
		for k, value := range headerMap {
			request.Header.Set(k, value)
		}
	}
	request.Header.Set("Content-Type", applicationJson)
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return 0, errors.WithMessagef(err, "do request failed: %+v", request)
	}
	defer resp.Body.Close()

	if v != nil {
		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, errors.WithMessagef(err, "read body failed")
		}

		if err := json.Unmarshal(respBytes, v); err != nil {
			return resp.StatusCode, errors.WithMessagef(err, "unmarshal failed: %s", string(respBytes))
		}
	}
	return resp.StatusCode, nil
}

func validatePtr(v *reflect.Value) error {
	// sequence is very important, IsNil must be called after checking Kind() with reflect.Ptr,
	// panic otherwise
	if !v.IsValid() || v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("error: not a valid pointer: %v", v)
	}

	return nil
}

func PostForm(path string, formData url.Values, v interface{}) (int, error) {
	resp, err := http.PostForm(path, formData)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if v != nil {
		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, errors.WithMessagef(err, "read body failed")
		}
		if err = json.Unmarshal(respBytes, v); err != nil {
			return resp.StatusCode, errors.WithMessagef(err, "unmarshal failed: %s", string(respBytes))
		}
	}
	return resp.StatusCode, nil
}
