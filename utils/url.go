package utils

import (
	"net/url"
	"strings"
)

func IsValidURL(urlString string) bool {
	urlString = strings.TrimSpace(urlString)

	if urlString == "" {
		return false
	}

	if !strings.Contains(urlString, "://") {
		urlString = "http://" + urlString
	}

	u, err := url.Parse(urlString)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	parts := strings.Split(u.Host, ".")
	if len(parts) < 2 {
		return false
	}
	tld := parts[len(parts)-1]
	
	return len(tld) >= 2
}

func NormalizeURL(urlString string) string {
	urlString = strings.TrimSpace(urlString)

	if urlString == "" {
		return urlString
	}
	
	if !strings.Contains(urlString, "://") {
		return "https://" + urlString
	}
	
	return urlString
}
