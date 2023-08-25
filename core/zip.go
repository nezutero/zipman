package core

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Zip(filename string, filesToCompress []string) {
	zipFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("[ERROR] error creating zip file:", err)
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range filesToCompress {
		err := addFileToZip(zipWriter, file)
		if err != nil {
			fmt.Println("[ERROR] error adding file to zip:", err)
			return
		}
	}

	fmt.Println("[SUCCESS] Zip compression completed successfully for file:", filename)
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filepath.Base(filename)

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}
