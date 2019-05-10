package utils

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpRequest struct {
	Url     string
	Param   map[string]string
	Headers map[string]string
	Body    []byte
}

type HttpResponse struct {
	Body []byte
}

func HttpGet(param *HttpRequest) (*HttpResponse, error){
	request, err := http.NewRequest("GET", param.Url, nil)
	if err != nil{
		log.Printf("utils.httpClient.HttpGet,http.NewRequest error,detail %s", err)
		return nil, err
	}

	if len(param.Headers) > 0 {
		for key, value := range param.Headers {
			request.Header.Set(key, value)
		}
	}

	body, err := doRequest(request)
	if err != nil {
		log.Printf("utils.httpClient.HttpPost,doRequest error,detail %s", err)
		return nil, err
	}
	response := &HttpResponse{
		Body: body,
	}
	return response, nil
}


func HttpPost(param *HttpRequest) (*HttpResponse, error) {
	byteBody := bytes.NewReader(param.Body)
	request, err := http.NewRequest("POST", param.Url, byteBody)
	if err != nil {
		log.Printf("utils.httpClient.HttpPost,http.NewRequest error,detail %s", err)
		return nil, err
	}

	if len(param.Headers) > 0 {
		for key, value := range param.Headers {
			request.Header.Set(key, value)
		}
	}

	body, err := doRequest(request)
	if err != nil {
		log.Printf("utils.httpClient.HttpPost,doRequest error,detail %s", err)
		return nil, err
	}
	response := &HttpResponse{
		Body: body,
	}
	return response, nil
}

func doRequest(param *http.Request) ([]byte, error) {
	var (
		resp *http.Response
		err  error
	)

	client := new(http.Client)

	resp, err = client.Do(param)
	if err != nil{
		log.Printf("utils.httpClient.doRequest,client.Do error, detail %s", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("utils.httpClient.doRequest,doRequest,ioutil.ReadAll error,detail %s", err)
		return nil, err
	}
	return body, nil
}
