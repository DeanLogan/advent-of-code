package webScraping

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

func DecompressBody(body []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	uncompressed, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return uncompressed, nil
}

func FetchDataWithHeaders(url string, headers map[string]string) (string, error) {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Set the provided headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Check if the response contains an indication of an invalid cookie
	if strings.Contains(string(body), "Invalid Cookie") {
		return "", fmt.Errorf("invalid cookie")
	}

	// Check if the response is compressed
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		body, err = DecompressBody(body)
		if err != nil {
			return "", err
		}
	}

	// Convert the body to a string and return
	return string(body), nil
}

func GetWebScrapedData(year string, day string) string{
	// URL to scrape
	url := "https://adventofcode.com/"+year+"/day/"+day+"/input"

	godotenv.Load()
	cookieSessionKey := os.Getenv("COOKIE")

	// Request headers
	headers := map[string]string{
		"Accept":           "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Accept-Encoding":  "gzip, deflate, br",
		"Accept-Language":  "en-GB,en-US;q=0.9,en;q=0.8",
		"Cache-Control":    "max-age=0",
		"Cookie":           "session="+cookieSessionKey,
		"Sec-Fetch-Dest":   "document",
		"Sec-Fetch-Mode":   "navigate",
		"Sec-Fetch-Site":   "same-origin",
		"Sec-Fetch-User":   "?1",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	}

	// Make the request
	htmlContent, err := FetchDataWithHeaders(url, headers)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	if strings.Contains(htmlContent, "500 Internal Server Error") {
		fmt.Println("Error: Invalid Cookie")
		return ""
	} 

	if htmlContent == "Please don't repeatedly request this endpoint before it unlocks! The calendar countdown is synchronized with the server time; the link will be enabled on the calendar the instant this puzzle becomes available." || htmlContent == "404 Page Not Found" {
		return ""
	}
	
	return htmlContent
}

func HtmlToReadme(htmlContent string) string {
	// Remove HTML tags
	re := regexp.MustCompile("<[^>]*>")
	cleanedInput := re.ReplaceAllString(htmlContent, "")

	// Replace multiple consecutive spaces with a single space
	cleanedInput = strings.Join(strings.Fields(cleanedInput), " ")

	// Extract relevant content using regex
	// match := regexp.MustCompile(`# Day[\s\S]*?Part One ---`).FindStringSubmatch(cleanedInput)
	// if len(match) < 1 {
	// 	return "Error: Could not find relevant content in input."
	// }
	content := cleanedInput

	// Replace specific patterns to format the text
	content = strings.ReplaceAll(content, "```\n", "")
	content = strings.ReplaceAll(content, "```", "")
	content = strings.ReplaceAll(content, "<em>", "")
	content = strings.ReplaceAll(content, "</em>", "")
	content = strings.ReplaceAll(content, "<code>", "")
	content = strings.ReplaceAll(content, "</code>", "")
	content = strings.ReplaceAll(content, "<a href", "[Here](https://adventofcode.com/2023/day/6) is the link")
	content = strings.ReplaceAll(content, "</a>", "")
	content = strings.ReplaceAll(content, "  [Share", "[Share")
	content = strings.ReplaceAll(content, "<span title", "")
	content = strings.ReplaceAll(content, "<p>", "")
	content = strings.ReplaceAll(content, "</p>", "")

	return content
}