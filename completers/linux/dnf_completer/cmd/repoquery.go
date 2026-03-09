package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repoqueryCmd = &cobra.Command{
	Use:     "repoquery [options] [<package-spec>...]",
	Aliases: []string{"rq"},
	Short:   "search for packages in repositories",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoqueryCmd).Standalone()

	repoqueryCmd.Flags().String("arch", "", "limit to packages of given architectures")
	repoqueryCmd.Flags().Bool("available", false, "query available packages")
	repoqueryCmd.Flags().Bool("disable-modular-filtering", false, "include packages of inactive module streams")
	repoqueryCmd.Flags().Bool("duplicates", false, "limit to installed duplicate packages")
	repoqueryCmd.Flags().Bool("exactdeps", false, "limit to packages that require the specified capability")
	repoqueryCmd.Flags().Bool("extras", false, "limit to installed packages not in any repository")
	repoqueryCmd.Flags().StringSlice("file", []string{}, "limit to packages that own files")
	repoqueryCmd.Flags().Bool("installed", false, "query installed packages")
	repoqueryCmd.Flags().Bool("installonly", false, "limit to installed installonly packages")
	repoqueryCmd.Flags().Int("latest-limit", 0, "limit to N latest packages")
	repoqueryCmd.Flags().Bool("leaves", false, "limit to groups of installed packages not required by others")

	rootCmd.AddCommand(repoqueryCmd)

	carapace.Gen(repoqueryCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(repoqueryCmd),
	)
}
