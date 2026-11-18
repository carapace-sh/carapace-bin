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

	docsCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	docsCmd.Flags().BoolP("help", "h", false, "Show this message.")
	docsCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	docsCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(docsCmd)
}
