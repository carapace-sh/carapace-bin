package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var apps_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Retrieve short-lived certificate for an app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_loginCmd).Standalone()

	apps_loginCmd.Flags().String("aws-role", "", "(For AWS CLI access only) Amazon IAM role ARN or role name.")
	apps_loginCmd.Flags().String("azure-identity", "", "(For Azure CLI access only) Azure managed identity name.")
	apps_loginCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect")
	apps_loginCmd.Flags().String("gcp-service-account", "", "(For GCP CLI access only) GCP service account name.")
	apps_loginCmd.Flags().Bool("no-quiet", false, "Quiet mode")
	apps_loginCmd.Flags().BoolP("quiet", "q", false, "Quiet mode")
	apps_loginCmd.Flag("no-quiet").Hidden = true
	appsCmd.AddCommand(apps_loginCmd)

	// TODO flag completion
	carapace.Gen(apps_loginCmd).FlagCompletion(carapace.ActionMap{
		"cluster": tsh.ActionClusters(),
	})
}
