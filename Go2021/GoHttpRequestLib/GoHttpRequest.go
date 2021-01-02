/**
 * @Author: ZhiWei.Yang
 * @Date:   2021/1/2 8:53 PM
 */

package GoHttpRequestLib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func HttpPost(baseurl string, endpoint string, body interface{}, headers map[string]string, cookies map[string]string) (*http.Response, error) {
	// https://juejin.im/post/6844903859312132103

	var bodyJson []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			log.Println("json.Marshal err: %v", err)
			return nil, errors.New("json.Marshal(body) failed")
		}
	}
	url := fmt.Sprintf("%s%s", baseurl, endpoint)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println("http.NewRequest err: %v", err)
		return nil, errors.New("NewRequest failed ")

	}
	req.Header.Set("Content-type", ContentType)

	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	if cookies != nil {
		for name, value := range cookies {
			req.AddCookie(&http.Cookie{Name: name, Value: value})
		}
	}
	defer req.Body.Close()

	client := &http.Client{}
	log.Printf("Go %s URL : %s", http.MethodPost, req.URL.String())

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, errors.New("client.Do(req) failed ")
	}
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("HttpPost--> endpoint: %v\nRequestBody: %v\nResponse: %v\n", endpoint, body, string(data))
	return res, err

}

func HttpGet(baseurl string, endpoint string, params map[string]interface{}, headers map[string]string, cookies map[string]string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", baseurl, endpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("NewRequest failed ")
	}

	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			if s, ok := val.(int); ok {
				q.Add(key, strconv.Itoa(s))
			} else if s, ok := val.(string); ok {
				q.Add(key, s)
			} else {
				return nil, nil
			}
		}
	}

	req.Header.Set("Content-type", ContentType)
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	if cookies != nil {
		for name, value := range cookies {
			req.AddCookie(&http.Cookie{Name: name, Value: value})
		}
	}

	client := &http.Client{}
	log.Printf("Go %s URL : %s", http.MethodGet, req.URL.String())
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, errors.New("client.Do(req) failed ")
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("HttpGet--> endpoint: %v\nRequestBody: %v\nResponse: %v\n", endpoint, params, string(data))

	return res, err
}
