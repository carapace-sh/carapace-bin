package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var ask_gitCmd = &cobra.Command{
	Use:   "git <prompt>",
	Short: "Generate Git commands from natural language (Experimental).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ask_gitCmd).Standalone()

	askCmd.AddCommand(ask_gitCmd)
}
