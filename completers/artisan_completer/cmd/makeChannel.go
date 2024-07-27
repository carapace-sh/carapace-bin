package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeChannelCmd = &cobra.Command{
	Use:   "make:channel [-f|--force] [--] <name>",
	Short: "Create a new channel class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeChannelCmd).Standalone()

	makeChannelCmd.Flags().Bool("force", false, "Create the class even if the channel already exists")
	rootCmd.AddCommand(makeChannelCmd)
}
