package main

import (
	"encoding/json"
	"math/rand"
	"time"
)

type Request struct {
	List []string `json:"list"`
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
	if in.List == nil {
		in.List = people
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(in.List), func(i, j int) {
		in.List[i], in.List[j] = in.List[j], in.List[i]
	})

	stringBody, _ := json.Marshal(in.List)

	return &Response{
		Body: string(stringBody),
	}, nil
}
