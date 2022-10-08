package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_setCredentialsCmd = &cobra.Command{
	Use:   "set-credentials",
	Short: "Set a user entry in kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCredentialsCmd).Standalone()
	config_setCredentialsCmd.Flags().String("auth-provider", "", "Auth provider for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().StringSlice("auth-provider-arg", []string{}, "'key=value' arguments for the auth provider")
	config_setCredentialsCmd.Flags().Bool("embed-certs", false, "Embed client cert/key for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().String("exec-api-version", "", "API version of the exec credential plugin for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().StringSlice("exec-arg", []string{}, "New arguments for the exec credential plugin command for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().String("exec-command", "", "Command for the exec credential plugin for the user entry in kubeconfig")
	config_setCredentialsCmd.Flags().StringArray("exec-env", []string{}, "'key=value' environment values for the exec credential plugin")
	config_setCredentialsCmd.Flag("embed-certs").NoOptDefVal = "true"
	configCmd.AddCommand(config_setCredentialsCmd)
}
