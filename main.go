package main

import (
	"fmt"
	"os"

	"github.com/tuhin37/urlcheck/internal/cmdparse"
	"github.com/tuhin37/urlcheck/internal/csvread"
	"github.com/tuhin37/urlcheck/internal/csvwrite"
	"github.com/tuhin37/urlcheck/internal/netcheck"
	"github.com/tuhin37/urlcheck/internal/print"
	"github.com/tuhin37/urlcheck/internal/urlcheck"
)

func main() {
	// ---------------------- parse the command ----------------------
	config := cmdparse.ParseArguments()

	// Print configuration if verbose mode is enabled
	if config.Verbose {
		cmdparse.PrintConfig(config)
	}

	// Check internet connecivity
	if !netcheck.CheckConnectivity() {
		fmt.Println("\n\nError: No internet connection. Please check your network settings.")
		os.Exit(1)
	}

	// ---------------------- read csv ----------------------
	// Read CSV data and get ordered URLs
	urlOrder, data := csvread.ReadCSV(config.InputFile, config.Offset, config.Limit)
	if config.Verbose {
		fmt.Println("\n\n=================== input data ===================")
		for _, url := range urlOrder {
			urlData := data[url]
			fmt.Printf("URL: %s, Timeout: %d ms, Status: %s\n", urlData.URL, urlData.TimeoutMs, urlData.Result)
		}
	}

	// ---------------------- check URLs ----------------------
	if config.Verbose {
		fmt.Println("\n\n=================== checking URLs ===================")
	}
	urlcheck.CheckURLs(data, config.NumWorkers)

	// Print results only if not in silent mode
	if !config.Silent {
		print.PrintResults(urlOrder, data, config.OutputFormat)
	}

	// Write CSV only if CsvOutFile is set
	if config.CsvOutFile != "" {
		err := csvwrite.WriteCSV(urlOrder, data, config.CsvOutFile, config.NumWorkers)
		if err != nil {
			fmt.Printf("Error writing CSV: %v\n", err)
		}
	}
}
