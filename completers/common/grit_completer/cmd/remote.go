package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var remoteCmd = &cobra.Command{
	Use:     "remote",
	Short:   "List remotes, or add one",
	GroupID: "remote",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remoteCmd).Standalone()

	remoteCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	remoteCmd.Flags().BoolP("verbose", "v", false, "Show fetch and push URLs (like `git remote -v`)")
	rootCmd.AddCommand(remoteCmd)
}
