package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print Go version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()
	versionCmd.Flags().SetInterspersed(false)

	versionCmd.Flags().BoolS("json", "json", false, "output the runtime/debug.BuildInfo in JSON format")
	versionCmd.Flags().BoolS("m", "m", false, "print each executable's embedded module version information")
	versionCmd.Flags().BoolS("v", "v", false, "report unrecognized files")
	rootCmd.AddCommand(versionCmd)

	carapace.Gen(versionCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if versionCmd.Flag("m").Changed ||
				versionCmd.Flag("json").Changed {
				return carapace.ActionFiles().FilterArgs()
			}
			return carapace.ActionValues()
		}),
	)
}
