package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:     "help",
	Short:   "Display help information about Git",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interrogator].ID,
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	helpCmd.Flags().BoolP("all", "a", false, "Prints all the available commands on the standard output")
	helpCmd.Flags().BoolP("config", "c", false, "List all available configuration variables")
	helpCmd.Flags().Bool("developer-interfaces", false, "Print list of file formats, protocols and other developer interfaces documentation")
	helpCmd.Flags().BoolP("guides", "g", false, "Prints a list of the Git concept guides on the standard output")
	helpCmd.Flags().BoolP("info", "i", false, "Display manual page for the command in the info format")
	helpCmd.Flags().BoolP("man", "m", false, "Display manual page for the command in the man format")
	helpCmd.Flags().Bool("no-aliases", false, "Exclude the listing of configured aliases")
	helpCmd.Flags().Bool("no-external-commands", false, "Exclude the listing of external \"git-*\" commands found in the $PATH")
	helpCmd.Flags().Bool("user-interfaces", false, "Prints a list of the repository, command and file interfaces documentation")
	helpCmd.Flags().Bool("verbose", false, "Print description for all recognized commands")
	helpCmd.Flags().BoolP("web", "w", false, "Display manual page for the command in the web (HTML) format")
	rootCmd.AddCommand(helpCmd)

	carapace.Gen(helpCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionCommands(rootCmd),
			git.ActionDeveloperInterfaces(),
		).ToA(),
	)
}
