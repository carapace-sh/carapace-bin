package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Opens a browser and retrieves your login from GitHub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_loginCmd).Standalone()

	git_loginCmd.Flags().Bool("force", false, "Force a login.")
	git_loginCmd.Flags().String("github-org", "", "GitHub organization.")
	git_loginCmd.Flags().Bool("no-force", false, "Force a login.")
	git_loginCmd.MarkFlagRequired("github-org")
	git_loginCmd.Flag("no-force").Hidden = true
	gitCmd.AddCommand(git_loginCmd)
}
