package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/defaults"
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename the given key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(renameCmd).Standalone()
	rootCmd.AddCommand(renameCmd)
	carapace.Gen(renameCmd).PositionalCompletion(
		defaults.ActionDomains(),
	)
}
