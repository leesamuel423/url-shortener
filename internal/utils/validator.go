package utils

import (
	"net/url"
	"strings"
)

func IsValidURL(str string) bool {
	if str == "" {
		return false
	}

	u, err := url.Parse(str)
	if err != nil {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	if u.Host == "" {
		return false
	}

	return true
}

func IsSameDomain(urlStr, domain string) bool {
	u, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	d, err := url.Parse(domain)
	if err != nil {
		return false
	}

	return strings.EqualFold(u.Host, d.Host)
}
