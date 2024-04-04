package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	if len(os.Args) < 2 {
		fmt.Println("Please provide a CSV file")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	outputFile, err := os.Create("results.csv")
	if err != nil {
		fmt.Println("Error creating results file:", err)
		return
	}
	defer outputFile.Close()

	var results []string
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		domain := record[0]
		wg.Add(1)
		go func(domain string) {
			defer wg.Done()
			results = append(results, checkRobots(domain)...)
		}(domain)
	}
	wg.Wait()

	fmt.Println(strings.Join(results, "\n"))
}

func checkRobots(domain string) []string {
	resp, err := http.Get("http://" + domain + "/robots.txt")
	if err != nil {
		return []string{fmt.Sprintf("%s,Unable to retrieve robots.txt. Maybe the domain doesn't exist?,", domain)}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return []string{fmt.Sprintf("%s,No robots.txt file found,", domain)}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []string{fmt.Sprintf("%s,Unable to read the contents of robots.txt,", domain)}
	}

	robotsTxt := string(body)
	sections := strings.Split(robotsTxt, "\n\n")

	googlebotDisallowed := false
	googleExtendedDisallowed := false

	for _, section := range sections {
		lines := strings.Split(section, "\n")
		if len(lines) > 0 && lines[0] == "User-agent: *" {
			for _, line := range lines[1:] {
				if line == "Disallow: /" {
					googlebotDisallowed = true
				}
			}
		} else if len(lines) > 0 && lines[0] == "User-agent: Google-Extended" {
			for _, line := range lines[1:] {
				if line == "Disallow: /" {
					googleExtendedDisallowed = true
				}
			}
		}
	}

	return []string{fmt.Sprintf("%s,%t,%t", domain, googlebotDisallowed, googleExtendedDisallowed)}
}
