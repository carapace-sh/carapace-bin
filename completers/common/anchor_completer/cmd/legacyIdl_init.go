package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var legacyIdl_initCmd = &cobra.Command{
	Use:   "init",
	Short: "[DEPRECATED] Initialize the legacy on-chain IDL account for the first time",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(legacyIdl_initCmd).Standalone()

	legacyIdl_initCmd.Flags().StringP("filepath", "f", "", "")
	legacyIdl_initCmd.Flags().BoolP("help", "h", false, "Print help")
	legacyIdl_initCmd.Flags().String("priority-fee", "", "")
	legacyIdl_initCmd.MarkFlagRequired("filepath")
	legacyIdlCmd.AddCommand(legacyIdl_initCmd)
}
