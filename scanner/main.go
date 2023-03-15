package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := []string{}

	for {
		scanner.Scan()

		line := scanner.Text()

		if line == "" {
			break
		}

		lines = append(lines, line)
	}

	fmt.Println(lines)
}
