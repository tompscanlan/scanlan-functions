package main

import (
	"encoding/json"
)

type Request struct {
	People []string `json:"people"`
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

var people = []string{
	//	"Hite",
	//	"Irma",
	"Julio",
	"Luis",
	"Nagashree",
	"Neeraj",
	"Prasanthi",
	"Scanlan",
	//	"Sri",
}

func Main(in Request) (*Response, error) {
	if in.People == nil {
		in.People = people
	}

	stringBody, _ := json.Marshal(in.People)

	return &Response{
		Body: string(stringBody),
	}, nil
}
