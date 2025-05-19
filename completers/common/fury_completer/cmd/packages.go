package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packagesCmd = &cobra.Command{
	Use:     "packages",
	Short:   "List packages in this account",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packagesCmd).Standalone()

	rootCmd.AddCommand(packagesCmd)
}
