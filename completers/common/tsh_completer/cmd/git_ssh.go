package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_sshCmd = &cobra.Command{
	Use:    "ssh",
	Short:  "Proxy Git commands using SSH.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_sshCmd).Standalone()

	git_sshCmd.Flags().String("github-org", "", "GitHub organization.")
	git_sshCmd.Flags().StringP("option", "o", "", "OpenSSH options in the format used in the configuration file.")
	git_sshCmd.MarkFlagRequired("github-org")
	gitCmd.AddCommand(git_sshCmd)
}
