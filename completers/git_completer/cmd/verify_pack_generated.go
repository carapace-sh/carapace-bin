package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var verify_packCmd = &cobra.Command{
	Use:     "verify-pack",
	Short:   "Validate packed Git archive files",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(verify_packCmd).Standalone()
	verify_packCmd.Flags().BoolP("stat-only", "s", false, "show statistics only")
	verify_packCmd.Flags().BoolP("verbose", "v", false, "verbose")
	rootCmd.AddCommand(verify_packCmd)
}
