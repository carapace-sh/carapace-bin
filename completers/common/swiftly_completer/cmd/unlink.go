package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unlinkCmd = &cobra.Command{
	Use:   "unlink",
	Short: "Unlink swiftly from the active toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unlinkCmd).Standalone()

	unlinkCmd.Flags().BoolP("help", "h", false, "Show help information")

	rootCmd.AddCommand(unlinkCmd)
}
