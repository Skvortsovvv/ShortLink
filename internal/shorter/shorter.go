package shorter

import (
	"crypto/sha256"
	"fmt"
	"regexp"
)

const (
	base78 = "aA0bB1cC2dD3eE4fF5gG6hH7iI8jJ9kK0lL1mM2nN3oO4pP5qQ6rR7sS8tT9uU0vV1wW2xX3yY4zZ5"
)

var (
	pattern = regexp.MustCompile(`([a-z]|[A-Z]|[0-9]){9}`)
)

func scanning(row []byte) (string, error) {
	loc := pattern.FindIndex(row)

	if loc == nil {
		return "", fmt.Errorf("pattern not found")
	}

	raw := row[loc[0]:loc[1]]
	position := (raw[0] + raw[8]) % 10
	var result string

	if position == byte(9) {
		result = string(raw) + "_"
	} else {
		result = string(raw[:position+1]) + "_" + string(raw[position+1:])
	}

	return result, nil
}

func sumMod2(input, key []byte) (output []byte) {
	for i := 0; i < 32; i++ {
		output = append(output, input[i]^key[i%len(key)])
	}
	return output
}

func base(data []byte) []byte {
	var result []byte
	for _, runeValue := range data {
		result = append(result, base78[int(runeValue)%78])
	}
	return result
}

func Shorter(longURL string) string {
	byteURL := []byte(longURL)
	for {
		hashURL := sha256.Sum256(byteURL)
		xorResult := sumMod2(hashURL[:], byteURL)
		based := base(xorResult)
		shortURL, err := scanning(based)
		if err != nil {
			byteURL = []byte(shortURL)
			continue
		}
		return shortURL
	}
}
