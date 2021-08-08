package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var globalCmd = &cobra.Command{
	Use:   "global",
	Short: "Work with global packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalCmd).Standalone()

	globalCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.AddCommand(globalCmd)
}
