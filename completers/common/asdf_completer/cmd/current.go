package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Display current version set or being used for all packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(currentCmd).Standalone()

	currentCmd.Flags().Bool("no-header", false, "Whether or not to print a header line")
	rootCmd.AddCommand(currentCmd)
}
