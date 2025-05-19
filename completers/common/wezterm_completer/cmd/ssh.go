package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh [OPTIONS] <USER_AT_HOST_AND_PORT> [PROG]...",
	Short: "Establish an ssh session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshCmd).Standalone()

	sshCmd.Flags().String("class", "", "Override the default windowing system class")
	sshCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	sshCmd.Flags().String("position", "", "Override the position for the initial window launched by this process")
	sshCmd.Flags().StringP("ssh-option", "o", "", "Override specific SSH configuration options")
	sshCmd.Flags().BoolS("v", "v", false, "Enable verbose ssh protocol tracing")
	rootCmd.AddCommand(sshCmd)

	carapace.Gen(sshCmd).PositionalCompletion(
		carapace.ActionMultiPartsN("@", 2, func(cAt carapace.Context) carapace.Action {
			return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					switch len(cAt.Parts) {
					case 0:
						return carapace.Batch(
							os.ActionUsers().Suffix("@"),
							net.ActionHosts(),
						).ToA()

					default:
						return net.ActionHosts()
					}
				default:
					return net.ActionKnownPorts()
				}
			})
		}),
	)
}
