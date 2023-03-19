package tower

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetComputeEnv() {

	url := "https://sm-k8s.dev-tower.net/api/compute-envs?workspaceId=251085962711837"
	token := "eyJ0aWQiOiA2fS5kNTQ1MDBkNzc4MGM4OWQ0YmYzYWFjYTQ3NTIwZGExN2EyNTAyZDAw"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}
