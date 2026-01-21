package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var copilotCmd = &cobra.Command{
	Use:   "copilot [flags] [args]",
	Short: "Run the GitHub Copilot CLI (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(copilotCmd).Standalone()
	copilotCmd.Flags().SetInterspersed(false)

	copilotCmd.Flags().Bool("remove", false, "Remove the downloaded Copilot CLI")
	rootCmd.AddCommand(copilotCmd)

	carapace.Gen(copilotCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("copilot"),
	)
}
