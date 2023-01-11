package main

import (
	"os"

	"github.com/devops-kung-fu/meadow/cmd"
)

func main() {
	defer os.Exit(0)
	cmd.Execute()
}
