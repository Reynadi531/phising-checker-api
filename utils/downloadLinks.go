package utils

import (
	"errors"
	"io"
	"net/http"
	"os"
)

var (
	LinksURL = "https://raw.githubusercontent.com/nikolaischunk/discord-phishing-links/main/txt/domain-list.txt"
)

func DownloadLinks(filepath string) error {
	resp, err := http.Get(LinksURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("download error: github doesnt respone with 200")
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
