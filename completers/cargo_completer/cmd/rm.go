package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a dependency from a Cargo.toml manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()

	rmCmd.Flags().BoolP("build", "B", false, "Remove crate as build dependency")
	rmCmd.Flags().BoolP("dev", "D", false, "Remove crate as development dependency")
	rmCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rmCmd.Flags().String("manifest-path", "", "Path to the manifest to remove a dependency from")
	rmCmd.Flags().StringP("package", "p", "", "Package id of the crate to remove this dependency from")
	rmCmd.Flags().BoolP("quiet", "q", false, "Do not print any output in case of success")
	rmCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.AddCommand(rmCmd)

	// TODO package?
	carapace.Gen(rmCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(".toml"),
	})

	carapace.Gen(rmCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionDependencies(rmCmd, false).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
