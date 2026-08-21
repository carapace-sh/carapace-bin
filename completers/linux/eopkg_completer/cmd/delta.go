package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/eopkg"
	"github.com/spf13/cobra"
)

var deltaCmd = &cobra.Command{
	Use:     "delta <oldpackage1> <newpackage>",
	Aliases: []string{"dt"},
	Short:   "construct a delta package between the given packages",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deltaCmd).Standalone()

	deltaCmd.Flags().StringP("newest-package", "t", "", "override the \"new\" package detection for explicit control of the process")
	deltaCmd.Flags().StringP("output-dir", "O", "", "override the output directory for the .delta.eopkg")
	deltaCmd.Flags().StringP("package-format", "F", "", "override the eopkg internal format")

	rootCmd.AddCommand(deltaCmd)

	carapace.Gen(deltaCmd).PositionalAnyCompletion(
		eopkg.ActionPackages().FilterArgs(),
	)
}
