package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	// Build the path to the file containing directory strings.
	dilipa := filepath.Join(home, "Projects", "dir.txt")

	// Open the file for reading.
	file, err := os.Open(dilipa)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read file contents.
	dirb, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	
	dirs := strings.TrimSpace(string(dirb))
	// Split the string into a slice using commas.
	split := strings.Split(dirs, ",")
	fmt.Println("Directories:", split)

	// Access a specific element (e.g., the third one, if it exists).
	if len(split) > 2 {
		fmt.Println("Third directory:", strings.TrimSpace(split[2]))
	}

	// Loop through each directory in the slice.
	for i := 0; i < len(split); i++ {
		// Clean up the directory string.
		dir := strings.TrimSpace(split[i])
		// Assuming there's a bash script named "script.sh" in each directory,
		// we create a command to run that script.
		cmd := exec.Command("bash", "script.sh")
		// Set the command's working directory.
		cmd.Dir = dir

		// Execute the command and capture the combined output (stdout and stderr).
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error executing script in %s: %v\n", dir, err)
			continue
		}
		fmt.Printf("Output in %s:\n%s\n", dir, string(output))
	}
}

