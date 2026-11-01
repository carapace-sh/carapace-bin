package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var teardownCmd = &cobra.Command{
	Use:   "teardown",
	Short: "Remove the git meta database and all meta refs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(teardownCmd).Standalone()

	teardownCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(teardownCmd)
}
