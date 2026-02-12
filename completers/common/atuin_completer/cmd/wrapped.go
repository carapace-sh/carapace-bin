package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wrappedCmd = &cobra.Command{
	Use:   "wrapped",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wrappedCmd).Standalone()

	wrappedCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(wrappedCmd)
}
