package zipman

import (
	"fmt"

	archiver "github.com/mholt/archiver/v3"
)

func CompressTo7z(filename string, filesToCompress []string) error {
	err := archiver.Archive(filesToCompress, filename)
	if err != nil {
		fmt.Println("[ERROR] error creating 7z:", err)
		return err
	}

	fmt.Println("[SUCCESS] 7z compression completed successfully for file:", filename)
	return nil
}

func Extract7z(filename, destFoldername string) error {
	err := archiver.Unarchive(filename, destFoldername)
	if err != nil {
		fmt.Println("[ERROR] error extracting 7z:", err)
		return err
	}

	fmt.Println("[SUCCESS] 7z extracted successfully")
	return nil
}
