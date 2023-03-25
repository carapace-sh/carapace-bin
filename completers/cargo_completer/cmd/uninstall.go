package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Short:   "Remove a Rust binary",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("uninstall"),
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().StringSlice("bin", []string{}, "Only uninstall the binary NAME")
	uninstallCmd.Flags().BoolP("help", "h", false, "Print help")
	uninstallCmd.Flags().StringSliceP("package", "p", []string{}, "Package to uninstall")
	uninstallCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	uninstallCmd.Flags().String("root", "", "Directory to uninstall packages from")
	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).FlagCompletion(carapace.ActionMap{
		"bin": action.ActionTargets(uninstallCmd, action.TargetOpts{Bin: true}),
		"package": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionInstalledPackages(uninstallCmd.Flag("root").Value.String())
		}),
		"root": carapace.ActionDirectories(),
	})

	carapace.Gen(uninstallCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionInstalledPackages(uninstallCmd.Flag("root").Value.String()).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
