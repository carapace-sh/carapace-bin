package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pkg_describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Generates a description of the given package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_describeCmd).Standalone()
	pkg_describeCmd.Flags().Bool("allow-local-deps", false, "Allow local dependencies and don't report them")
	pkg_describeCmd.Flags().Bool("disallow-local-deps", false, "Always disallow local dependencies and report them as error")
	pkg_describeCmd.Flags().String("out-dir", "", "Output directory of description files")
	pkg_describeCmd.Flags().BoolP("verbose", "v", false, "Show more information")
	pkgCmd.AddCommand(pkg_describeCmd)

	carapace.Gen(pkg_describeCmd).FlagCompletion(carapace.ActionMap{
		"out-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(pkg_describeCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
