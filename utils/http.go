package utils

import (
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

func HttpGetJsonRes(url string, result interface{}) error {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != http.StatusOK {
		return err
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		return err
	}

	return nil

}
