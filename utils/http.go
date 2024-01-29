package utils

import (
	"bff/model/common/response"
	"bff/model/system"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HttpRequest(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		//c.Status(http.StatusServiceUnavailable)
		return "", err
	}
	return getStrfromResponse(response), nil

}

func getStrfromResponse(response *http.Response) string {
	reader := response.Body
	defer reader.Close()
	body, _ := io.ReadAll(reader)
	bodystr := string(body)
	fmt.Println(bodystr)
	return bodystr
}

func HttpGetJsonRes(url string) (*response.Response, error) {
	var posts []system.Post
	result := response.Response{
		Data: &posts,
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	} else if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("backend went wrong: %d", res.StatusCode)
		return nil, err
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil

}
