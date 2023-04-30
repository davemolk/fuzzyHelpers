package fuzzyHelpers

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

type clientOptions struct {
	allowRedirects bool
	connections    int
	noSkip         bool
	proxy          string
	timeout        int
}

type optionClient func(*clientOptions)

func NewClient(opts ...optionClient) *http.Client {
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.MaxIdleConnsPerHost = 30
	tr.MaxConnsPerHost = 30
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout:   5000 * time.Millisecond,
		Transport: tr,
	}
	c := &clientOptions{}
	for _, opt := range opts {
		opt(c)
	}

	// populate any options
	if c.connections != 0 {
		tr.MaxIdleConnsPerHost = c.connections
		tr.MaxConnsPerHost = c.connections
	}
	if c.noSkip {
		tr.TLSClientConfig.InsecureSkipVerify = false
	}
	if c.proxy != "" {
		parsed, err := url.Parse(c.proxy)
		// if error, won't set a proxy
		if err == nil {
			tr.Proxy = http.ProxyURL(parsed)
		}
	}
	if c.allowRedirects {
		client.CheckRedirect = nil
	}
	if c.timeout > 0 {
		client.Timeout = time.Duration(c.timeout) * time.Millisecond
	}
	return client
}

func WithConnections(n int) optionClient {
	return func(c *clientOptions) {
		if n <= 0 {
			return
		}
		c.connections = n
	}
}

func WithNoSkip(v bool) optionClient {
	return func(c *clientOptions) {
		c.noSkip = true
	}
}

func WithProxy(p string) optionClient {
	return func(c *clientOptions) {
		if p == "" {
			return
		}
		c.proxy = p
	}
}

func WithAllowRedirects(r bool) optionClient {
	return func(c *clientOptions) {
		c.allowRedirects = r
	}
}

func WithTimeout(t int) optionClient {
	return func(c *clientOptions) {
		if t <= 0 {
			return
		}
		c.timeout = t
	}
}

func Client() *http.Client {
	return NewClient()
}
