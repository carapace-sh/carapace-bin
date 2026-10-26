package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var bumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "Increases the lower limit of your composer.json requirements to the currently installed versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bumpCmd).Standalone()

	bumpCmd.Flags().BoolP("dev-only", "D", false, "Only bump requirements in \"require-dev\".")
	bumpCmd.Flags().Bool("dry-run", false, "Outputs the packages to bump, but will not execute anything.")
	bumpCmd.Flags().BoolP("no-dev-only", "R", false, "Only bump requirements in \"require\".")

	rootCmd.AddCommand(bumpCmd)

	carapace.Gen(bumpCmd).PositionalAnyCompletion(
		action.ActionRootRequirements(bumpCmd).FilterArgs(),
	)
}
