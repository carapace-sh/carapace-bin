package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dryRunCmd = &cobra.Command{
	Use:   "dry-run",
	Short: "Show what store paths would be built or downloaded, but otherwise do nothing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dryRunCmd).Standalone()
	rootCmd.AddCommand(dryRunCmd)
}
