# fuzzyHelpers
* provides headers that mimic chrome and firefox 
* provides a client with helpful defaults for fuzzing a site

# basic usage
```
url := "https://example.com"
req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
if err != nil {
    return err
}
req.Header = fuzzyHelpers.Headers()
resp, err := fuzzyHelpers.Client().Do(req)
```

# using options
```
url := "https://example.com"
req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
if err != nil {
    return err
}

// pass in options
h := fuzzyHelpers.NewHeaders(
    // will generate a linux, mac, or windows ua and
    // matching "sec-ch-ua-platform" header if returning
    // a set of chrome headers
    fuzzyHelpers.WithOS("any"),
    // will include a "Host" header
    fuzzyHelpers.WithURL(url),
    // include custom header(s)
    // format: space-separated key=value
    // this gives the famous 'foo' and 'go' headers 
    // with values 'bar' and 'pher', respectively
    fuzzyHelpers.WithCustomHeaders("foo=bar go=pher"),
)

// call Headers to generate
req.Header = h.Headers()

c := fuzzyHelpers.NewClient(
    // maybe we want to send through burp suite, for instance
    fuzzyHelpers.WithProxy("http://127.0.0.1:8080"),
    // set a timeout via client if you're not using
    // context.WithTimeout
    fuzzyHelpers.WithTimeout(15000),
)
resp, err := c.Do(req)
etc...
```
### headers overview
fuzzyHelpers selects randomly between firefox headers and chrome headers. use WithOS to provide choose an appropriate (random) user agent and sec-ch-ua-platform header (if chrome is selected). 
```
firefox
    User-Agent = a random ua
    Accept = text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8
    Accept-Language: en-US,en;q=0.5
    DNT = 1
    Connection = keep-alive
    Upgrade-Insecure-Requests = 1
    Sec-Fetch-Dest = document
    Sec-Fetch-Mode = navigate
    Sec-Fetch-Site = none
    Sec-Fetch-User = ?1
    Sec-GCP = 1

chrome
    Connection = keep-alive
    Cache-Control = max-age=0
    sec-ch-ua = "Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99"
    sec-ch-ua-mobile = ?0
    sec-ch-ua-platform = Linux, Macintosh, or Windows, depending on your input
    Upgrade-Insecure-Requests = 1
    User-Agent = a random ua
    Accept = text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8
    Sec-Fetch-Site = none
    Sec-Fetch-Mode = navigate
    Sec-Fetch-User = ?1
    Sec-Fetch-Dest = document
    Accept-Language = en-US,en;q=0.5
```
### client overview
```
fuzzyHelpers provides a client with the following customized defaults
    MaxIdleConnsPerHost = 30
    MaxConnsPerHost = 30
    InsecureSkipVerify = true
    Timeout = 10000*time.Millisecond
    CheckRedirect = func(req *http.Request, via []*http.Request) error {
                        return http.ErrUseLastResponse
                    }
```
### user-supplied options
```
header options:
  WithOS
    	used in "sec-ch-ua-platform" chrome header
        possible values are "l, m, w, or any"
        "any" will select randomly between l, m, and w
        default value is "w" 
  WithCustomHeaders
        include custom header(s) as space-separated key=value
        e.g. WithCustomHeaders("foo=bar go=pher")
        giving 'foo' and 'go' headers the values 'bar' and 'pher'.
        note: custom headers are not overwritten by default values
  SuppressHeaders
        include space-separated header(s) to suppress from request
  WithURL
        include the request url for fuzzyHelper to set the Host header
  ChromeOnly
        only use chrome headers
  FirefoxOnly
        only use firefox headers

client options
  WithConnections
    	sets MaxIdleConnsPerHost and MaxConnsPerHost
        try entering your number of concurrent requests
  WithNoSkip
    	pass in true if you want InsecureSkipVerify = false
  WithProxy
    	pass in a proxy
  WithAllowedRedirects
    	pass in true if you want to allow redirects
  WithTimeout
        measured in ms
```
### note
Go unfortunately doesn't preserve header order, so if that's important to you and what you're up to, you'll need to look elsewhere. 