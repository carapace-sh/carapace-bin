package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var awsProfileCmd = &cobra.Command{
	Use:   "aws-profile",
	Short: "Generate AWS config profiles by syncing with your integrated AWS IAM Identity Center account(s). Other profiles in the config file are left untouched.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(awsProfileCmd).Standalone()

	awsProfileCmd.Flags().String("aws-sso-region", "", "AWS region for SSO. Auto-detected from cluster if not specified.")
	awsProfileCmd.Flags().Bool("dry-run", false, "Print the configuration that will be applied without modifying the AWS config file.")
	awsProfileCmd.Flags().Bool("no-dry-run", false, "Print the configuration that will be applied without modifying the AWS config file.")
	awsProfileCmd.Flag("no-dry-run").Hidden = true
	rootCmd.AddCommand(awsProfileCmd)
}
