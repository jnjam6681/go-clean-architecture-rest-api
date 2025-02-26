package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

func ValidName(name string) error {
	validName := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	if !validName.MatchString(name) {
		return fmt.Errorf("(%s) It should only contain letters, numbers, underscores, or hyphens", name)
	}
	return nil
}

func ValidHostKey(input string) (string, error) {
	const (
		minLength = 8
		maxLength = 256
	)

	if strings.HasPrefix(input, "file://") {
		filePath := strings.TrimPrefix(input, "file://")
		data, err := os.ReadFile(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				return "", fmt.Errorf("file not found: %s", filePath)
			}
			return "", fmt.Errorf("failed to read file: %v", err)
		}
		input = strings.TrimSpace(string(data))
	}

	length := len(input)
	if length < minLength || length > maxLength {
		return "", fmt.Errorf("host key must be between %d and %d characters long", minLength, maxLength)
	}

	validKeyRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !validKeyRegex.MatchString(input) {
		return "", errors.New("host key must only contain letters and numbers (no special characters)")
	}
	return input, nil
}
