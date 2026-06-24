package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var refs_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify ref database consistency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refs_verifyCmd).Standalone()

	refs_verifyCmd.Flags().Bool("strict", false, "enable strict checking")
	refs_verifyCmd.Flags().Bool("verbose", false, "be verbose")
	refsCmd.AddCommand(refs_verifyCmd)
}
