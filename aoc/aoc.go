package aoc

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func dayUrl(day int) string {
	return fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
}

func DayInput(day int) (string, error) {
	filename := fmt.Sprintf("inputs/day%d.txt", day)
	input, err := os.ReadFile(filename)

	if err == nil {
		return string(input), nil
	}

	err = godotenv.Load(".env")

	if err != nil {
		return "", fmt.Errorf("error loading .env file: %s", err)
	}

	cookie := http.Cookie{Name: "session", Value: os.Getenv("session")}

	req, err := http.NewRequest(http.MethodGet, dayUrl(day), nil)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error making http request: %s", err)
		os.Exit(1)
	}

	req.AddCookie(&cookie)
	req.Header.Add("User-Agent", "https://github.com/FractalBoy/advent-of-code-2024 by reisner.marc@gmail.com")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", fmt.Errorf("error making http request: %s", err)
	}

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		return "", fmt.Errorf("client: could not read response body: %s", err)
	}

	os.Mkdir("inputs", 0755)
	os.WriteFile(filename, []byte(resBody), 0644)

	return string(resBody), nil
}
