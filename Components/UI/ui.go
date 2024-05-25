package UI

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func PrintBanner(centered bool, lines ...interface{}) {
	width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80 // Default width if unable to get terminal size
	}

	banner := []string{
		fmt.Sprintf("%s", strings.Repeat("=", width)),
	}
	var msg string
	for _, line := range lines {
		if centered {
			msg = fmt.Sprintf(CenterText(line.(string), width))
		} else {
			msg = fmt.Sprintf(line.(string))
		}
		banner = append(banner, msg)
	}

	banner = append(banner, fmt.Sprintf("%s", strings.Repeat("=", width)))

	for _, line := range banner {
		fmt.Println(line)
	}
}

func CenterText(text string, width int) string {
	padding := (width - len(text)) / 2
	return strings.Repeat(" ", padding) + text
}

func Read(multiLineHead bool, headLines ...interface{}) (string, error) {
	head := []string{}
	for _, line := range headLines {
		head = append(head, fmt.Sprintf(line.(string)))
	}
	scanner := bufio.NewScanner(os.Stdin)
	for _, line := range head {
		if multiLineHead {
			fmt.Println(line)
		} else {
			fmt.Print(line)
		}
	}
	for scanner.Scan() {
		input := scanner.Text()
		if len(input) > 0 {
			return input, nil
		}
	}
	return "", errors.New("Please input correctly!")
}

func ClearScreen() error {
	var clearFunc func()

	switch runtime.GOOS {
	case "linux":
		clearFunc = func() {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	case "android":
		clearFunc = func() {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	case "windows":
		clearFunc = func() {
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	default:
		return errors.New("Your platform is unsupported! I can't clear terminal screen :(")
	}
	clearFunc()
	return nil
}
