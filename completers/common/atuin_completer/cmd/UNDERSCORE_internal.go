package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var __internalCmd = &cobra.Command{
	Use:    "__internal",
	Short:  "Internal subcommands, not for direct use by users",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(__internalCmd).Standalone()

	rootCmd.AddCommand(__internalCmd)
}
