package common_utils

import (
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type Options struct {
	Headers map[string]string
	Proxy   string
}

func Fetch(urlStr string, options Options) (*goquery.Document, error) {
	// Create a new request using http
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}

	// Add headers to the request
	for key, value := range options.Headers {
		req.Header.Add(key, value)
	}

	var transport *http.Transport

	if options.Proxy != "" {
		proxyURL, err := url.Parse(options.Proxy)
		if err != nil {
			return nil, err
		}

		// Create a custom Transport that uses the proxy
		transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	} else {
		transport = &http.Transport{}
	}

	// Create a custom http.Client with the Transport
	client := &http.Client{
		Transport: transport,
	}
	// Send the request via a client
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Defer the closing of the body
	defer res.Body.Close()

	data, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
