package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hooks_help_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a hook by name (passes `--` arguments through to the script)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hooks_help_runCmd).Standalone()

	hooks_helpCmd.AddCommand(hooks_help_runCmd)
}
