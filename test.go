package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

const (
	appName    = "MyApp"
	appVersion = "1.0.0"
	credit     = "Credit to Eikarna"
)

func printBanner() {
	width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80 // Default width if unable to get terminal size
	}

	banner := []string{
		fmt.Sprintf("%s", strings.Repeat("=", width)),
		fmt.Sprintf("%s", centerText(appName, width)),
		fmt.Sprintf("%s", centerText(credit, width)),
		fmt.Sprintf("%s", centerText("Version: "+appVersion, width)),
		fmt.Sprintf("%s", strings.Repeat("=", width)),
	}

	for _, line := range banner {
		fmt.Println(line)
	}
}

func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	return strings.Repeat(" ", padding) + text
}

func main() {
	printBanner()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter some input:")

	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			fmt.Println("Exiting...")
			break
		}
		fmt.Println("You entered:", input)
		fmt.Println("Please enter some input:")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
