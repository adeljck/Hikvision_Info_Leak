package modules

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
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
func Decrypt(encrypted string) (string, error) {
	client := resty.New()
	client.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.0.0")
	client.SetHeader("Referer", "https://forum.ywhack.com/decrypt.php")
	client.SetBaseURL("https://forum.ywhack.com/")
	DecryptData := map[string]string{"pwd": encrypted, "type": "Hikvision_iSecure_dbDePass_decrypt"}
	res, err := client.R().SetFormData(DecryptData).Post("decrypt.php")
	if err != nil {
		return "", err
	}
	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(res.Body()))
	if err != nil {
		return "", err
	}
	results := dom.Find(".add_line > p > font")
	if len(results.Nodes) != 1 {
		return "", errors.New("Decrypt Failed.")
	}
	DecryptResults := results.Text()
	return DecryptResults, nil
}
