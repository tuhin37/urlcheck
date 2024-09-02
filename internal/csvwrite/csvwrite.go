package csvwrite

import (
	"fmt"
	"os"
	"sync"

	"github.com/tuhin37/urlcheck/internal/models"
)

func WriteCSV(urlOrder []string, data map[string]*models.URLData, filename string, numWorkers int) error {
	// Create the output file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Convert map to slice for easier partitioning
	var entries []string
	for _, url := range urlOrder {
		urlData, ok := data[url]
		if !ok {
			continue
		}
		entry := fmt.Sprintf("<%s>,%s\n", urlData.URL, urlData.Result)
		entries = append(entries, entry)
	}

	// Use a single goroutine to write entries in order
	var wg sync.WaitGroup
	var mu sync.Mutex

	writeEntries := func(startIndex int) {
		defer wg.Done()
		for i := startIndex; i < len(entries); i++ {
			mu.Lock()
			_, err := file.WriteString(entries[i])
			mu.Unlock()
			if err != nil {
				fmt.Printf("Failed to write to file: %v\n", err)
			}
		}
	}

	// Start a single worker to write all entries in order
	wg.Add(1)
	go writeEntries(0)

	// Wait for the worker to finish
	wg.Wait()

	return nil
}
