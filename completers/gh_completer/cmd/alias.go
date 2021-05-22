package cmd

import (
	"github.com/spf13/cobra"
)

var aliasCmd = &cobra.Command{
	Use:   "alias <command>",
	Short: "Create command shortcuts",
}

func init() {
	rootCmd.AddCommand(aliasCmd)
}
