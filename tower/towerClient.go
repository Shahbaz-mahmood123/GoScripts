package tower

import (
	"io/ioutil"
	"net/http"
)

func httpClient(uri string) (string, error) {

	url := "https://sm-k8s.dev-tower.net/api/" + uri
	token := "eyJ0aWQiOiA2fS5kNTQ1MDBkNzc4MGM4OWQ0YmYzYWFjYTQ3NTIwZGExN2EyNTAyZDAw"

	req, err := http.NewRequest("GET", url, nil)
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
