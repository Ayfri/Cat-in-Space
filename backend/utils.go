package main

import (
	"encoding/json"
	"fmt"
	"html/template"
)

func ToJSON(v interface{}) template.JS {
	r, err := json.Marshal(v)
	if err != nil {
		return template.JS("Cannot convert to JSON : " + fmt.Sprint(v))
	}
	return template.JS(r)
}
