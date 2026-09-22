package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var idl_createBufferCmd = &cobra.Command{
	Use:   "create-buffer",
	Short: "Create a buffer account for metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(idl_createBufferCmd).Standalone()

	idl_createBufferCmd.Flags().StringP("filepath", "f", "", "Path to the metadata file")
	idl_createBufferCmd.Flags().BoolP("help", "h", false, "Print help")
	idl_createBufferCmd.Flags().String("priority-fee", "", "Priority fees in micro-lamports per compute unit")
	idl_createBufferCmd.MarkFlagRequired("filepath")
	idlCmd.AddCommand(idl_createBufferCmd)
}
