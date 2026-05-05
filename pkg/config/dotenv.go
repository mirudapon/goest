package config

import (
	"bufio"
	"os"
	"strings"
)

func loadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)

		var key string = strings.TrimSpace(parts[0])

		if os.Getenv(key) != "" {
			continue
		}

		var value string
		for _, v := range parts[1:] {
			value += v
		}

		value = strings.TrimSpace(value)

		os.Setenv(key, value)
	}

	return nil
}
