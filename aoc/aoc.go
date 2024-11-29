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

func DayInput(day int) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading .env file: %s\n", err)
		os.Exit(1)
	}

	cookie := http.Cookie{Name: "session", Value: os.Getenv("session")}

	req, err := http.NewRequest(http.MethodGet, dayUrl(day), nil)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error making http request: %s\n", err)
		os.Exit(1)
	}

	req.AddCookie(&cookie)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	return string(resBody)
}
