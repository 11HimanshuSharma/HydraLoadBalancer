package backend

import (
	"net/http/httputil"
	"net/url"
	"sync/atomic"
	"time"
)

type Backend struct {

	// immutable fields
	URL *url.URL
	Weight int

	// request routing
	Proxy *httputil.ReverseProxy

	// mutable state
	alive atomic.Bool
	activeConnections atomic.Int64
	failureCount atomic.Int64

	lastHealthCheck stomic.Int64

}

func NewBackend(rawURL string, weight int, proxy *httputil.ReverseProxy) (*Backend, error) {

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	b := &Backend {
		URL: parsedURL,
		Weight: weight,
		Proxy: proxy,
	}

	b.alive.Store(true)

	b.lastHealthCheck.Store(
		time.Now().Unix(),
	)

	return b, nil

}