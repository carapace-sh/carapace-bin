package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devenvCmd = &cobra.Command{
	Use:   "devenv",
	Short: "Run commands in developer environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devenvCmd).Standalone()

	devenvCmd.Flags().StringS("C", "C", "", "Path to build directory")
	devenvCmd.Flags().String("dump", "", "Only print required environment (Since 0.62.0) Takes an optional file path (Since 1.1.0)")
	devenvCmd.Flags().String("dump-format", "", "Format used with --dump (Since 1.1.0)")
	devenvCmd.Flags().StringP("workdir", "w", "", "Directory to cd into before running (default: builddir, Since 1.0.0)")
	applyCommonFlagCompletion(devenvCmd)
	rootCmd.AddCommand(devenvCmd)
}
