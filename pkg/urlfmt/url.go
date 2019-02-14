package urlfmt

import (
	"fmt"
	"net/url"
	"strings"
)

// Scheme defines which protocol a URL should use.
type Scheme uint8

// SlashHandling defines the way that slashes should be used in the URL.
type SlashHandling uint8

// These are the defined URL schemes.
const (
	HTTPS Scheme = iota
	InsecureHTTP
)

// These are the defined slash-handling modes.
const (
	NoTrailingSlash SlashHandling = iota
	TrailingSlash
)

func (s Scheme) String() string {
	switch s {
	case HTTPS:
		return "https"
	case InsecureHTTP:
		return "http"
	default:
		return fmt.Sprintf("%d", s)
	}
}

// FormatURL takes in an endpoint, whether to prepend https if no scheme is specified and if the url should end in a slash
func FormatURL(endpoint string, defaultScheme Scheme, slash SlashHandling) (string, error) {
	if !strings.HasPrefix(endpoint, "http") {
		endpoint = fmt.Sprintf("%s://%s", defaultScheme, endpoint)
	}
	if slash == TrailingSlash && !strings.HasSuffix(endpoint, "/") {
		return endpoint + "/", nil
	} else if slash == NoTrailingSlash {
		return strings.TrimRight(endpoint, "/"), nil
	}
	return endpoint, nil
}

// FullyQualifiedURL returns a URL in the proper format or returns an error if the format is invalid
func FullyQualifiedURL(endpoint string, values url.Values, args ...string) (string, error) {
	endpoint = strings.TrimRight(endpoint, "/")
	for i, s := range args {
		s = strings.TrimLeft(s, "/")
		s = strings.TrimRight(s, "/")
		args[i] = s
	}
	fullPath := strings.Join(append([]string{endpoint}, args...), "/")
	url, err := url.Parse(fullPath)
	if err != nil {
		return "", err
	}
	url.RawQuery = values.Encode()
	return url.String(), nil
}

// GetServerFromURL takes a url and returns the server and port without a scheme or the rest of the URL path.
// In order for this to parse correctly, the endpoint must contain a scheme
func GetServerFromURL(endpoint string) string {
	u, err := url.Parse(endpoint)
	if err != nil {
		return ""
	}
	return u.Host
}

// TrimHTTPPrefixes cuts off the http prefixes if they exist on the URL
func TrimHTTPPrefixes(url string) string {
	url = strings.TrimPrefix(url, "http://")
	return strings.TrimPrefix(url, "https://")
}
