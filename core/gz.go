package core

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CompressToGz(gzFilename string, filesToCompress []string) error {
	outputFile, err := os.Create(gzFilename)
	if err != nil {
		fmt.Println("[ERROR] error creating gzip")
		return err
	}
	defer outputFile.Close()

	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	for _, filename := range filesToCompress {
		inputFile, err := os.Open(filename)
		if err != nil {
			fmt.Println("[ERROR] error opening gzip")
			return err
		}
		defer inputFile.Close()

		_, err = io.Copy(gzipWriter, inputFile)
		if err != nil {
			fmt.Println("[ERROR] error copying")
			return err
		}
	}

	fmt.Println("[SUCCESS] gzip compression completed successfully for file:", gzFilename)
	return nil
}

func AddFileToGz(existingFilename, filename string) error {
	existingFile, err := os.Open(existingFilename)
	if err != nil {
		fmt.Println("[ERROR] error opening existing gzip file")
		return err
	}
	defer existingFile.Close()

	newFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("[ERROR] error opening file")
		return err
	}
	defer newFile.Close()

	outputFilename := existingFilename
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	io.Copy(gzipWriter, existingFile)
	io.Copy(gzipWriter, newFile)

	return nil
}

func ExtractFromGz(gzFilename, destFolder string) error {
	inputFile, err := os.Open(gzFilename)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	gzipReader, err := gzip.NewReader(inputFile)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	for {
		filename, err := gzipReader.Header()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		extractedFilePath := filepath.Join(destFolder, filename.Name)
		extractedFile, err := os.Create(extractedFilePath)
		if err != nil {
			return err
		}
		defer extractedFile.Close()

		_, err = io.Copy(extractedFile, gzipReader)
		if err != nil {
			return err
		}
	}

	return nil
}
