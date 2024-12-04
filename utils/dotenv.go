package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadDotEnv() error {
	dotenv, err := os.ReadFile(".env")

	if err != nil {
		return err
	}

	lines := SplitLinesFieldsWithDelim(string(dotenv), "=")

	for i, line := range lines {
		if len(line) != 2 {
			actualLine := strings.Join(line, "=")
			return fmt.Errorf("error on line %d (%s): number of fields is not 2", i, actualLine)
		}

		key := line[0]
		value := line[1]

		_, exists := os.LookupEnv(key)

		if !exists {
			os.Setenv(key, value)
		}
	}

	return nil
}
