package fetcher

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
)

var ErrSessionCookieRequired = errors.New("a session cookie is required when input has not yet been cached locally")

// Fetcher retrieves the input for a problem.
type Fetcher interface {
	Fetch() (io.Reader, error)
}

// Caching fetcher fetches and caches input.
type CachingFetcher struct {
	// cookie is the session cookie used to retrieve input.
	cookie string
	// path is the path to the input cached on disk.
	path string
	// url is the URL of the input sourced from the network.
	url *url.URL
}

func (cf CachingFetcher) Fetch() (io.ReadSeeker, error) {
	contents, err := os.ReadFile(cf.path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}

		req, err := http.NewRequest(http.MethodGet, cf.url.String(), nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("cookie", cf.cookie)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		contents, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if err := os.WriteFile(cf.path, contents, os.ModePerm); err != nil {
			return nil, err
		}
	}

	f, err := os.Open(cf.path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// NewCachingFetcher returns a CachingFetcher that retrieves input from the URL
// specified in source and stores input in the file specified by the path local.
func NewCachingFetcher(source string, cookie string, local string) (CachingFetcher, error) {
	url, err := url.Parse(source)
	if err != nil {
		return CachingFetcher{}, err
	}

	if _, err := os.Stat(local); err != nil {
		if errors.Is(err, os.ErrNotExist) && cookie == "" {
			return CachingFetcher{}, ErrSessionCookieRequired
		}
	}

	return CachingFetcher{
		cookie: cookie,
		path:   local,
		url:    url,
	}, nil
}
