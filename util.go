package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func promptInput(message string) string {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
