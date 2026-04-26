package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stackCmd = &cobra.Command{
	Use:     "stack",
	Short:   "Install the official stack extension",
	GroupID: "extension",
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stackCmd).Standalone()

	rootCmd.AddCommand(stackCmd)
}
