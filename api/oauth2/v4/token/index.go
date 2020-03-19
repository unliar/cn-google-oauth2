package handler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// 代理google oauth
func Handler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	req, _ := http.NewRequest("POST", "https://www.googleapis.com/oauth2/v4/token", bytes.NewBuffer(reqBody))
	req.Header.Set("accept", "application/json")
	client := &http.Client{}
	client.Timeout = 15 * time.Second
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body))
}
