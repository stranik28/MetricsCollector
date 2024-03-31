package server

import "net/http"

func SendReq(url string, method string) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	code := resp.StatusCode
	if code != 200 {
		panic("Server code not 200")
	}
}
