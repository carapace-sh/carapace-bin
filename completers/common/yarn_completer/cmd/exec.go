package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:     "exec",
	Short:   "Execute a shell script",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()
	execCmd.Flags().SetInterspersed(false)

	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
