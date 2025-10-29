package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_override_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List directory toolchain overrides",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_override_listCmd).Standalone()

	help_overrideCmd.AddCommand(help_override_listCmd)
}
