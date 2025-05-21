package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Print the current Node.js version",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	rootCmd.AddCommand(currentCmd)

	addCommonFlags(currentCmd)

	carapace.Gen(currentCmd).Standalone()
}
