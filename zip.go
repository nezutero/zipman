package zipman

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CompressToZip(filename string, filesToCompress []string) {
	zipFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("[ERROR] error creating zip file:", err)
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range filesToCompress {
		err := AddFileToZip(zipWriter, file)
		if err != nil {
			fmt.Println("[ERROR] error adding file to zip:", err)
			return
		}
	}

	fmt.Println("[SUCCESS] Zip compression completed successfully for file:", filename)
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("[ERROR] error opening zip")
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		fmt.Println("[ERROR] error getting file info")
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		fmt.Println("[ERROR] error getting file header")
		return err
	}

	header.Name = filepath.Base(filename)

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		fmt.Println("[ERROR] error creating file header")
		return err
	}

	_, err = io.Copy(writer, file)
	fmt.Println("[SUCCESS] added to zip successfully")
	return err
}

func ExtractZip(zipFilename, destFoldername string) error {
	zipFile, err := zip.OpenReader(zipFilename)
	if err != nil {
		fmt.Println("[ERROR] error opening zip")
		return err
	}
	defer zipFile.Close()

	err = os.MkdirAll(destFoldername, 0755)
	if err != nil {
		fmt.Println("[ERROR] error creating folder: ", destFoldername)
		return err
	}

	for _, file := range zipFile.File {
		destPath := filepath.Join(destFoldername, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(destPath, file.Mode())
			continue
		}

		writer, err := os.Create(destPath)
		if err != nil {
			fmt.Println("[ERROR] error creating zip")
			return err
		}
		defer writer.Close()

		reader, err := file.Open()
		if err != nil {
			fmt.Println("[ERROR] error reading zip")
			return err
		}
		defer reader.Close()

		_, err = io.Copy(writer, reader)
		if err != nil {
			fmt.Println("[ERROR] error copying zip")
			return err
		}
	}

	fmt.Println("[SUCCESS] zip extracted successfully")
	return nil
}
