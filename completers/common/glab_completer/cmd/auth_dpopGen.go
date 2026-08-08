package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/ssh"
	"github.com/spf13/cobra"
)

var auth_dpopGenCmd = &cobra.Command{
	Use:   "dpop-gen",
	Short: "Generate a DPoP (demonstrating-proof-of-possession) proof JWT. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_dpopGenCmd).Standalone()

	auth_dpopGenCmd.Flags().String("hostname", "", "The hostname of the GitLab instance to authenticate with.")
	auth_dpopGenCmd.Flags().String("pat", "", "Personal access token (PAT) to generate a DPoP proof for. Defaults to the token set with 'glab auth login'.")
	auth_dpopGenCmd.Flags().StringP("private-key", "p", "", "Location of the private SSH key on the local system.")
	auth_dpopGenCmd.MarkFlagRequired("private-key")
	authCmd.AddCommand(auth_dpopGenCmd)

	carapace.Gen(auth_dpopGenCmd).FlagCompletion(carapace.ActionMap{
		"hostname": net.ActionHosts(),
		"private-key": carapace.Batch(
			ssh.ActionPrivateKeys(),
			carapace.ActionFiles(),
		).ToA(),
	})
}
