package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
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

func Main(in Request) (*Response, error) {
	if in.List == nil {
		namesURL := fmt.Sprintf("https://urchin-app-lcj2a.ondigitalocean.app/%s", "team/team-names")
		res, err := http.Get(namesURL)

		if err != nil {
			return &Response{
				StatusCode: res.StatusCode,
			}, err
		}

		var req Request
		err = json.NewDecoder(res.Body).Decode(&req.List)
		if err != nil {
			return &Response{
					StatusCode: http.StatusFailedDependency},
				err
		}

		in.List = req.List

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
