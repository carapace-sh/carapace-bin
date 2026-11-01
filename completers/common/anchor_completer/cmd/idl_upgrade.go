package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var idl_upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrades the IDL to the new file. An alias for first writing and then then setting the idl buffer account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(idl_upgradeCmd).Standalone()

	idl_upgradeCmd.Flags().StringP("filepath", "f", "", "")
	idl_upgradeCmd.Flags().BoolP("help", "h", false, "Print help")
	idl_upgradeCmd.Flags().String("priority-fee", "", "")
	idl_upgradeCmd.MarkFlagRequired("filepath")
	idlCmd.AddCommand(idl_upgradeCmd)

	carapace.Gen(idl_upgradeCmd).FlagCompletion(carapace.ActionMap{
		"filepath": carapace.ActionFiles(".json"),
	})
}
