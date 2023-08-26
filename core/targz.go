package core

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CompressToTarGz(tarGzFilename string, filesToCompress []string) error {
	tarGzFile, err := os.Create(tarGzFilename)
	if err != nil {
		fmt.Println("[ERROR] error creating tar")
		return err
	}
	defer tarGzFile.Close()

	gzWriter := gzip.NewWriter(tarGzFile)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	for _, file := range filesToCompress {
		err := addFileToTar(tarWriter, file)
		if err != nil {
			fmt.Println("[ERROR] error compresssing files")
			return err
		}
	}

	fmt.Println("[SUCCESS] tar.gz compresssion completed successfully for  file: ", tarGzFilename)
	return nil
}

func addFileToTar(tarWriter *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("[ERROR] error opening tar")
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		fmt.Println("[ERROR] error getting file info")
		return err
	}

	header := &tar.Header{
		Name:    filepath.Base(filename),
		Size:    info.Size(),
		Mode:    int64(info.Mode()),
		ModTime: info.ModTime(),
	}

	err = tarWriter.WriteHeader(header)
	if err != nil {
		fmt.Println("[ERROR] error writing tar header")
		return err
	}

	_, err = io.Copy(tarWriter, file)
	fmt.Println("[SUCCESS] added to tar.gz successfully")
	return err
}

func ExtractTarGz(tarGzFilename string, destFolder string) error {
	tarGzFile, err := os.Open(tarGzFilename)
	if err != nil {
		fmt.Println("[ERROR] opening tar")
		return err
	}
	defer tarGzFile.Close()

	gzReader, err := gzip.NewReader(tarGzFile)
	if err != nil {
		fmt.Println("[ERROR] error reading tar")
		return err
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)

	err = os.MkdirAll(destFolder, 0755)
	if err != nil {
		fmt.Println("[ERROR] error creating directory")
		return err
	}

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		destPath := filepath.Join(destFolder, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(destPath, os.FileMode(header.Mode))
		case tar.TypeReg:
			file, err := os.Create(destPath)
			if err != nil {
				fmt.Println("[ERROR] error creating directory")
				return err
			}
			defer file.Close()

			_, err = io.Copy(file, tarReader)
			if err != nil {
				fmt.Println("[ERROR] error copying file")
				return err
			}
		}
	}

	fmt.Println("[SUCCESS] tar.gz extracted successfully")
	return nil
}
