package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Not implemented in pnpm. Use the npm CLI directly",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(editCmd).Standalone()

	editCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(editCmd)
}
