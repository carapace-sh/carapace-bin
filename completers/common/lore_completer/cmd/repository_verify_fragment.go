package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_verify_fragmentCmd = &cobra.Command{
	Use:   "fragment",
	Short: "Verify a specific fragment in the local store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_verify_fragmentCmd).Standalone()

	repository_verify_fragmentCmd.Flags().String("context", "", "Context part of the address to verify")
	repository_verify_fragmentCmd.Flags().Bool("heal", false, "Attempt to heal if verification fails (remote only)")
	repository_verify_fragmentCmd.Flags().BoolP("help", "h", false, "Print help")
	repository_verifyCmd.AddCommand(repository_verify_fragmentCmd)

	carapace.Gen(repository_verify_fragmentCmd).PositionalCompletion(
		carapace.ActionValues(), // hash
	)
}
