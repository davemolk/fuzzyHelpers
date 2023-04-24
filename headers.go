package fuzzyHelpers

import (
	"math/rand"
	"net/url"
	"time"
)

type headers struct {
	osys      string
	headerMap map[string][]string
}

type option func(*headers) error

func init() {
	rand.Seed(time.Now().Unix())
}

func NewHeaders(opts ...option) (*headers, error) {
	h := &headers{
		osys:      "w",
		headerMap: make(map[string][]string),
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
		u, err := url.ParseRequestURI(s)
		if err != nil {
			// don't set Host on headers
			return nil
		}
		h.headerMap["Host"] = []string{u.Host}
		return nil
	}
}

func (h *headers) Headers() map[string][]string {
	if rand.Intn(2) == 1 {
		h.firefox()
	} else {
		h.chrome()
	}
	return h.headerMap
}

func Headers() map[string][]string {
	h, err := NewHeaders()
	if err != nil {
		panic(err)
	}
	return h.Headers()
}

func (h *headers) firefox() {
	uAgent := h.ffUA()
	h.headerMap["User-Agent"] = []string{uAgent}
	h.headerMap["Accept"] = []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"}
	h.headerMap["Accept-Language"] = []string{"en-US,en;q=0.5"}
	h.headerMap["DNT"] = []string{"1"}
	h.headerMap["Connection"] = []string{"keep-alive"}
	h.headerMap["Upgrade-Insecure-Requests"] = []string{"1"}
	h.headerMap["Sec-Fetch-Dest"] = []string{"document"}
	h.headerMap["Sec-Fetch-Mode"] = []string{"navigate"}
	h.headerMap["Sec-Fetch-Site"] = []string{"none"}
	h.headerMap["Sec-Fetch-User"] = []string{"?1"}
	h.headerMap["Sec-GCP"] = []string{"1"}
}

func (h *headers) chrome() {
	uAgent := h.chromeUA()
	h.headerMap["Connection"] = []string{"keep-alive"}
	h.headerMap["Cache-Control"] = []string{"max-age=0"}
	h.headerMap["sec-ch-ua"] = []string{`"Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99"`}
	h.headerMap["sec-ch-ua-mobile"] = []string{"?0"}
	h.headerMap["Upgrade-Insecure-Requests"] = []string{"1"}
	h.headerMap["User-Agent"] = []string{uAgent}
	h.headerMap["Accept"] = []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"}
	h.headerMap["Sec-Fetch-Site"] = []string{"none"}
	h.headerMap["Sec-Fetch-Mode"] = []string{"navigate"}
	h.headerMap["Sec-Fetch-User"] = []string{"?1"}
	h.headerMap["Sec-Fetch-Dest"] = []string{"document"}
	h.headerMap["Accept-Language"] = []string{"en-US,en;q=0.5"}

	switch h.osys {
	case "m":
		h.headerMap["sec-ch-ua-platform"] = []string{"Macintosh"}
	case "l":
		h.headerMap["sec-ch-ua-platform"] = []string{"Linux"}
	default:
		h.headerMap["sec-ch-ua-platform"] = []string{"Windows"}
	}
}

func (h *headers) ffUA() string {
	var userAgents []string
	switch h.osys {
	case "m":
		userAgents = []string{
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 13.13.1; rv:110.0) Gecko/20100101 Firefox/110.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 13.13; rv:108.0) Gecko/20100101 Firefox/108.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 13.13; rv:106.0) Gecko/20100101 Firefox/106.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 11.7.6; rv:108.0) Gecko/20100101 Firefox/108.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 11.1; rv:108.0) Gecko/20100101 Firefox/108.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 11.1; rv:110.0) Gecko/20100101 Firefox/110.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15.7; rv:106.0) Gecko/20100101 Firefox/106.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 12.6.5; rv:110.0) Gecko/20100101 Firefox/110.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 12.1; rv:104.0) Gecko/20100101 Firefox/104.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:102.0) Gecko/20100101 Firefox/102.0",
		}
	case "l":
		userAgents = []string{
			"Mozilla/5.0 (X11; Linux x86_64; rv:110.0) Gecko/20100101 Firefox/110.0",
			"Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/109.0",
			"Mozilla/5.0 (X11; Linux x86_64; rv:108.0) Gecko/20100101 Firefox/108.0",
			"Mozilla/5.0 (X11; Linux x86_64; rv:107.0) Gecko/20100101 Firefox/107.0",
			"Mozilla/5.0 (X11; Linux x86_64; rv:106.0) Gecko/20100101 Firefox/106.0",
			"Mozilla/5.0 (X11; Linux x86_64; rv:105.0) Gecko/20100101 Firefox/105.0",
			"Mozilla/5.0 (X11; Linux x86_64; rv:104.0) Gecko/20100101 Firefox/104.0",
			"Mozilla/5.0 (X11; Linux x86_64; rv:103.0) Gecko/20100101 Firefox/103.0",
			"Mozilla/5.0 (X11; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0",
			"Mozilla/5.0 (X11; Linux x86_64; rv:93.0) Gecko/20100101 Firefox/93.0",
		}
	default:
		userAgents = []string{
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:110.0) Gecko/20100101 Firefox/110.0",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/109.0",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:108.0) Gecko/20100101 Firefox/108.0",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:107.0) Gecko/20100101 Firefox/107.0",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:106.0) Gecko/20100101 Firefox/106.0",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:105.0) Gecko/20100101 Firefox/105.0",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:104.0) Gecko/20100101 Firefox/104.0",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:99.0) Gecko/20100101 Firefox/99.0",
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
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 13_13_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_7_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 12_6_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36",
		}
	case "l":
		userAgents = []string{
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.5615.137 Safari/537.36",
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
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.5615.137 Safari/537.36",
		}
	}
	random := rand.Intn(len(userAgents))
	return userAgents[random]
}
