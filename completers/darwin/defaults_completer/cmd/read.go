package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/defaults"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Print the value for the given key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(readCmd).Standalone()
	rootCmd.AddCommand(readCmd)
	carapace.Gen(readCmd).PositionalCompletion(
		defaults.ActionDomains(),
	)
}
