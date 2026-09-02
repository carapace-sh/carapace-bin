package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var deprecateCmd = &cobra.Command{
	Use:   "deprecate",
	Short: "Deprecate a version of a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deprecateCmd).Standalone()
	deprecateCmd.Flags().Bool("dry-run", false, "only report changes")
	deprecateCmd.Flags().String("otp", "", "one-time password")
	deprecateCmd.Flags().String("registry", "", "base URL of the npm registry")

	rootCmd.AddCommand(deprecateCmd)

	carapace.Gen(deprecateCmd).PositionalCompletion(
		action.ActionPackages(deprecateCmd),
	)
}
