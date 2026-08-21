package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/eopkg"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check <package?>",
	Short: "check the installation status of all packages, or the provided package names",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()

	checkCmd.Flags().StringP("component", "c", "", "check installed packages under the given component")
	checkCmd.Flags().Bool("config", false, "only check the status of configuration files")

	rootCmd.AddCommand(checkCmd)

	carapace.Gen(checkCmd).PositionalAnyCompletion(
		eopkg.ActionPackages().FilterArgs(),
	)
}
