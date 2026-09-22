package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var legacyIdl_writeBufferCmd = &cobra.Command{
	Use:   "write-buffer",
	Short: "[DEPRECATED] Write an IDL into a legacy buffer account. Use with set-buffer to upgrade",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(legacyIdl_writeBufferCmd).Standalone()

	legacyIdl_writeBufferCmd.Flags().StringP("filepath", "f", "", "")
	legacyIdl_writeBufferCmd.Flags().BoolP("help", "h", false, "Print help")
	legacyIdl_writeBufferCmd.Flags().String("priority-fee", "", "")
	legacyIdl_writeBufferCmd.MarkFlagRequired("filepath")
	legacyIdlCmd.AddCommand(legacyIdl_writeBufferCmd)
}
