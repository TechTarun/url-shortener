package idgen

import "errors"

const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func EncodeToBase62(num uint64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	result := make([]byte, 0)

	for num > 0 {
		rem := num % 62
		result = append([]byte{base62Chars[rem]}, result...)
		num = num / 62
	}

	return string(result)
}

func PadShortCode(code string, length int) (string, error) {
	if len(code) > length {
		return "", errors.New("Configuration Error: generated code length more than fixed length")
	} else {
		for len(code) < length {
			code = string(base62Chars[0]) + code
		}
	}
	return code, nil
}
