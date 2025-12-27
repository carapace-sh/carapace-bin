package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resetAliasCmd = &cobra.Command{
	Use:   "reset-alias",
	Short: "Reset controller alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resetAliasCmd).Standalone()
	rootCmd.AddCommand(resetAliasCmd)
}
