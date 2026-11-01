package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var allowCmd = &cobra.Command{
	Use:   "allow",
	Short: "Allow auto-activation for the current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(allowCmd).Standalone()

	rootCmd.AddCommand(allowCmd)
}
