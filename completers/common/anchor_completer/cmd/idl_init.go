package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var idl_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a program's IDL account. Can only be run once",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(idl_initCmd).Standalone()

	idl_initCmd.Flags().StringP("filepath", "f", "", "")
	idl_initCmd.Flags().BoolP("help", "h", false, "Print help")
	idl_initCmd.Flags().Bool("non-canonical", false, "Create non-canonical metadata account (third-party metadata)")
	idl_initCmd.Flags().String("priority-fee", "", "")
	idl_initCmd.MarkFlagRequired("filepath")
	idlCmd.AddCommand(idl_initCmd)
}
