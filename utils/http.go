package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JsonResponse struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
}

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

func HttpGetJsonRes(url string, target interface{}) (error) {
    res := JsonResponse{
        Data: target,
    }

    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    return json.NewDecoder(resp.Body).Decode(&res)
}


func HttpPostJsonRes(url string, body interface{},target interface{}) (error) {
	// Convert body to JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Send HTTP POST request
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

		// Check HTTP response status code
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("bad status: %s", bodyBytes)
	}

	// Parse JSON response
	res := JsonResponse{
		Data: target,
	}

    return json.NewDecoder(resp.Body).Decode(&res)
}