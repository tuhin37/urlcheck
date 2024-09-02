package csvread

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/tuhin37/urlcheck/internal/models"
)

type URLData struct {
	URL       string
	TimeoutMs int
	Result    string
}

func ReadCSV(filePath string, offset, limit int) ([]string, map[string]*models.URLData) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	if len(records) <= 1 {
		panic("CSV file has no data rows")
	}

	// Skip the header (first row)
	records = records[1:]

	// Apply offset and limit
	if offset > 0 {
		if offset >= len(records) {
			return nil, make(map[string]*models.URLData) // Offset is beyond the length of records
		}
		records = records[offset:]
	}
	if limit > 0 && limit < len(records) {
		records = records[:limit]
	}

	urlOrder := make([]string, 0, len(records))
	data := make(map[string]*models.URLData)

	// Process records sequentially with a single worker
	for _, record := range records {
		if len(record) < 2 {
			continue
		}

		url := record[0]
		timeoutMs, err := strconv.Atoi(record[1])
		if err != nil {
			timeoutMs = 1000 // Default timeout if conversion fails
		}

		// Create URLData and set initial status as "pending"
		urlData := &models.URLData{
			URL:       url,
			TimeoutMs: timeoutMs,
			Result:    "pending",
		}

		data[url] = urlData
		urlOrder = append(urlOrder, url) // Maintain the order
	}

	return urlOrder, data
}
