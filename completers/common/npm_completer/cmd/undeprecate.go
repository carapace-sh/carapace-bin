package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var undeprecateCmd = &cobra.Command{
	Use:   "undeprecate",
	Short: "Undeprecate a version of a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(undeprecateCmd).Standalone()
	undeprecateCmd.Flags().String("otp", "", "one-time password")

	rootCmd.AddCommand(undeprecateCmd)

	carapace.Gen(undeprecateCmd).PositionalCompletion(
		action.ActionPackages(undeprecateCmd),
	)
}
