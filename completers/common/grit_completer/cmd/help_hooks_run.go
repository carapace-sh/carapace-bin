package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_hooks_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a hook by name (passes `--` arguments through to the script)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_hooks_runCmd).Standalone()

	help_hooksCmd.AddCommand(help_hooks_runCmd)
}
