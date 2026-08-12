package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current repository status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_statusCmd).Standalone()

	repository_statusCmd.Flags().Bool("check-dirty", false, "Verify already-dirty files against the filesystem without a full scan")
	repository_statusCmd.Flags().Bool("count", false, "Count directories and files (staged state if present, else current revision; view-filtered)")
	repository_statusCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	repository_statusCmd.Flags().Bool("reset", false, "Drop the existing staged anchor before computing status. Combine with --scan to scan from a clean slate")
	repository_statusCmd.Flags().Bool("revision-only", false, "Only show revision info, skip all diffs")
	repository_statusCmd.Flags().Bool("scan", false, "Walk the filesystem under the given paths and reconcile every file against the current revision")
	repository_statusCmd.Flags().String("targets", "", "Path to a targets file")
	repository_statusCmd.Flags().Bool("unstaged", false, "Alias for --scan (backward compatibility)")
	repository_statusCmd.Flag("unstaged").Hidden = true
	repositoryCmd.AddCommand(repository_statusCmd)

	carapace.Gen(repository_statusCmd).FlagCompletion(carapace.ActionMap{
		"targets": carapace.ActionFiles(),
	})

	carapace.Gen(repository_statusCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
