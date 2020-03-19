package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("https://oauth2.googleapis.com/tokeninfo?%s", r.URL.RawQuery)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("accept", "application/json")
	client := &http.Client{}
	client.Timeout = 15 * time.Second
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body))
}
