package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docsCmd = &cobra.Command{
	Use:     "docs",
	Short:   "Open Homebrew's online documentation at <https://docs.brew.sh> in a browser",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docsCmd).Standalone()

	docsCmd.Flags().Bool("debug", false, "Display any debugging information.")
	docsCmd.Flags().Bool("help", false, "Show this message.")
	docsCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	docsCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(docsCmd)
}
