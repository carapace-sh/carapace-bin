package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Manage command aliases.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aliasCmd).Standalone()

	aliasCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(aliasCmd)
}
