package modules

import (
	"net"
	"net/url"
	"strconv"
)

func UrlChecker(target string) (string, string, bool) {
	Schema, err := url.ParseRequestURI(target)
	if err != nil {
		return "", "", false
	}
	return Schema.Scheme + "://" + Schema.Host, Schema.Hostname(), true
}
func IPChecker(ip string) bool {
	address := net.ParseIP(ip)
	if address == nil {
		return false
	} else {
		return true
	}
}
func PortChecker(port string) bool {
	p, err := strconv.Atoi(port)
	if err != nil {
		return false
	}
	if p <= 0 || p >= 65535 {
		return false
	}
	return true
}
