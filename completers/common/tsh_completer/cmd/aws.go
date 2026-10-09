package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Access AWS API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(awsCmd).Standalone()

	awsCmd.Flags().String("app", "", "Optional Name of the AWS application to use if logged into multiple.")
	awsCmd.Flags().String("aws-role", "", "(For AWS CLI access only) Amazon IAM role ARN or role name.")
	awsCmd.Flags().BoolP("endpoint-url", "e", false, "Run local proxy to serve as an AWS endpoint URL. If not specified, local proxy serves as an HTTPS proxy.")
	awsCmd.Flags().String("exec", "", "Execute different commands (e.g. terraform) under Teleport credentials.")
	awsCmd.Flags().Bool("no-endpoint-url", false, "Run local proxy to serve as an AWS endpoint URL. If not specified, local proxy serves as an HTTPS proxy.")
	awsCmd.Flag("endpoint-url").Hidden = true
	awsCmd.Flag("no-endpoint-url").Hidden = true
	rootCmd.AddCommand(awsCmd)
}
