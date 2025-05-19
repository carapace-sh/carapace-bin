package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [PATHS]...",
	Short: "Remove a directory from the database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().BoolP("help", "h", false, "Print help")
	removeCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.AddCommand(removeCmd)
}
