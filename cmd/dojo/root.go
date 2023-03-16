package dojo

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use: "dojo",
  Short: "dojo - a simple CLI wrapper for DefectDojo's API",

  Run: func(cmd *cobra.Command, args []string){

  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintf(os.Stderr, "There was an error while initializing dojo CLI: '%s", err)
    os.Exit(1)
  }
}
