package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_show_runtimeEnvCmd = &cobra.Command{
	Use:   "runtime-env",
	Short: "List run environment variables for all stacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_show_runtimeEnvCmd).Standalone()

	debug_showCmd.AddCommand(debug_show_runtimeEnvCmd)
}
