package main

import (
	"fmt"
	"io"
	"os"
)

// --------------------- IO ---------------------
func appendToFile(file *os.File, content string) error {
	// Write the content to the file.
	_, err := file.WriteString(fmt.Sprintln(content))
	handleErr(err, false, fmt.Sprintf("Could not write %v to file: %v", content, err))
	return nil
}

func createDirIfNotExists(dirname string) error {
	if _, err := os.Stat("output"); os.IsNotExist(err) {
		// Directory does not exist
		err := os.MkdirAll("output", 0755)
		handleErr(err, true, fmt.Sprintf("Could not create directory: %v", err))
	}

	return nil
}

func copyFile(src, dst string) error {
	// Open the source file for reading.
	srcFile, err := os.Open(src)
	handleErr(err, false, fmt.Sprintf("Could not open source file: %v", err))
	defer srcFile.Close()

	// Create the destination file for writing.
	dstFile, err := os.Create(dst)
	handleErr(err, false, fmt.Sprintf("Could not create destination file: %v", err))
	defer dstFile.Close()

	// Copy the contents from the source file to the destination file.
	_, err = io.Copy(dstFile, srcFile)
	handleErr(err, false, fmt.Sprintf("Could not copy contents: %v", err))

	return nil
}
