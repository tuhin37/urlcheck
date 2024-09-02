package print

import (
	"fmt"
	"strings"

	"github.com/tuhin37/urlcheck/internal/models"
)

// PrintResults prints the results in either CSV or table format.
func PrintResults(urlOrder []string, data map[string]*models.URLData, format string) {
	if strings.ToUpper(format) == "TABLE" {
		printTable(urlOrder, data)
	} else {
		printCSV(urlOrder, data)
	}
}

// printCSV prints the results in the specified CSV format.
func printCSV(urlOrder []string, data map[string]*models.URLData) {
	// No header, only two columns: URL and Status
	for _, url := range urlOrder {
		urlData := data[url]
		fmt.Printf("<%s>, %s\n", urlData.URL, urlData.Result)
	}
}

// printTable prints the results in a tabular format.
func printTable(urlOrder []string, data map[string]*models.URLData) {
	// Determine column widths
	urlWidth := len("URL")
	timeoutWidth := len("Timeout(ms)")
	statusWidth := len("Status")

	for _, url := range urlOrder {
		urlData := data[url]
		if len(url) > urlWidth {
			urlWidth = len(url)
		}
		if len(fmt.Sprintf("%d", urlData.TimeoutMs)) > timeoutWidth {
			timeoutWidth = len(fmt.Sprintf("%d", urlData.TimeoutMs))
		}
		if len(urlData.Result) > statusWidth {
			statusWidth = len(urlData.Result)
		}
	}

	// Print table header
	fmt.Printf("%-*s  %-*s  %-*s\n", urlWidth, "URL", timeoutWidth, "Timeout(ms)", statusWidth, "Status")
	fmt.Printf("%s  %s  %s\n", strings.Repeat("-", urlWidth), strings.Repeat("-", timeoutWidth), strings.Repeat("-", statusWidth))

	// Print table rows
	for _, url := range urlOrder {
		urlData := data[url]
		fmt.Printf("%-*s  %-*d  %-*s\n", urlWidth, urlData.URL, timeoutWidth, urlData.TimeoutMs, statusWidth, urlData.Result)
	}
}
