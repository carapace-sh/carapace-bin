package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var systemAliasCmd = &cobra.Command{
	Use:   "system-alias <name>",
	Short: "Set controller alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(systemAliasCmd).Standalone()
	rootCmd.AddCommand(systemAliasCmd)
}
