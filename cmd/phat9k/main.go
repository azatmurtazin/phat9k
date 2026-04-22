package main

import (
	"fmt"
	"os"

	"github.com/azatmurtazin/phat9k/cmd"
)

func main() {
	if err := cmd.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
