package cmdparse

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	InputFile    string
	OutputFormat string
	NumWorkers   int
	Verbose      bool
	Silent       bool
	CsvOutFile   string
	Offset       int
	Limit        int
}

func ParseArguments() Config {
	var config Config

	// Define flags
	flag.StringVar(&config.OutputFormat, "output-format", "csv", "Output format: csv (default) or table")
	flag.IntVar(&config.NumWorkers, "workers", 1, "Number of workers for parallel processing")
	flag.BoolVar(&config.Verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&config.Silent, "silent", false, "Suppress all output")
	flag.StringVar(&config.CsvOutFile, "csv-out", "", "Write output to a CSV file")
	flag.IntVar(&config.Offset, "offset", 0, "Offset for reading records (default: 0)")
	flag.IntVar(&config.Limit, "limit", 0, "Limit for reading records (default: 0)")

	// Parse the flags
	flag.Parse()

	// Handle positional argument (input file)
	if len(flag.Args()) != 1 {
		fmt.Println("Error: Exactly one input file must be provided.")
		printHelp()
		os.Exit(1)
	}
	config.InputFile = flag.Args()[0]

	// Validate the command
	if config.OutputFormat != "csv" && config.OutputFormat != "table" {
		fmt.Println("Error: Invalid output format. Allowed values are 'csv' or 'table'.")
		printHelp()
		os.Exit(1)
	}

	if config.Silent && config.Verbose {
		fmt.Println("Error: --silent and --verbose cannot be used together.")
		printHelp()
		os.Exit(1)
	}

	return config
}

// PrintConfig prints the parsed configuration details
func PrintConfig(config Config) {
	fmt.Println("=================== command parsed ===================")
	fmt.Println("Parsed arguments:")
	fmt.Printf("  Input-file:      %s\n", config.InputFile)
	fmt.Printf("  Read-offset:     %d\n", config.Offset)
	fmt.Printf("  Read-limit:      %d\n", config.Limit)
	fmt.Printf("  Verbose:         %t\n", config.Verbose)
	fmt.Printf("  Silent:          %t\n", config.Silent)
	fmt.Printf("  Workers:         %d\n", config.NumWorkers)
	fmt.Printf("  Output-format:   %s\n", config.OutputFormat)
	fmt.Printf("  Output-file:     %s\n", config.CsvOutFile)
}

// printHelp prints a basic help message when the user provides incorrect arguments
func printHelp() {
	// Define the header and option lines
	header := "\nUsage: urlcheck [OPTIONS] file.csv"
	options := []struct {
		option string
		desc   string
	}{
		{"--output-format csv/table", "Print format (default: csv)"},
		{"--workers n", "Number of workers for URL checking (default: 1)"},
		{"--verbose", "Enable verbose output (default: disabled)"},
		{"--silent", "Suppress all output (default: disabled)"},
		{"--write-csv filename.csv", "Write output to a CSV file (default: none)"},
		{"--read-offset n", "Offset for reading records (default: 0 no-offset)"},
		{"--read-limit n", "Number of records from the Offset (default: 0 no-limit)"},
		{"-h, --help", "Display this help message"},
	}

	// Calculate the maximum width for options column
	maxOptionWidth := 0
	for _, opt := range options {
		if len(opt.option) > maxOptionWidth {
			maxOptionWidth = len(opt.option)
		}
	}

	// Print the header
	fmt.Println(header)
	fmt.Println("Options:")

	// Print each option with aligned columns
	for _, opt := range options {
		fmt.Printf("  %-*s %s\n", maxOptionWidth, opt.option, opt.desc)
	}
}
