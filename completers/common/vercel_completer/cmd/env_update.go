package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var env_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the value of an existing Environment Variable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(env_updateCmd).Standalone()

	env_updateCmd.Flags().String("project", "", "Project name or ID")
	env_updateCmd.Flags().Bool("sensitive", false, "Mark the environment variable as sensitive")
	env_updateCmd.Flags().String("value", "", "Value for the environment variable")
	env_updateCmd.Flags().String("visibility", "", "Visibility of the environment variable")
	env_updateCmd.Flags().Bool("yes", false, "Skip confirmation")

	envCmd.AddCommand(env_updateCmd)

	carapace.Gen(env_updateCmd).FlagCompletion(carapace.ActionMap{
		"visibility": carapace.ActionValues("auto", "encrypted", "plain"),
	})

	carapace.Gen(env_updateCmd).PositionalCompletion(
		action.ActionEnvironmentVariables(env_updateCmd),
		action.ActionEnvironments(),
	)
}
