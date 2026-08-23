package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var multichannelCmd = &cobra.Command{
	Use:   "multichannel",
	Short: "Print multichannel attributes of shares",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(multichannelCmd).Standalone()

	multichannelCmd.Flags().BoolS("a", "a", false, "Print attributes of all mounted shares")
	multichannelCmd.Flags().BoolS("c", "c", false, "Print client interface information")
	multichannelCmd.Flags().StringS("f", "f", "", "Output format")
	multichannelCmd.Flags().BoolS("i", "i", false, "Print session information")
	multichannelCmd.Flags().StringS("m", "m", "", "Print attributes of the share mounted at path")
	multichannelCmd.Flags().BoolS("s", "s", false, "Print server interface information")
	multichannelCmd.Flags().BoolS("x", "x", false, "Print established connection information")

	carapace.Gen(multichannelCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionValues("Json"),
		"m": carapace.ActionDirectories(),
	})
}
