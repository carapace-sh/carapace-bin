package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:    "completion [options] {bash|zsh|fish|powershell}",
	Short:  "Generate shell autocompletions",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionCmd).Standalone()

	completionCmd.Flags().StringP("file", "f", "", "Output the completion to file rather than stdout.")
	completionCmd.Flags().Bool("no-desc", false, "Don't include descriptions in the completion output.")
	rootCmd.AddCommand(completionCmd)
}
