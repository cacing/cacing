package utils

import (
	"net/url"
)

// ParseURL func
func ParseURL(rawURL string) (*url.URL, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}
