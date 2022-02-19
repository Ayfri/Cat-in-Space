package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Pair struct {
	Key   string
	Value string
}

type Requester struct {
	Body           map[string]string
	Client         http.Client
	Headers        map[string]string
	Method         string
	URL            string
	URLParamsArray []Pair
	URLParams      map[string]string
}

func (r *Requester) DoRequest() ([]byte, error) {
	req, err := http.NewRequest(r.Method, r.URL, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	q := url.Values{}
	for k, v := range r.URLParams {
		q.Add(k, v)
	}
	for _, p := range r.URLParamsArray {
		q.Add(p.Key, p.Value)
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
	if resp.StatusCode == http.StatusBadRequest {
		res, _ := io.ReadAll(resp.Body)
		log.Printf("bad request: %s", string(res))
	}
	return io.ReadAll(resp.Body)
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
