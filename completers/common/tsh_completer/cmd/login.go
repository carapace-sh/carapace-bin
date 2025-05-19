package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to a cluster and retrieve the session certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()

	loginCmd.Flags().String("browser", "", "Set to 'none' to suppress browser opening on login")
	loginCmd.Flags().StringP("format", "f", "", "Identity format: file, openssh (for OpenSSH compatibility) or kubernetes (for kubeconfig)")
	loginCmd.Flags().String("kube-cluster", "", "Name of the Kubernetes cluster to login to")
	loginCmd.Flags().Bool("no-overwrite", false, "Whether to overwrite the existing identity file.")
	loginCmd.Flags().Bool("no-request-nowait", false, "Finish without waiting for request resolution")
	loginCmd.Flags().Bool("no-verbose", false, "Show extra status information")
	loginCmd.Flags().String("out", "", "Identity output")
	loginCmd.Flags().Bool("overwrite", false, "Whether to overwrite the existing identity file.")
	loginCmd.Flags().String("request-id", "", "Login with the roles requested in the given request")
	loginCmd.Flags().Bool("request-nowait", false, "Finish without waiting for request resolution")
	loginCmd.Flags().String("request-reason", "", "Reason for requesting additional roles")
	loginCmd.Flags().String("request-reviewers", "", "Suggested reviewers for role request")
	loginCmd.Flags().String("request-roles", "", "Request one or more extra roles")
	loginCmd.Flags().BoolP("verbose", "v", false, "Show extra status information")
	loginCmd.Flag("no-overwrite").Hidden = true
	loginCmd.Flag("no-request-nowait").Hidden = true
	loginCmd.Flag("no-verbose").Hidden = true
	rootCmd.AddCommand(loginCmd)

	carapace.Gen(loginCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("file", "openssh", "kubernetes"),
	})
}
