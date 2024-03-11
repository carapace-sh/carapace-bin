package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/ssh"
	"github.com/spf13/cobra"
)

var logins_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a Gitea login",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logins_addCmd).Standalone()

	logins_addCmd.Flags().BoolP("insecure", "i", false, "Disable TLS verification")
	logins_addCmd.Flags().StringP("name", "n", "", "Login name")
	logins_addCmd.Flags().Bool("no-version-check", false, "Do not check version of Gitea instance")
	logins_addCmd.Flags().String("password", "", "Password for basic auth (will create token)")
	logins_addCmd.Flags().StringP("ssh-agent-key", "a", "", "Use SSH public key or SSH fingerprint to login (needs a running ssh-agent with ssh key loaded)")
	logins_addCmd.Flags().StringP("ssh-agent-principal", "c", "", "Use SSH certificate with specified principal to login (needs a running ssh-agent with certificate loaded)")
	logins_addCmd.Flags().StringP("ssh-key", "s", "", "Path to a SSH key/certificate to use, overrides auto-discovery")
	logins_addCmd.Flags().StringP("token", "t", "", "Access token. Can be obtained from Settings > Applications")
	logins_addCmd.Flags().StringP("url", "u", "", "Server URL")
	logins_addCmd.Flags().String("user", "", "User for basic auth (will create token)")
	loginsCmd.AddCommand(logins_addCmd)

	// TODO completion
	carapace.Gen(logins_addCmd).FlagCompletion(carapace.ActionMap{
		"ssh-key": carapace.Batch(
			ssh.ActionPrivateKeys(),
			carapace.ActionFiles(),
		).ToA(),
	})
}
