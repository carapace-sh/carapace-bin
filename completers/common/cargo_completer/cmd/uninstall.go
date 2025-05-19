package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/cargo_completer/cmd/action"
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

	uninstallCmd.Flags().StringSlice("bin", nil, "Only uninstall the binary NAME")
	uninstallCmd.Flags().BoolP("help", "h", false, "Print help")
	uninstallCmd.Flags().StringSliceP("package", "p", nil, "Package to uninstall")
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
			return action.ActionInstalledPackages(uninstallCmd.Flag("root").Value.String()).FilterArgs()
		}),
	)
}
