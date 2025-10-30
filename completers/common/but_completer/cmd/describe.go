package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use:     "describe",
	Short:   "Edit the commit message of the specified commit",
	Aliases: []string{"desc"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(describeCmd).Standalone()

	describeCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(describeCmd)
}
