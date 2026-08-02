package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repoInfoCmd = &cobra.Command{
	Use:   "info [options] [<repo-spec>...]",
	Short: "print details about repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoInfoCmd).Standalone()

	repoInfoCmd.Flags().Bool("all", false, "Show all repositories")
	repoInfoCmd.Flags().Bool("disabled", false, "Show disabled repositories")
	repoInfoCmd.Flags().Bool("enabled", false, "Show enabled repositories (default)")
	repoInfoCmd.Flags().Bool("json", false, "Request json output format")

	repoCmd.AddCommand(repoInfoCmd)
}
