package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Show an overview of the workspace state",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "inspection",
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	statusCmd.Flags().BoolS("f", "f", false, "Also list the files changed by each commit")
	statusCmd.Flags().Bool("files", false, "Also list the files changed by each commit")
	statusCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	statusCmd.Flags().Bool("no-hint", false, "Disable hints about available commands at the end of output")
	statusCmd.Flags().BoolP("refresh-prs", "r", false, "Forces a sync of pull requests from the forge before showing status")
	statusCmd.Flags().Bool("short", false, "")
	statusCmd.Flags().BoolP("upstream", "u", false, "Show detailed list of upstream commits that haven't been integrated yet")
	statusCmd.Flags().BoolP("verbose", "v", false, "Show verbose output with commit author and timestamp")
	statusCmd.Flag("files").Hidden = true
	statusCmd.Flag("short").Hidden = true
	rootCmd.AddCommand(statusCmd)
}
