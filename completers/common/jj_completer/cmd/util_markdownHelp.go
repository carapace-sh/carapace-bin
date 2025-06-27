package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var util_markdownHelpCmd = &cobra.Command{
	Use:   "markdown-help",
	Short: "Print the CLI help for all subcommands in Markdown",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(util_markdownHelpCmd).Standalone()

	util_markdownHelpCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	utilCmd.AddCommand(util_markdownHelpCmd)
}
