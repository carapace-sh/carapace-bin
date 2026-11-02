package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var issue_artifactCmd = &cobra.Command{
	Use:     "artifact <command>",
	Short:   "Work with issue artifacts (preview)",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_artifactCmd).Standalone()

	issueCmd.AddCommand(issue_artifactCmd)
}
