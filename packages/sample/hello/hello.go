package main

import (
	"fmt"
	"os"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func Main(in Request) (*Response, error) {
	if in.Name == "" {
		in.Name = "stranger"
	}
	var env string
	for _, e := range os.Environ() {
		env += fmt.Sprintf("%s\n", e)
	}
	return &Response{
		Body: fmt.Sprintf("Hello %s\n%s", in.Name, env),
	}, nil
}
