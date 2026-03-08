package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repoInfoCmd = &cobra.Command{
	Use:   "info [repo-spec]...",
	Short: "show detailed info about repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoInfoCmd).Standalone()

	repoInfoCmd.Flags().Bool("all", false, "show information about all known repositories")
	repoInfoCmd.Flags().Bool("disabled", false, "show information about disabled repositories")
	repoInfoCmd.Flags().Bool("enabled", false, "show information about enabled repositories")
	repoInfoCmd.Flags().Bool("json", false, "show information in JSON format")

	repoCmd.AddCommand(repoInfoCmd)
}
