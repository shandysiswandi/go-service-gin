package jsonplaceholder

import (
	"encoding/json"
	"net/http"
	"time"
)

// Posts is
type Posts struct {
	endPoint string
	client   *http.Client
}

// NewPosts is
func NewPosts(baseURL string) *Posts {
	return &Posts{baseURL + "/posts", &http.Client{
		Timeout: time.Minute / 2,
	}}
}

// Fetch is
func (p *Posts) Fetch() ([]Post, error) {
	var posts []Post

	req, err := http.NewRequest(http.MethodGet, p.endPoint, nil)
	if err != nil {
		return nil, err
	}

	response, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		return nil, err
	}

	return posts, nil
}
