package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all supported platforms",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().BoolS("json", "json", false, "produce JSON output")
	listCmd.Flags().BoolS("v", "v", false, "verbosity")

	rootCmd.AddCommand(listCmd)
}
