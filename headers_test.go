package fuzzyHelpers

import (
	"testing"
)

func TestChrome(t *testing.T) {
	t.Parallel()
	h := NewHeaders()

	t.Run("no headers at initialization", func(t *testing.T) {
		want := 0
		got := len(h.headerMap)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("chrome() provides correct number of chrome headers", func(t *testing.T) {
		want := 13
		h.chrome()
		got := len(h.headerMap)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestFirefox(t *testing.T) {
	t.Parallel()
	h := NewHeaders()

	t.Run("no headers at initialization", func(t *testing.T) {
		want := 0
		got := len(h.headerMap)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("firefox() provides correct number of headers", func(t *testing.T) {
		want := 11
		h.firefox()
		got := len(h.headerMap)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestWithURL(t *testing.T) {
	t.Parallel()
	h := NewHeaders(
		WithURL("https://example.com/foo"),
	)
	foo := h.Headers()
	want := "example.com"
	got := foo["Host"][0]
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestBadWithURLInput(t *testing.T) {
	t.Parallel()
	h := NewHeaders(
		WithURL("kl klsajdf; jkl"),
	)
	foo := h.Headers()
	if v, ok := foo["Host"]; ok {
		t.Errorf("wanted no host header, got %v", v)
	}
}

func TestWithOSInHeader(t *testing.T) {
	t.Parallel()
	h := NewHeaders(
		WithOS("m"),
	)
	// check chrome for sec-ch-ua-platform header
	h.chrome()
	want := "Macintosh"
	got := h.headerMap["sec-ch-ua-platform"][0]
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestWithOSDefaultWindows(t *testing.T) {
	t.Parallel()
	h := NewHeaders(
		WithOS("foo"),
	)
	// check chrome for sec-ch-ua-platform header
	h.chrome()
	want := "Windows"
	got := h.headerMap["sec-ch-ua-platform"][0]
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestWithOSAny(t *testing.T) {
	t.Parallel()
	h := NewHeaders(
		WithOS("any"),
	)
	if h.osys != "l" && h.osys != "m" && h.osys != "w" {
		t.Errorf("wanted l, m, or w, got %s", h.osys)
	}
}

func TestCaseInsensitiveOS(t *testing.T) {
	t.Parallel()
	h := NewHeaders(
		WithOS("M"),
	)
	if h.osys != "m" {
		t.Errorf("got %s wanted 'm'", h.osys)
	}
}

func TestBadOSInputUsesDefault(t *testing.T) {
	t.Parallel()
	h := NewHeaders(
		WithOS("foo"),
	)
	if h.osys != "w" {
		t.Errorf("got %s wanted 'w'", h.osys)
	}
}

func TestWithOSInUA(t *testing.T) {
	t.Parallel()
	h := NewHeaders(
		WithOS("m"),
	)
	h.chrome()
	ua := []string{
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
	if !assertCorrectUA(t, h.headerMap["User-Agent"][0], ua) {
		t.Errorf("wanted %s to be a Macintosh ua", h.headerMap["User-Agent"][0])
	}
}

func assertCorrectUA(t *testing.T, ua string, possible []string) bool {
	t.Helper()
	for _, v := range possible {
		if ua == v {
			return true
		}
	}
	return false
}

func TestHeaders(t *testing.T) {
	t.Parallel()
	headers := Headers()
	t.Run("got headers", func(t *testing.T) {
		if len(headers) == 0 {
			t.Errorf("want headers, got none")
		}
	})
	t.Run("no host set", func(t *testing.T) {
		if v, ok := headers["Host"]; ok {
			t.Errorf("wanted no Host header, got %v", v)
		}
	})
	t.Run("correct number of headers", func(t *testing.T) {
		if len(headers) != 11 && len(headers) != 13 {
			t.Errorf("number of headers was %d, wanted 11 or 13", len(headers))
		}
	})
}

func TestCustomHeaders(t *testing.T) {
	t.Parallel()
	h := NewHeaders(
		WithCustomHeaders("Host=example.com User-Agent=foobar"),
	)
	headers := h.Headers()
	t.Run("got headers", func(t *testing.T) {
		if len(headers) == 0 {
			t.Errorf("wanted headers, got none")
		}
	})
	t.Run("set custom host", func(t *testing.T) {
		if headers["Host"][0] != "example.com" {
			t.Errorf("got %s wanted %q", headers["Host"][0], "example.com")
		}
	})
	t.Run("custom user agent isn't overwritten", func(t *testing.T) {
		if headers["User-Agent"][0] != "foobar" {
			t.Errorf("got %s wanted %q", headers["User-Agent"][0], "foobar")
		}
	})
	t.Run("correct number of headers", func(t *testing.T) {
		// 11 or 13 as default plus the Host header
		if len(headers) != 12 && len(headers) != 14 {
			t.Errorf("number of headers was %d, wanted 12 or 14", len(headers))
		}
	})
}

func TestSuppressHeaders(t *testing.T) {
	t.Parallel()
	h := NewHeaders(
		SuppressHeaders("User-Agent"),
	)
	headers := h.Headers()
	if v, ok := headers["User-Agent"]; ok {
		t.Errorf("got %s wanted no header", v)
	}
}

func TestSuppressMultipleHeaders(t *testing.T) {
	t.Parallel()
	h := NewHeaders(
		SuppressHeaders("User-Agent Accept"),
	)
	headers := h.Headers()
	if v, ok := headers["User-Agent"]; ok {
		t.Errorf("got %s wanted no User-Agent header", v)
	}
	if v, ok := headers["Accept"]; ok {
		t.Errorf("got %s wanted no Accept header", v)
	}
}

func TestChromeOnly(t *testing.T) {
	t.Parallel()
	h := NewHeaders(
		ChromeOnly(true),
	)
	headers := h.Headers()
	// ff doesn't have this header
	if _, ok := headers["sec-ch-ua-platform"]; !ok {
		t.Error("wanted 'sec-ch-ua-platform' header but got none")
	}
}

func TestFirefoxOnly(t *testing.T) {
	t.Parallel()
	h := NewHeaders(
		FirefoxOnly(true),
	)
	headers := h.Headers()
	if v, ok := headers["sec-ch-ua-platform"]; ok {
		t.Errorf("got %v but wanted no 'sec-ch-ua-platform' header", v)
	}
}

func TestChromePriority(t *testing.T) {
	t.Parallel()
	h := NewHeaders(
		FirefoxOnly(true),
		ChromeOnly(true),
	)
	headers := h.Headers()
	if _, ok := headers["sec-ch-ua-platform"]; !ok {
		t.Error("wanted 'sec-ch-ua-platform' header but got none")
	}
}
