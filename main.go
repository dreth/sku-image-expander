package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/xuri/excelize/v2"
)

func main() {
	filePath := filepath.Join("files", "images.xlsx")
	f, err := excelize.OpenFile(filePath)
	handleErr(err, true, fmt.Sprintf("Could not open file: %v", err))

	// Get all the rows in the sheet data
	rows, err := f.GetRows("data")
	handleErr(err, true, fmt.Sprintf("The sheet with the table contents must be called 'data': %v", err))

	// Create output directory if it doesn't exist
	createDirIfNotExists("output")

	// Collect SKUs without filename
	swf := []string{}

	// Create file to log the process
	logFile, err := os.Create("logfile.txt") // It's good practice to use a .txt extension
	handleErr(err, true, fmt.Sprintf("Could not create log file: %v", err))
	defer logFile.Close() // Ensure the log file is closed when the program exits

	// Iterate over rows
	for idx, row := range rows {
		if idx > 0 {
			// Get SKU (should always be present)
			sku := row[0]

			// Check if filename exists in the row
			if len(row) < 3 || row[2] == "" {
				// Add process to log
				logMsg := fmt.Sprintf("No filename found for SKU: %s", sku)
				fmt.Println(logMsg)
				appendToFile(logFile, logMsg)

				// Append SKU to slice
				swf = append(swf, sku)

				// Continue the loop
				continue
			}

			// Get the filename from the row
			filename := row[2]

			// Notify the process in console
			logMsg := fmt.Sprintf("Processing SKU: %s, filename: %s", sku, filename)
			fmt.Println(logMsg)
			appendToFile(logFile, logMsg)

			// Extract the file extension
			ext := filepath.Ext(filename)
			if ext == "" {
				// If there's no extension, default to .jpg or handle as needed
				ext = ".jpg"
			}

			// Construct output filename with original extension
			outputFilename := fmt.Sprintf("%s%s", sku, ext)

			// Copy the file to the output directory
			err := copyFile(filepath.Join("files", filename), filepath.Join("output", outputFilename))
			handleErr(err, false, fmt.Sprintf("Could not copy file: %v", err))
		}
	}

	// If there are any SKUs without filenames, create a text file with the SKUs
	if len(swf) > 0 {
		// Create a text file with a timestamp
		swfFilename := fmt.Sprintf("SKUs-without-files-%s.txt", time.Now().Format(time.RFC3339))
		swfFile, err := os.Create(swfFilename)
		handleErr(err, true, fmt.Sprintf("Could not create SKUs without files file: %v", err))
		defer swfFile.Close()

		// Append SKUs to the file
		for _, sku := range swf {
			appendToFile(swfFile, sku)
		}
	}
}
