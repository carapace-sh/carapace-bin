package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pub"
	"github.com/spf13/cobra"
)

var pub_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an executable from a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_runCmd).Standalone()

	pub_runCmd.Flags().StringP("directory", "C", "", "Run this in the directory <dir>.")
	pub_runCmd.Flags().Bool("enable-asserts", false, "Enable assert statements.")
	pub_runCmd.Flags().String("enable-experiment", "", "Runs the executable in a VM with the given experiments enabled.")
	pub_runCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_runCmd.Flags().Bool("no-enable-asserts", false, "Disable assert statements.")
	pubCmd.AddCommand(pub_runCmd)

	carapace.Gen(pub_runCmd).FlagCompletion(carapace.ActionMap{
		"directory": carapace.ActionDirectories(),
	})

	carapace.Gen(pub_runCmd).PositionalCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return pub.ActionDependencies().Invoke(c).Suffix(":").ToA()
			case 1:
				return pub.ActionHostedExecutables(pub.HostedExecutablesOpts{Name: c.Parts[0], Version: ""})
			default:
				return carapace.ActionValues()
			}

		}),
	)
}
