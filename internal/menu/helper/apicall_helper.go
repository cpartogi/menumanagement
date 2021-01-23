package helper

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// APICall struct
type APICall struct {
	URL       string
	Method    string
	FormParam string
	Header    map[string]string
}

type UrlHttpResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Body       string      `json:"body"`
	Header     http.Header `json:"header"`
}

type APIResponse struct {
	Code          string            `json:"code"`
	Error         []string          `json:"error"`
	Data          interface{}       `json:"data"`
	Message       string            `json:"message"`
	MessageDetail map[string]string `json:"message_detail,omitempty"`
	Status        string            `json:"status"`
}

// Call call to third party endpoint
func (apicall *APICall) Call() (UrlHttpResponse, error) {
	var result UrlHttpResponse

	// Http request
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	req, err := http.NewRequest(apicall.Method, apicall.URL, bytes.NewBuffer([]byte(apicall.FormParam)))
	if err != nil {
		log.Println("Error new request | UrlHelper ApiCall ", err)
		return result, err
	}

	// Set header
	// -- Content type
	req.Header.Add("Content-Type", "application/json")
	for index, value := range apicall.Header {
		req.Header.Add(index, value)
	}

	// Do request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%+v\n", err)
		fmt.Println(err.Error())
		log.Println("Error doing request | UrlHelper ApiCall")
		return result, err
	}

	defer resp.Body.Close()

	// Get string body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error read body")
		log.Println(err.Error())
		return result, err
	}

	result.Status = resp.Status
	result.StatusCode = resp.StatusCode
	result.Body = string(body)
	result.Header = resp.Header

	return result, nil
}
