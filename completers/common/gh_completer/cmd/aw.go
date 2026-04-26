package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var awCmd = &cobra.Command{
	Use:     "aw",
	Short:   "Install the official aw extension",
	GroupID: "extension",
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(awCmd).Standalone()

	rootCmd.AddCommand(awCmd)
}
