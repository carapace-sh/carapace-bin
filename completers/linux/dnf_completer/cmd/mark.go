package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var markCmd = &cobra.Command{
	Use:   "mark <subcommand> [options] [<group-id>] <package-spec>...",
	Short: "change the reason of an installed package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var markUserCmd = &cobra.Command{
	Use:   "user <package-spec>...",
	Short: "mark package as user-installed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var markDependencyCmd = &cobra.Command{
	Use:   "dependency <package-spec>...",
	Short: "mark package as a dependency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var markWeakCmd = &cobra.Command{
	Use:   "weak <package-spec>...",
	Short: "mark package as a weak dependency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var markGroupCmd = &cobra.Command{
	Use:   "group <group-id> <package-spec>...",
	Short: "mark package as installed by group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(markCmd).Standalone()

	markCmd.AddCommand(markUserCmd)
	markCmd.AddCommand(markDependencyCmd)
	markCmd.AddCommand(markWeakCmd)
	markCmd.AddCommand(markGroupCmd)

	for _, subcmd := range []*cobra.Command{markUserCmd, markDependencyCmd, markWeakCmd, markGroupCmd} {
		subcmd.Flags().Bool("skip-unavailable", false, "allow skipping packages that are not installed")
	}

	rootCmd.AddCommand(markCmd)

	for _, subcmd := range []*cobra.Command{markUserCmd, markDependencyCmd, markWeakCmd, markGroupCmd} {
		carapace.Gen(subcmd).PositionalAnyCompletion(
			action.ActionPackageSearch(subcmd),
		)
	}
}
