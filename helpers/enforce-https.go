package helpers

func EnforceHTTPS(url string) string {
	if url[:4] != "http" {
		return "https://" + url
	}

	return url
}
