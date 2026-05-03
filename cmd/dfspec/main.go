// Package main provides the dfspec CLI for managing dark factory specs.
package main

import (
	"os"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
