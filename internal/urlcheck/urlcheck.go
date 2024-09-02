package urlcheck

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/tuhin37/urlcheck/internal/models"
)

// CheckURLs checks the URLs using a fixed number of goroutines.
func CheckURLs(data map[string]*models.URLData, numWorkers int) {
	var wg sync.WaitGroup
	dataChunks := splitMap(data, numWorkers)

	// Launch fixed number of goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(chunk map[string]*models.URLData) {
			defer wg.Done()
			for _, urlData := range chunk {
				checkURL(urlData)
			}
		}(dataChunks[i])
	}

	wg.Wait()
}

// splitMap splits the map into a slice of smaller maps, one for each worker.
func splitMap(data map[string]*models.URLData, numWorkers int) []map[string]*models.URLData {
	// Initialize a slice of maps to hold the data chunks
	dataChunks := make([]map[string]*models.URLData, numWorkers)
	for i := range dataChunks {
		dataChunks[i] = make(map[string]*models.URLData)
	}

	// Distribute the data among the chunks
	i := 0
	for url, urlData := range data {
		dataChunks[i%numWorkers][url] = urlData
		i++
	}

	return dataChunks
}

// checkURL performs the actual check for a single URL and updates its result.
func checkURL(urlData *models.URLData) {
	timeout := time.Duration(urlData.TimeoutMs) * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Handle URL Malformation
	parsedURL, err := url.Parse(urlData.URL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		urlData.Result = "URL_MALFORMED"
		return
	}

	req, err := http.NewRequestWithContext(ctx, "GET", urlData.URL, nil)
	if err != nil {
		urlData.Result = "FAIL_CONNECT"
		return
	}

	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: timeout,
			}).DialContext,
			TLSClientConfig: &tls.Config{
				// Optional: configure SSL settings if needed
			},
		},
	}

	// Resolve DNS
	if _, err := net.LookupHost(parsedURL.Host); err != nil {
		urlData.Result = "FAIL_RESOLVE"
		return
	}

	// Make HTTP request
	resp, err := client.Do(req)
	if err != nil {
		if strings.Contains(err.Error(), "timeout") {
			urlData.Result = "FAIL_TIMEOUT"
		} else if strings.Contains(err.Error(), "redirect") {
			urlData.Result = "TOO_MANY_REDIRECTS"
		} else if strings.Contains(err.Error(), "certificate") || strings.Contains(err.Error(), "SSL") {
			urlData.Result = "SSL_ERROR"
		} else {
			urlData.Result = "FAIL_CONNECT"
		}
		return
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		urlData.Result = "OK"
	} else {
		urlData.Result = "FAIL_RESPONSE"
	}
}

// getHost extracts the host part from a URL.
// func getHost(rawURL string) string {
// 	parsedURL, err := url.Parse(rawURL)
// 	if err != nil {
// 		return rawURL // Fallback: return the original URL if parsing fails
// 	}
// 	return parsedURL.Host
// }
