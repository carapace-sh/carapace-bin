package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var smbstatCmd = &cobra.Command{
	Use:   "smbstat",
	Short: "List information about a file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(smbstatCmd).Standalone()

	smbstatCmd.Flags().StringS("f", "f", "", "Output format")

	carapace.Gen(smbstatCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionValues("Json"),
	})

	carapace.Gen(smbstatCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
