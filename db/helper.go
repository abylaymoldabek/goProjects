package db

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/abylaymoldabek/redisExample/models"
)

type ResourceError struct {
	URL      string
	HTTPCode int
	Message  string
	Body     interface{}
	Err      error `json:"-"`
}

func (re *ResourceError) Error() string {
	return fmt.Sprintf(
		"Resource error: URL: %s, status code: %v,  err: %v, body: %v",
		re.URL,
		re.HTTPCode,
		re.Err,
		re.Body,
	)
}

func RequestJSON(method, url string, data []byte, headers map[string]string, responseStruct interface{}) (httpStatus int, responseBody []byte, err error) {
	if headers == nil {
		headers = map[string]string{"Content-Type": "application/json"}
	} else {
		headers["Content-Type"] = "application/json"
	}

	httpStatus, responseBody, err = send(method, url, "", data, headers)
	if err != nil {
		return
	}
	if responseStruct != nil && len(responseBody) != 0 {
		err = json.Unmarshal(responseBody, responseStruct)
	}
	return
}

func RequestXML(method, url string, data []byte, headers map[string]string, responseStruct interface{}) (httpStatus int, responseBody []byte, err error) {
	if headers == nil {
		headers = map[string]string{"Content-Type": "text/xml"}
	} else {
		headers["Content-Type"] = "text/xml"
	}

	httpStatus, responseBody, err = send(method, url, "", data, headers)
	if err != nil {
		return
	}
	if responseStruct != nil && len(responseBody) != 0 {
		err = xml.Unmarshal(responseBody, responseStruct)
	}
	return
}

func send(method, urlString, token string, data []byte, headers map[string]string) (httpStatus int, buf []byte, err error) {
	client := &http.Client{}
	request, err := http.NewRequest(method, urlString, bytes.NewBuffer(data))
	if err != nil {
		return httpStatus, nil, &ResourceError{URL: urlString, Err: err}
	}

	//Отрабатываем по header
	for key, value := range headers {
		request.Header.Add(key, value)
	}

	//Отрабатываем по авторизацию
	if token != "" {
		request.Header.Add("Authorization", token)
	}

	//Отрабатываем по параметрам
	if strings.ContainsAny(urlString, "?") {
		urlTemp, err := url.Parse(urlString)
		if err != nil {
			return httpStatus, nil, &ResourceError{URL: urlString, Err: err}
		}
		urlQuery := urlTemp.Query()
		urlTemp.RawQuery = urlQuery.Encode()
		urlString = urlTemp.String()
	}

	response, err := client.Do(request)
	if err != nil {
		return httpStatus, nil, &ResourceError{URL: urlString, Err: err}
	}
	defer response.Body.Close()

	buf, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return httpStatus, nil, &ResourceError{URL: urlString, Err: err, HTTPCode: response.StatusCode}
	}

	httpStatus = response.StatusCode
	if response.StatusCode > 399 {
		return httpStatus, buf, &ResourceError{
			URL:      urlString,
			Err:      fmt.Errorf("incorrect status code"),
			HTTPCode: response.StatusCode,
			Message:  "incorrect response.StatusCode",
			Body:     string(data),
		}
	}

	return
}

func RequestData(urlSite string) (ResponseData models.ResponseData, err error) {
	var res models.ResponseData
	_, respbody, err := RequestXML("GET", urlSite, nil, nil, &res)
	if err != nil {
		fmt.Println(string(respbody))
		return
	}
	return res, nil
}
