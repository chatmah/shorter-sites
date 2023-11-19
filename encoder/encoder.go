package encoder

import (
	"shorter-sites/handler"
)

const (
	base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func EncodeLinkToShort(linkToEncode string) (string, error) {
	count, err := handler.CountHandler()
	encodedCount := intToBase62(count)
	if err != nil {
		return "error", err
	}
	return encodedCount, err
}

func intToBase62(n int) string {
	base := len(base62Chars)
	result := ""
	for n > 0 {
		remainder := n % base
		result = string(base62Chars[remainder]) + result
		n = n / base
	}
	// Pad with leading zeros to make the result 9 characters long
	for len(result) < 9 {
		result = "0" + result
	}
	return result
}
