package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var multiPackIndex_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify the contents of the MIDX file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(multiPackIndex_verifyCmd).Standalone()

	multiPackIndexCmd.AddCommand(multiPackIndex_verifyCmd)
}
