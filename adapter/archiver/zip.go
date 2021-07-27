package archiver

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
)

func GetTargetFile(url, targetFile, targetDir string) (string, error) {
	zipReader, err := getZipReader(url, targetFile)
	if err != nil {
		return "", err
	}
	defer zipReader.Close()

	if fs.ValidPath(targetDir) {
		return "", fmt.Errorf("")
	}
	filePath := filepath.Join(targetDir, targetFile)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, zipReader)
	if err != nil {
		return "", err
	}
	err = os.Chmod(filePath, os.ModePerm)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func getZipReader(url string, targetFile string) (io.ReadCloser, error) {
	if targetFile == "" {
		return nil, fmt.Errorf("set target file name")
	}

	fmt.Printf("downloading %s ..\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buff := bytes.NewBuffer([]byte{})
	size, err := io.Copy(buff, resp.Body)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(buff.Bytes())
	// Open a zip archive for reading.
	zipReader, err := zip.NewReader(reader, size)

	for _, file := range zipReader.File {
		if filepath.Base(file.Name) == targetFile {
			return file.Open()
		}
	}
	return nil, fmt.Errorf("cannot find target file in zip")
}
