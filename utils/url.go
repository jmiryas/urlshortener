package utils

import (
	"net/url"
	"strings"
)

// IsValidURL mem-validasi sebuah URL.
// Jika urlString tidak memiliki scheme, fungsi ini akan mencoba menambahkan "http://" untuk validasi.
func IsValidURL(urlString string) bool {
	urlString = strings.TrimSpace(urlString)
	if urlString == "" {
		return false
	}

	// Jika user mengirim "example.com" tanpa scheme, tambahkan http:// untuk validasi
	if !strings.Contains(urlString, "://") {
		urlString = "http://" + urlString
	}

	u, err := url.Parse(urlString)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	// Hanya terima http/https
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	// Validasi TLD minimal 2 karakter
	parts := strings.Split(u.Host, ".")
	if len(parts) < 2 {
		return false
	}
	tld := parts[len(parts)-1]
	
	return len(tld) >= 2

	// if len(tld) < 2 {
	// 	return false
	// }

	// return true
}

// Optional: helper untuk men-normalize url (mengembalikan url dengan scheme)
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
