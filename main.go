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

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	handleErr(err, true, fmt.Sprintf("The sheet with the table contents must be called 'data': %v", err))

	// create output directory if it doesnt exist
	createDirIfNotExists("output")

	// collect skus without filename
	swf := []string{}

	// create file to log the process
	log_file, err := os.Create(fmt.Sprintf("log-%v.txt", time.Now().Format(time.RFC3339)))
	handleErr(err, true, fmt.Sprintf("Could not create file: %v", err))

	// iterate over rows
	for idx, row := range rows {
		if idx > 0 {
			// get sku (should always be present)
			sku := row[0]

			// catch rows where there's no filename
			if len(row) < 3 {
				// add process to log
				logmsg := fmt.Sprint("No filename found for SKU: ", sku)
				fmt.Println(logmsg)
				appendToFile(log_file, logmsg)

				// appends sku to slice
				swf = append(swf, sku)

				// continue the loop
				continue

			} else {
				// get the filename from the row
				filename := row[2]

				// notify the process in console
				logmsg := fmt.Sprint("Processing SKU: ", sku, ", filename: ", filename)
				fmt.Println(logmsg)
				appendToFile(log_file, logmsg)

				// construct output filename
				outputFilename := fmt.Sprintf("%v_1.jpg", sku)

				// copy the file to the output directory
				err := copyFile(filepath.Join("files", filename), filepath.Join("output", outputFilename))
				handleErr(err, false, fmt.Sprintf("Could not copy file: %v", err))
			}

		}
	}

	// if there's any sku without filename, create a text file with the skus
	if len(swf) > 0 {
		// create a text file if it doesnt exist
		swf_file, err := os.Create(fmt.Sprintf("SKUs-without-files-%v.txt", time.Now().Format(time.RFC3339)))
		handleErr(err, true, fmt.Sprintf("Could not create file: %v", err))

		// append skus to file
		for _, sku := range swf {
			appendToFile(swf_file, sku)
		}
	}

}
