package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:     "branch",
	Short:   "List and manipulate virtual branches",
	Aliases: []string{"branches"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branchCmd).Standalone()

	branchCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(branchCmd)
}
