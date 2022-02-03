package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Requester struct {
	Body      map[string]string
	Client    http.Client
	Headers   map[string]string
	Method    string
	URL       string
	URLParams map[string]string
}

func (r *Requester) DoRequest() ([]byte, error) {
	req, err := http.NewRequest(r.Method, r.URL, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")

	q := url.Values{}
	for k, v := range r.URLParams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	log.Printf("Requested: %s %s %s", req.Method, req.URL, resp.Status)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing body: %s", err)
		}
	}(resp.Body)
	return ioutil.ReadAll(resp.Body)
}

func (r *Requester) DoRequestTo(data interface{}) error {
	res, err := r.DoRequest()
	if err != nil {
		return err
	}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return err
	}
	return nil
}
