package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var token_gitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "Show the GitLab token mise uses for a host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_gitlabCmd).Standalone()

	token_gitlabCmd.Flags().BoolP("help", "h", false, "Print help")
	token_gitlabCmd.Flags().Bool("unmask", false, "Show the full unmasked token")
	tokenCmd.AddCommand(token_gitlabCmd)
}
