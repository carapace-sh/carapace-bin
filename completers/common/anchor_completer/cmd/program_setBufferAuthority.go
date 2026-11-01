package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var program_setBufferAuthorityCmd = &cobra.Command{
	Use:   "set-buffer-authority",
	Short: "Set a new buffer authority",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(program_setBufferAuthorityCmd).Standalone()

	program_setBufferAuthorityCmd.Flags().BoolP("help", "h", false, "Print help")
	programCmd.AddCommand(program_setBufferAuthorityCmd)
}
