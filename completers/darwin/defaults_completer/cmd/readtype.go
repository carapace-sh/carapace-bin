package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/defaults"
	"github.com/spf13/cobra"
)

var readTypeCmd = &cobra.Command{
	Use:   "read-type",
	Short: "Print the plist type for the given key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(readTypeCmd).Standalone()
	rootCmd.AddCommand(readTypeCmd)
	carapace.Gen(readTypeCmd).PositionalCompletion(
		defaults.ActionDomains(),
	)
}
