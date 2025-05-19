package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clamdtop",
	Short: "monitor the Clam AntiVirus Daemon",
	Long:  "http://www.clamav.net/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("config-file", "c", "", "Read clamd's configuration files from FILE")
	rootCmd.Flags().BoolP("defaultcolors", "d", false, "Use default terminal colors")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help")
	rootCmd.Flags().BoolP("version", "V", false, "Show version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return net.ActionHosts().NoSpace()
				case 1:
					return net.ActionPorts()
				default:
					return carapace.ActionValues()
				}
			}).UnlessF(condition.CompletingPath),
		).ToA(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
