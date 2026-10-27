package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mvCmd = &cobra.Command{
	Use:     "mv",
	Short:   "Rename or move tracked paths",
	GroupID: "changes",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mvCmd).Standalone()

	mvCmd.Flags().Bool("dry-run", false, "List moves that would happen without changing the index or work tree")
	mvCmd.Flags().BoolP("force", "f", false, "Replace existing files at the destination")
	mvCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(mvCmd)

	carapace.Gen(mvCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
