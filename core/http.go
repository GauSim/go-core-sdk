package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ClientRequest is to issue a http request agains the api
type ClientRequest struct {
	method string
	url    string
	data   interface{}
	format string
	auth   Auth
}

func generateJSONRequestData(cr ClientRequest) ([]byte, string, error) {
	body, err := json.Marshal(cr.data)
	if err != nil {
		return nil, "", err
	}
	return body, "application/json", nil
}

func (r ClientConfig) Request(cr ClientRequest, re interface{}) error {
	if r.Debug {
		log.Printf("[http] request: %s %v \n\n", cr.method, cr.url)
	}

	body, ct, err := generateJSONRequestData(cr)
	if err != nil {
		return err
	}

	if r.Debug {
		log.Printf("[http] body \n\n %s \n\n", string(body))
	}

	client := &http.Client{}
	req, err := http.NewRequest(cr.method, cr.url, bytes.NewReader(body))
	if err != nil {
		log.Fatalf("[http] http.NewRequest() error: %v\n", err)
	}
	req.Header.Add("Content-Type", ct)

	req.Header.Add("Authorization", fmt.Sprintf("%s %s", cr.auth.TokenType, cr.auth.AccessToken))
	req.Header.Add("X-Client-ID", r.Identifier)
	req.Header.Add("X-Client-Version", r.Version)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Charset", "utf-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("[http] httpGet() error: %v\n", err)
		return err
	}

	defer resp.Body.Close()
	c, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("[http] ioutil.ReadAll() error: %v\n", err)
		return err
	}

	if r.Debug {
		log.Printf("[http] request resp: \n\n %s \n\n", string(c))
	}

	var env envelope
	err2 := json.Unmarshal(c, &env)
	if err2 != nil {
		log.Fatalf("[http] envelope error: %v\n", err2)
		return err2
	}

	if env.Data != nil {
		c, _ = json.Marshal(env.Data)
	}
	return json.Unmarshal(c, &re)
}
