package tower

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func httpClient(uri string, method string, payload []byte) (string, error) {
	url := "https://sm-k8s.dev-tower.net/api/" + uri
	token := "eyJ0aWQiOiA2fS5kNTQ1MDBkNzc4MGM4OWQ0YmYzYWFjYTQ3NTIwZGExN2EyNTAyZDAw"

	var req *http.Request
	var err error

	if method == "GET" {
		req, err = http.NewRequest(method, url, nil)
	} else if method == "POST" {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
	} else {
		return "", errors.New("Invalid HTTP method")
	}

	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
