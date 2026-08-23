package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statsharesCmd = &cobra.Command{
	Use:   "statshares",
	Short: "Print attributes of mounted shares",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statsharesCmd).Standalone()

	statsharesCmd.Flags().BoolS("a", "a", false, "Print attributes of all mounted shares")
	statsharesCmd.Flags().StringS("f", "f", "", "Output format")
	statsharesCmd.Flags().StringS("m", "m", "", "Print attributes of the share mounted at path")

	carapace.Gen(statsharesCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionValues("Json"),
		"m": carapace.ActionDirectories(),
	})
}
