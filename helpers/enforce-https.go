package helpers

func EnforceHTTPS(url string) string {
	if url[:4] != "http" && url[:4] != "data" {
		return "https://" + url
	}

	return url
}
