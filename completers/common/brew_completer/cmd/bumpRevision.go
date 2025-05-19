package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bumpRevisionCmd = &cobra.Command{
	Use:     "bump-revision",
	Short:   "Create a commit to increment the revision of <formula>",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bumpRevisionCmd).Standalone()

	bumpRevisionCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bumpRevisionCmd.Flags().Bool("dry-run", false, "Print what would be done rather than doing it.")
	bumpRevisionCmd.Flags().Bool("help", false, "Show this message.")
	bumpRevisionCmd.Flags().Bool("message", false, "Append <message> to the default commit message.")
	bumpRevisionCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bumpRevisionCmd.Flags().Bool("remove-bottle-block", false, "Remove the bottle block in addition to bumping the revision.")
	bumpRevisionCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	bumpRevisionCmd.Flags().Bool("write-only", false, "Make the expected file modifications without taking any Git actions.")
	rootCmd.AddCommand(bumpRevisionCmd)
}
