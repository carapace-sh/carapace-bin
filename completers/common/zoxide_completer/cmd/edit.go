package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(editCmd).Standalone()

	editCmd.Flags().BoolP("help", "h", false, "Print help")
	editCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.AddCommand(editCmd)
}
