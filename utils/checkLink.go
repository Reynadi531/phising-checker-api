package utils

import (
	"io/ioutil"
	"strings"
)

func CheckLink(domain string) (bool, error) {

	f, err := ioutil.ReadFile("links.txt")
	if err != nil {
		return false, err
	}

	s := string(f)

	return strings.Contains(s, domain), nil
}
