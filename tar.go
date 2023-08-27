package zipman

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CompressToTar(tarFilename string, filesToCompress []string) error {
	tarFile, err := os.Create(tarFilename)
	if err != nil {
		fmt.Println("[ERROR] error creating tar")
		return err
	}
	defer tarFile.Close()

	tarWriter := tar.NewWriter(tarFile)
	defer tarWriter.Close()

	for _, file := range filesToCompress {
		err := addFileToTar(tarWriter, file)
		if err != nil {
			fmt.Println("[ERROR] error adding file to tar")
			return err
		}
	}

	fmt.Println("[SUCCESS] tar compression successfully completed")
	return nil
}

func AddFileToTar(tarWriter *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("[ERROR] error adding file to tar")
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
	fmt.Println("[SUCCESS] added file to tar successfully")
	return err
}

func ExtractTar(tarFilename string, destFolder string) error {
	tarFile, err := os.Open(tarFilename)
	if err != nil {
		fmt.Println("[ERROR] error opening tar")
		return err
	}
	defer tarFile.Close()

	tarReader := tar.NewReader(tarFile)

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
				fmt.Println("[ERROR] error creating destination path")
				return err
			}
			defer file.Close()

			_, err = io.Copy(file, tarReader)
			if err != nil {
				fmt.Println("[ERROR] error copying")
				return err
			}
		}
	}

	fmt.Println("[SUCCESS] tar extracted successfully")
	return nil
}
