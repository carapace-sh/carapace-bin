package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Print running jobs",
	Long:  "https://fishshell.com/docs/current/cmds/jobs.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("c", "c", false, "print command name")
	rootCmd.Flags().Bool("command", false, "print command name")
	rootCmd.Flags().BoolS("g", "g", false, "print group ID")
	rootCmd.Flags().Bool("group", false, "print group ID")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().BoolS("l", "l", false, "print only last job")
	rootCmd.Flags().Bool("last", false, "print only last job")
	rootCmd.Flags().BoolS("p", "p", false, "print process IDs")
	rootCmd.Flags().Bool("pid", false, "print process IDs")
	rootCmd.Flags().BoolS("q", "q", false, "query jobs")
	rootCmd.Flags().Bool("query", false, "query jobs")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		shell.ActionJobSpecs(),
	)
}
