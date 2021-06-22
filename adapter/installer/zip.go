package installer

import (
	"fmt"
	"net/http"
)

func downloadZip(dir, url string) (string, error) {
	fmt.Printf("downloading %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	zip.de
}
