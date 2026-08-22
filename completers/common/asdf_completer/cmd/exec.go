package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:                "exec",
	Short:              "Executes the command shim for current version",
	DisableFlagParsing: true,
	Run:                func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	rootCmd.AddCommand(execCmd)
}
