package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io"
	"strings"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	dilipa := filepath.Join(home, "Projects", "dir.txt")

	// Create the file
	file, err := os.OpenFile(dilipa, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	file.Close() // Close after creation

	// Open the file (FIXED VARIABLE NAME)
	file, err = os.Open(dilipa)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	fmt.Println("File opened successfully:", file.Name())

	// Close the file
	defer file.Close()

	dirb, _ := io.ReadAll(file)
	fmt.Println(string(dirb))
	dirs := string(dirb)
	split := strings.Split(dirs, ",")
	fmt.Println(split)
	fmt.Println(split[2])





}

