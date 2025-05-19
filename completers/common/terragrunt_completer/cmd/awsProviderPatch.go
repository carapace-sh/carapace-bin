package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var awsProviderPatchCmd = &cobra.Command{
	Use:     "aws-provider-patch",
	Short:   "Overwrite settings on nested AWS providers to work around a Terraform bug",
	GroupID: "terragrunt",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(awsProviderPatchCmd).Standalone()

	rootCmd.AddCommand(awsProviderPatchCmd)
}
