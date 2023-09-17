package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/xuri/excelize/v2"
)

func main() {
	filePath := filepath.Join("files", "Imagenes.xlsx")
	f, err := excelize.OpenFile(filePath)
	handleErr(err, true, fmt.Sprintf("Could not open file: %v", err))

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	handleErr(err, true, fmt.Sprintf("The sheet with the table contents must be called 'data': %v", err))

	// create output directory if it doesnt exist
	createDirIfNotExists("output")

	// collect skus without filename
	swf := []string{}

	// iterate over rows
	for idx, row := range rows {
		if idx > 0 {
			// get sku (should always be present)
			sku := row[0]

			// catch rows where there's no filename
			if len(row) == 2 {
				fmt.Println("No filename found for SKU: ", sku)

				// appends sku to slice
				swf = append(swf, sku)

				// continue the loop
				continue

			} else {

				// get the filename from the row
				filename := row[2]

				fmt.Printf("Processing SKU: %v, filename: %v", sku, filename)
				fmt.Println()
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
