package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flagCaskroomCmd = &cobra.Command{
	Use:   "caskroom",
	Short: "Display Homebrew's Caskroom path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flagCaskroomCmd).Standalone()

	flagCaskroomCmd.Flags().Bool("debug", false, "Display any debugging information.")
	flagCaskroomCmd.Flags().Bool("help", false, "Show this message.")
	flagCaskroomCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	flagCaskroomCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
}
