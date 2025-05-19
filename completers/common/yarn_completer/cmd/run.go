package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Run a script defined in the package.json",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().BoolP("binaries-only", "B", false, "Ignore any user defined scripts and only check for binaries")
	runCmd.Flags().Bool("inspect", false, "Forwarded to the underlying Node process when executing a binary")
	runCmd.Flags().Bool("inspect-brk", false, "Forwarded to the underlying Node process when executing a binary")
	runCmd.Flags().BoolP("top-level", "T", false, "Check the root workspace for scripts and/or binaries instead of the current one")
	rootCmd.AddCommand(runCmd)

	// TODO positional completion
}
