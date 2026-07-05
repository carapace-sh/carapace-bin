package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/defaults"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Read a plist from stdin and set it as the value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()
	rootCmd.AddCommand(importCmd)
	carapace.Gen(importCmd).PositionalCompletion(
		defaults.ActionDomains(),
		carapace.ActionFiles(),
	)
}
