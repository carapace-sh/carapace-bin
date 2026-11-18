package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var gistLogsCmd = &cobra.Command{
	Use:     "gist-logs",
	Short:   "Upload logs for a failed build of <formula> to a new Gist",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gistLogsCmd).Standalone()

	gistLogsCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	gistLogsCmd.Flags().BoolP("help", "h", false, "Show this message.")
	gistLogsCmd.Flags().BoolP("new-issue", "n", false, "Automatically create a new issue in the appropriate GitHub repository after creating the Gist.")
	gistLogsCmd.Flags().BoolP("private", "p", false, "The Gist will be marked private and will not appear in listings but will be accessible with its link.")
	gistLogsCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	gistLogsCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	gistLogsCmd.Flags().Bool("with-hostname", false, "Include the hostname in the Gist.")
	rootCmd.AddCommand(gistLogsCmd)

	carapace.Gen(gistLogsCmd).PositionalAnyCompletion(
		action.ActionList(gistLogsCmd).FilterArgs(),
	)
}
