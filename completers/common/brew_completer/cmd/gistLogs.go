package cmd

import (
	"github.com/carapace-sh/carapace"
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

	gistLogsCmd.Flags().Bool("debug", false, "Display any debugging information.")
	gistLogsCmd.Flags().Bool("help", false, "Show this message.")
	gistLogsCmd.Flags().Bool("new-issue", false, "Automatically create a new issue in the appropriate GitHub repository after creating the Gist.")
	gistLogsCmd.Flags().Bool("private", false, "The Gist will be marked private and will not appear in listings but will be accessible with its link.")
	gistLogsCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	gistLogsCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	gistLogsCmd.Flags().Bool("with-hostname", false, "Include the hostname in the Gist.")
	rootCmd.AddCommand(gistLogsCmd)
}
