package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	specsDir string
	docsDir  string
)

var rootCmd = &cobra.Command{
	Use:   "dfspec",
	Short: "Dark Factory Spec Manager",
	Long: `dfspec is a CLI tool for managing dark factory specifications.

It provides commands to:
  - Validate specs against schemas and constitution
  - Initialize new feature directories
  - Move features between statuses (active, backlog, completed, archived)
  - Generate human-readable documentation from specs`,
	Version: "0.1.0",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&specsDir, "specs-dir", "specs", "Path to specs directory")
	rootCmd.PersistentFlags().StringVar(&docsDir, "docs-dir", "docs", "Path to docs directory")

	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(moveCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(executeCmd)
	rootCmd.AddCommand(checkpointCmd)
}

// printf is a helper for consistent output
func printf(format string, args ...any) {
	fmt.Printf(format, args...)
}

// errorf prints to stderr
func errorf(format string, args ...any) {
	_, _ = fmt.Fprintf(rootCmd.ErrOrStderr(), format, args...)
}
