package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildSystemTargetsCmd = &cobra.Command{
	Use:   "build-system-targets",
	Short: "Print list of build system targets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildSystemTargetsCmd).Standalone()
	rootCmd.AddCommand(buildSystemTargetsCmd)
}
