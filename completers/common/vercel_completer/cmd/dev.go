package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var devCmd = &cobra.Command{
	Use:     "dev",
	Aliases: []string{"develop"},
	Short:   "Start a local development server",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devCmd).Standalone()

	devCmd.Flags().BoolP("confirm", "c", false, "(deprecated)")
	devCmd.Flags().StringP("listen", "l", "0.0.0.0:3000", "Specify a URI endpoint on which to listen")
	devCmd.Flags().BoolP("local", "L", false, "Start the dev server without linking to a Vercel project")
	devCmd.Flags().String("project", "", "Project name or ID")
	devCmd.Flags().BoolP("yes", "y", false, "Skip questions")

	rootCmd.AddCommand(devCmd)

	carapace.Gen(devCmd).FlagCompletion(carapace.ActionMap{
		"listen": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("0.0.0.0", "127.0.0.1").Invoke(c).Suffix(":").ToA()
			case 1:
				return net.ActionPorts()
			default:
				return carapace.ActionValues()
			}
		}),
		"project": action.ActionProjects(devCmd),
	})

	carapace.Gen(devCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
