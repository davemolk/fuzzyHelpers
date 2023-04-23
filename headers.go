package fuzzyHelpers

import (
	"math/rand"
	"net/http"
	"net/url"
)

type headers struct {
	osys      string
	HeaderMap map[string][]string
}

type option func(*headers) error

func NewHeaders(opts ...option) (*headers, error) {
	h := &headers{
		osys:      "w",
		HeaderMap: make(map[string][]string),
	}
	for _, opt := range opts {
		err := opt(h)
		if err != nil {
			return &headers{}, err
		}
	}
	return h, nil
}

func WithOS(osys string) option {
	return func(h *headers) error {
		switch osys {
		case "l", "m", "w":
			h.osys = osys
		default:
			// prob get rid of error in option if i'm
			// just going to enter defaults on errors...
			h.osys = "w"
		}
		return nil
	}
}

func WithURL(s string) option {
	return func(h *headers) error {
		u, err := url.Parse(s)
		if err != nil {
			// don't set Host on headers
			return nil
		}
		h.HeaderMap["Host"] = []string{u.Host}
		return nil
	}
}

func Headers() map[string][]string {
	h, err := NewHeaders()
	if err != nil {
		panic(err)
	}
	if rand.Intn(2) == 1 {
		h.Firefox()
	} else {
		h.Chrome()
	}
	return h.HeaderMap
}

func (h *headers) Headers() map[string][]string {
	if rand.Intn(2) == 1 {
		h.Firefox()
	} else {
		h.Chrome()
	}
	return h.HeaderMap
}

func (h *headers) Firefox() {
	uAgent := h.ffUA()
	h.HeaderMap = http.Header{
		"User-Agent":                {uAgent},
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"},
		"Accept-Language":           {"en-US,en;q=0.5"},
		"DNT":                       {"1"},
		"Connection":                {"keep-alive"},
		"Upgrade-Insecure-Requests": {"1"},
		"Sec-Fetch-Dest":            {"document"},
		"Sec-Fetch-Mode":            {"navigate"},
		"Sec-Fetch-Site":            {"none"},
		"Sec-Fetch-User":            {"?1"},
		"Sec-GCP":                   {"1"},
	}
}

func (h *headers) Chrome() {
	uAgent := h.chromeUA()
	h.HeaderMap = http.Header{
		"Connection":                {"keep-alive"},
		"Cache-Control":             {"max-age=0"},
		"sec-ch-ua":                 {`" Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99"`},
		"sec-ch-ua-mobile":          {"?0"},
		"Upgrade-Insecure-Requests": {"1"},
		"User-Agent":                {uAgent},
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"},
		"Sec-Fetch-Site":            {"none"},
		"Sec-Fetch-Mode":            {"navigate"},
		"Sec-Fetch-User":            {"?1"},
		"Sec-Fetch-Dest":            {"document"},
		"Accept-Language":           {"en-US,en;q=0.5"},
	}

	switch h.osys {
	case "m":
		h.HeaderMap["sec-ch-ua-platform"] = []string{"Macintosh"}
	case "l":
		h.HeaderMap["sec-ch-ua-platform"] = []string{"Linux"}
	default:
		h.HeaderMap["sec-ch-ua-platform"] = []string{"Windows"}
	}
}

func (h *headers) ffUA() string {
	var userAgents []string
	switch h.osys {
	case "m":
		userAgents = []string{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:107.0) Gecko/20100101 Firefox/107.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:106.0) Gecko/20100101 Firefox/106.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:104.0) Gecko/20100101 Firefox/104.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:102.0) Gecko/20100101 Firefox/102.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:101.0) Gecko/20100101 Firefox/101.0",
		}
	case "l":
		userAgents = []string{
			"Mozilla/5.0 (X11; Linux x86_64; rv:93.0) Gecko/20100101 Firefox/93.0 ",
			"Mozilla/5.0 (X11; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0 ",
			"Mozilla/5.0 (X11; Linux x86_64; rv:104.0) Gecko/20100101 Firefox/104.0 ",
			"Mozilla/5.0 (X11; Linux x86_64; rv:106.0) Gecko/20100101 Firefox/106.0 ",
			"Mozilla/5.0 (X11; Linux x86_64; rv:107.0) Gecko/20100101 Firefox/107.0 ",
		}
	default:
		userAgents = []string{
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:99.0) Gecko/20100101 Firefox/99.0",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:108.0) Gecko/20100101 Firefox/108.0",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:106.0) Gecko/20100101 Firefox/106.0",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:104.0) Gecko/20100101 Firefox/104.0",
		}
	}
	random := rand.Intn(len(userAgents))
	return userAgents[random]
}

func (h *headers) chromeUA() string {
	var userAgents []string
	switch h.osys {
	case "m":
		userAgents = []string{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4692.56 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4889.0 Safari/537.36",
		}
	case "l":
		userAgents = []string{
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4692.56 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4889.0 Safari/537.36",
		}
	default:
		userAgents = []string{
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36",
			"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36",
			"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36",
		}
	}
	random := rand.Intn(len(userAgents))
	return userAgents[random]
}
