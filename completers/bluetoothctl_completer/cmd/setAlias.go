package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setAliasCmd = &cobra.Command{
	Use:   "set-alias <alias>",
	Short: "Set device alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setAliasCmd).Standalone()
	rootCmd.AddCommand(setAliasCmd)
}
