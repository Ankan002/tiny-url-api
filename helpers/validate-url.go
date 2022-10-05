package helpers

import (
	"os"
	"strings"
)

func ValidateUrl(url string) bool {
	urlPartsArray := strings.Split(url, "/")
	currentHost := os.Getenv("DOMAIN")

	for _, urlPart := range urlPartsArray {
		if urlPart == currentHost {
			return false
		}
	}

	return true
}
