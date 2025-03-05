package save

import (
	resp "url-shortener/internal/lib/api/response"
)

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}
