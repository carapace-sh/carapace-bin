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

	flagCaskroomCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	flagCaskroomCmd.Flags().BoolP("help", "h", false, "Show this message.")
	flagCaskroomCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	flagCaskroomCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
}
