package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Validate the MD5/SHA1 on a drive image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifyCmd).Standalone()

	verifyCmd.Flags().BoolP("fix", "f", false, "fix the SHA-1 if it is incorrect")
	verifyCmd.Flags().StringP("input", "i", "", "input file name")
	verifyCmd.Flags().StringP("inputparent", "ip", "", "parent file name for input CHD")
	rootCmd.AddCommand(verifyCmd)

	carapace.Gen(verifyCmd).FlagCompletion(carapace.ActionMap{
		"input":       carapace.ActionFiles(),
		"inputparent": carapace.ActionFiles(),
	})
}
