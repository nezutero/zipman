package zipman

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/nwaples/rardecode"
)

func ExtractRar(rarFilename, destFoldername string) error {
	rarFile, err := os.Open(rarFilename)
	if err != nil {
		return err
	}
	defer rarFile.Close()

	archive, err := rardecode.NewReader(rarFile, "")
	if err != nil {
		return err
	}

	err = os.MkdirAll(destFoldername, 0755)
	if err != nil {
		return err
	}

	for {
		header, err := archive.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		destPath := filepath.Join(destFoldername, header.Name)

		if header.IsDir {
			os.MkdirAll(destPath, os.FileMode(header.Mode()))
			continue
		}

		writer, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer writer.Close()

		_, err = io.Copy(writer, archive)
		if err != nil {
			return err
		}
	}

	fmt.Println("[SUCCESS] rar extracted successfully")
	return nil
}
