package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var commandCmd = &cobra.Command{
	Use:     "command",
	Short:   "Display the path to the file being used when invoking `brew` <cmd>",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commandCmd).Standalone()

	commandCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	commandCmd.Flags().BoolP("help", "h", false, "Show this message.")
	commandCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	commandCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(commandCmd)

	carapace.Gen(commandCmd).PositionalAnyCompletion(
		carapace.ActionCommands(rootCmd),
	)
}
