package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var enterCmd = &cobra.Command{
	Use:     "enter [OPTION…] INSTANCE COMMAND [ARGUMENT…",
	Short:   "Run a command inside a running sandbox",
	GroupID: "run",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enterCmd).Standalone()

	enterCmd.Flags().BoolP("help", "h", false, "Show help options")
	enterCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	enterCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(enterCmd)

	// TODO positional
}
