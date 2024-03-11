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

	commandCmd.Flags().Bool("debug", false, "Display any debugging information.")
	commandCmd.Flags().Bool("help", false, "Show this message.")
	commandCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	commandCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(commandCmd)
}
