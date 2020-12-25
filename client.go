package amply

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type client struct {
	url, accessToken string
}

type Response struct {
	StatusCode int
	Body       string
}

func checkResponse(httpResp *http.Response, data []byte) (*Response, error) {
	statusCode := httpResp.StatusCode

	resp := Response{
		StatusCode: statusCode,
		Body:       string(data),
	}

	if statusCode == 204 {
		return &resp, nil
	} else if statusCode == 301 || statusCode == 302 {
		return &resp, nil
	} else if statusCode == 401 {
		return &resp, errors.New("Unauthorized")
	} else if statusCode == 403 {
		return &resp, errors.New("Forbidden")
	} else if statusCode == 404 {
		return &resp, errors.New("The resource was not found while making an API request")
	} else if statusCode == 422 {
		return &resp, errors.New("A validation error occurred while making an API request")
	} else if statusCode < 200 || statusCode >= 300 {
		return &resp, errors.New("An error occurred while making an API request")
	}

	return &resp, nil
}

func (e client) post(path string, data interface{}) (*Response, error) {
	httpClient := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	postData, _ := json.Marshal(data)
	request, err := http.NewRequest("POST", e.url+path, bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+e.accessToken)

	resp, _ := httpClient.Do(request)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return checkResponse(resp, body)
}
