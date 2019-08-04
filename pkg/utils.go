package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func TerminalSize() (int, int)  {
	var width, height int
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output() // 46 94
	fields := strings.Fields(string(out))
	width, _ = strconv.Atoi(fields[1])
	height, _ = strconv.Atoi(fields[0])

	return width, height
}

func ClearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func BeautifulPrint(char string) (bool) {
	if len(char) == 1 {
		width, _ := TerminalSize()
		for i := 1; i <= width; i++ {
			fmt.Printf(char)
		}
		fmt.Printf("\n")
		return false
	} else {
		return true
	}

}