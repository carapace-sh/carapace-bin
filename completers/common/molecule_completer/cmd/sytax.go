package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/molecule"
	"github.com/spf13/cobra"
)

var syntaxCmd = &cobra.Command{
	Use:   "syntax [flags]",
	Short: "Use the provisioner to syntax check the role",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syntaxCmd)

	syntaxCmd.Flags().Bool("help", false, "Show help message and exit")
	syntaxCmd.Flags().StringP("scenario-name", "s", "default", "Name of the scenario to target")

	carapace.Gen(syntaxCmd).FlagCompletion(carapace.ActionMap{
		"scenario-name": molecule.ActionScenarios(),
	})

	rootCmd.AddCommand(syntaxCmd)
}
