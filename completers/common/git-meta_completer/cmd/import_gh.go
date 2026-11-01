package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var import_ghCmd = &cobra.Command{
	Use:   "gh",
	Short: "Import merged pull request metadata from GitHub using gh",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_ghCmd).Standalone()

	import_ghCmd.Flags().Bool("dry-run", false, "Show what would be imported without writing")
	import_ghCmd.Flags().Bool("force", false, "Reprocess PRs even when they were previously imported")
	import_ghCmd.Flags().BoolP("help", "h", false, "Print help")
	import_ghCmd.Flags().Bool("include-comments", false, "Import PR comments and review bodies")
	import_ghCmd.Flags().String("limit", "", "Maximum number of merged PRs to fetch")
	import_ghCmd.Flags().Bool("no-tags", false, "Skip release tag mapping")
	import_ghCmd.Flags().String("repo", "", "GitHub repository in OWNER/NAME form")
	import_ghCmd.Flags().String("since", "", "Only import PRs merged on or after this date (YYYY-MM-DD)")
	importCmd.AddCommand(import_ghCmd)

	carapace.Gen(import_ghCmd).FlagCompletion(carapace.ActionMap{
		"repo": gh.ActionOwnerRepositories(gh.HostOpts{}),
	})
}
