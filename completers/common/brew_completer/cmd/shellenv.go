package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shellenvCmd = &cobra.Command{
	Use:   "shellenv",
	Short: "Print export statements that, when run in a shell, will add Homebrew to the environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellenvCmd).Standalone()

	shellenvCmd.Flags().Bool("debug", false, "Display any debugging information.")
	shellenvCmd.Flags().Bool("help", false, "Show this message.")
	shellenvCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	shellenvCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(shellenvCmd)
}
