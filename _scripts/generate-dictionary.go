package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dictionaryContent := "\n\n"

	i := 1

	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("failed to read each file: %w", err)
			}

			if !info.IsDir() {
				return nil
			}

			if strings.Contains(path, ".") || strings.Contains(path, "_scripts") {
				return nil
			}

			title := path

			title = strings.ReplaceAll(path, "-", " ")
			title = strings.Title(strings.ToLower(title))

			dictionaryContent += fmt.Sprintf("%d. %s - [main.go](%s/main.go)\n", i, title, path)

			i++

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	newReadme, err := generateNewReadme(dictionaryContent)
	if err != nil {
		log.Fatalf("failed to generate new readme: %v", err)
	}

	err = writeNewReadme(newReadme, "README.md")
	if err != nil {
		log.Fatalf("failed to write new readme: %v", err)
	}
}

func generateNewReadme(dictionaryReadme string) (string, error) {
	readmeFile, err := os.Open("README.md")
	if err != nil {
		return "", err
	}
	defer readmeFile.Close()

	scanner := bufio.NewScanner(readmeFile)

	skipLine := false
	newReadme := ""

	for scanner.Scan() {
		if skipLine {
			// continue reads the content if end dictionary label found
			if scanner.Text() == "<!-- end dictionary -->" {
				skipLine = false

				newReadme += scanner.Text() + "\n"
			}

			continue
		}

		newReadme += scanner.Text()

		// start skips the content if start dictionary label found
		// and put the readme dictionary on the content
		if scanner.Text() == "<!-- start dictionary -->" {
			skipLine = true

			newReadme += dictionaryReadme
		}

		newReadme += "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return newReadme, nil
}

// writeNewReadme writes the new readme file.
func writeNewReadme(newReadme, path string) error {
	newReadmeFile, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}

	defer newReadmeFile.Close()

	_, err = newReadmeFile.WriteString(newReadme)

	if err != nil {
		return err
	}

	return nil
}
