package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current repository status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	statusCmd.Flags().Bool("check-dirty", false, "Verify already-dirty files against the filesystem without a full scan")
	statusCmd.Flags().Bool("count", false, "Count directories and files (staged state if present, else current revision; view-filtered)")
	statusCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	statusCmd.Flags().Bool("reset", false, "Drop the existing staged anchor before computing status. Combine with --scan to scan from a clean slate")
	statusCmd.Flags().Bool("revision-only", false, "Only show revision info, skip all diffs")
	statusCmd.Flags().Bool("scan", false, "Walk the filesystem under the given paths and reconcile every file against the current revision")
	statusCmd.Flags().String("targets", "", "Path to a targets file")
	statusCmd.Flags().Bool("unstaged", false, "Alias for --scan (backward compatibility)")
	statusCmd.Flag("unstaged").Hidden = true
	rootCmd.AddCommand(statusCmd)

	carapace.Gen(statusCmd).FlagCompletion(carapace.ActionMap{
		"targets": carapace.ActionFiles(),
	})

	carapace.Gen(statusCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
