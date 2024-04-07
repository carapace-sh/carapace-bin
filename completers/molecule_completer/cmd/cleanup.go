package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/molecule"
	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup [flags]",
	Short: "Use the provisioner to cleanup any changes made to external systems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanupCmd)

	cleanupCmd.Flags().Bool("help", false, "Show help message and exit")
	cleanupCmd.Flags().StringP("scenario-name", "s", "default", "Name of the scenario to target")

	carapace.Gen(cleanupCmd).FlagCompletion(carapace.ActionMap{
		"scenario-name": molecule.ActionScenarios(),
	})

	rootCmd.AddCommand(cleanupCmd)
}
