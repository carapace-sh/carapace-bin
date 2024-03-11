package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var commandsCmd = &cobra.Command{
	Use:     "commands",
	Short:   "Show lists of built-in and external commands",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commandsCmd).Standalone()

	commandsCmd.Flags().Bool("debug", false, "Display any debugging information.")
	commandsCmd.Flags().Bool("help", false, "Show this message.")
	commandsCmd.Flags().Bool("include-aliases", false, "Include aliases of internal commands.")
	commandsCmd.Flags().Bool("quiet", false, "List only the names of commands without category headers.")
	commandsCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(commandsCmd)
}
