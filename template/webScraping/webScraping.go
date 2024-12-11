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

func GetWebScrapedData(year string, day string, getInput bool) string{
	// URL to scrape
	url := "https://adventofcode.com/"+year+"/day/"+day

	if getInput {
		url +="/input"
	}

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

func HtmlToReadme(htmlContent string, year string, day string) string {
	// Extract content within <main> tags
	mainTagRegex := regexp.MustCompile(`(?s)<main>(.*?)</main>`)
	mainContentMatch := mainTagRegex.FindStringSubmatch(htmlContent)
	if len(mainContentMatch) < 2 {
		return ""
	}
	mainContent := mainContentMatch[1]

	// Process <h2> tags: Remove leading/trailing "---" and whitespace, add "##" and double newline
	h2TagRegex := regexp.MustCompile(`(?s)<h2>(.*?)</h2>`)
	h2Found := false
	mainContent = h2TagRegex.ReplaceAllStringFunc(mainContent, func(match string) string {
		inner := h2TagRegex.FindStringSubmatch(match)[1]
		inner = strings.Trim(inner, "- \t\n") // Remove "---" and surrounding whitespace
		header := fmt.Sprintf("## %s\n\n", inner)
		if !h2Found {
			h2Found = true
			linkText := fmt.Sprintf("[Here](https://adventofcode.com/%s/day/%s) is the link to the problem page on advent of code.\n\nThe input data for the puzzle can be found in the text file input.txt.\n\n# Part 1\n\n", year, day)
			return header + linkText
		}
		return header
	})

	// Prepend "- " to each <li> tag content
	liTagRegex := regexp.MustCompile(`(?s)<li>(.*?)</li>`)
	mainContent = liTagRegex.ReplaceAllString(mainContent, "- $1")

	// Replace <code> tags with backticks (`).
	codeTagRegex := regexp.MustCompile(`(?s)<code>(.*?)</code>`)
	mainContent = codeTagRegex.ReplaceAllString(mainContent, "`$1`")

	// Replace <p> tags with double newlines.
	pTagRegex := regexp.MustCompile(`(?s)<p>(.*?)</p>`)
	mainContent = pTagRegex.ReplaceAllString(mainContent, "$1\n")

	// Process <em> tags: Wrap content in ** unless inside <code>
	emTagRegex := regexp.MustCompile(`(?s)<em>(.*?)</em>`)
	mainContent = emTagRegex.ReplaceAllStringFunc(mainContent, func(match string) string {
		inner := emTagRegex.FindStringSubmatch(match)[1]
		// Skip if nested in <code>
		if strings.Contains(match, "<code>") { 
			return match
		}
		return fmt.Sprintf("**%s**", inner)
	})

	// Remove all remaining HTML tags.
	remainingTagsRegex := regexp.MustCompile(`<[^>]*>`)
	mainContent = remainingTagsRegex.ReplaceAllString(mainContent, "")

	// Ignore everything below "To begin, get your puzzle input."
	splitPhrase := "To begin, get your puzzle input."
	if idx := strings.Index(mainContent, splitPhrase); idx != -1 {
		mainContent = mainContent[:idx]
	}

	return strings.TrimSpace(mainContent)
}