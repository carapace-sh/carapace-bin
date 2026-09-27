package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Alias of `clean`: same behavior, except a `purge` script (not a `clean` script) overrides it when present",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_purgeCmd).Standalone()

	helpCmd.AddCommand(help_purgeCmd)
}
