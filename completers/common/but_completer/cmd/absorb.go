package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var absorbCmd = &cobra.Command{
	Use:   "absorb",
	Short: "Amends changes into the appropriate commits where they belong.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(absorbCmd).Standalone()

	absorbCmd.Flags().Bool("dry-run", false, "Show the absorption plan without making any changes")
	absorbCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	absorbCmd.Flags().BoolP("new", "n", false, "Create new commits, instead of amending existing ones. This is useful when you want to preserve existing commits and add new ones for the absorbed changes")
	rootCmd.AddCommand(absorbCmd)

	carapace.Gen(absorbCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionLocalBranches(),
			but.ActionCliIds(but.CliIdsOpts{}.Default()), // TODO which change ids?
		).ToA(),
	)
}
