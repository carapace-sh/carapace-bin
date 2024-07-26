package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [--debug] [--] [<shell>]",
	Short: "Dump the shell completion script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionCmd).Standalone()

	completionCmd.Flags().Bool("debug", false, "Tail the completion debug log")
	rootCmd.AddCommand(completionCmd)
}
