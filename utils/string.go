package utils

import "net/url"

func URLDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}
