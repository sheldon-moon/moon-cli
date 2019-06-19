package http

import (
	"io/ioutil"
	"log"
	"net/http"
)

func PostRequest(uri string, params map[string]interface{}) {

}

func FormRequest(uri string, params map[string]interface{}) {

}

func GetRequest(uri string, params map[string]interface{}) interface{} {
	if len(params) > 0 {
		for k , v := range params{
			k = k+"="+v.(string)
		}
	}
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal("get request error !", "uri : ", uri, "param :", params, err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("parse request error !", "uri : ", uri, "body :", resp, err)
		return nil
	}
	return string(body)
}
