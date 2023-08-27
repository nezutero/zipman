package zipman

import (
	"fmt"
	"io"
	"os"

	"github.com/dsnet/compress/bzip2"
)

func CompressToBzip2(filename, compressedFilename string) error {
	inputFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("[ERROR] error opening file: ", filename)
		return err
	}
	defer inputFile.Close()

	compressedFile, err := os.Create(compressedFilename)
	if err != nil {
		fmt.Println("[ERROR] error creating file: ", filename)
		return err
	}
	defer compressedFile.Close()

	bz2Writer, err := bzip2.NewWriter(compressedFile, nil)
	if err != nil {
		fmt.Println("[ERROR] error creating bzip2 writer")
		return err
	}
	defer bz2Writer.Close()

	_, err = io.Copy(bz2Writer, inputFile)
	if err != nil {
		fmt.Println("[ERROR] error copying")
		return err
	}

	fmt.Println("[SUCCESS] bzip2 compression completed succesfully for file: ", filename)
	return nil
}

func ExtractFromBzip2(compressedFilename, extractedFilename string) error {
	compressedFile, err := os.Open(compressedFilename)
	if err != nil {
		return err
	}
	defer compressedFile.Close()

	bz2Reader, err := bzip2.NewReader(compressedFile, nil)
	if err != nil {
		fmt.Println("[ERROR] error creating bzip2 reader")
		return err
	}

	extractedFile, err := os.Create(extractedFilename)
	if err != nil {
		fmt.Println("[ERROR] error creating extracted file")
		return err
	}
	defer extractedFile.Close()

	_, err = io.Copy(extractedFile, bz2Reader)
	if err != nil {
		fmt.Println("[ERROR] error copying")
		return err
	}

	fmt.Println("[SUCCESS] bzip2 extraction completed successfully")
	return nil
}
