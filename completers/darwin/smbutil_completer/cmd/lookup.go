package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "Resolve NetBIOS name to IP address",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookupCmd).Standalone()

	lookupCmd.Flags().BoolS("e", "e", false, "Unpercent escape NetBIOS names")
	lookupCmd.Flags().StringS("t", "t", "", "NetBIOS name type")
	lookupCmd.Flags().StringS("w", "w", "", "NetBIOS name server")

	carapace.Gen(lookupCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
