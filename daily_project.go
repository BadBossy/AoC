package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	today := time.Now()
	currentYear := today.Year()

	// Iterate over all days from the start date to today
	for year := 2024; year <= currentYear; year++ {
		yearDir := fmt.Sprintf("%d", year)
		// Create the year directory if it doesn't exist
		err := os.Mkdir(yearDir, 0755)
		if err != nil && !os.IsExist(err) {
			fmt.Printf("Error creating year directory %s: %v\n", yearDir, err)
			continue
		}

		startDate := time.Date(year, 12, 1, 0, 0, 0, 0, time.UTC)
		endDate := today
		if year < currentYear {
			endDate = time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)
		}

		for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
			dirName := fmt.Sprintf("%s/day_%s", yearDir, d.Format("2006-01-02"))

			// Create the directory for the day
			err := os.Mkdir(dirName, 0755)
			if err != nil {
				if os.IsExist(err) {
					fmt.Printf("Directory %s already exists\n", dirName)
					continue
				}
				fmt.Printf("Error creating directory: %v\n", err)
				continue
			}
			fmt.Printf("Created directory: %s\n", dirName)

			// Create the main.go file
			mainFilePath := fmt.Sprintf("%s/main.go", dirName)
			mainFileContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, Advent of Code!")
}`
			err = os.WriteFile(mainFilePath, []byte(mainFileContent), 0644)
			if err != nil {
				fmt.Printf("Error creating main.go: %v\n", err)
				continue
			}
			fmt.Printf("Created file: %s\n", mainFilePath)

			// Create a README.md file
			readmeFilePath := fmt.Sprintf("%s/README.md", dirName)
			readmeContent := fmt.Sprintf("# Day %s\n\nSolution for Advent of Code, Day %s.", d.Format("2006-01-02"), d.Format("2006-01-02"))
			err = os.WriteFile(readmeFilePath, []byte(readmeContent), 0644)
			if err != nil {
				fmt.Printf("Error creating README.md: %v\n", err)
				continue
			}
			fmt.Printf("Created file: %s\n", readmeFilePath)
		}
	}

	fmt.Println("Project setup complete!")
}

