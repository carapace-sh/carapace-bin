package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
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

	bumpRevisionCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	bumpRevisionCmd.Flags().BoolP("dry-run", "n", false, "Print what would be done rather than doing it.")
	bumpRevisionCmd.Flags().BoolP("help", "h", false, "Show this message.")
	bumpRevisionCmd.Flags().String("message", "", "Append <message> to the default commit message.")
	bumpRevisionCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	bumpRevisionCmd.Flags().Bool("remove-bottle-block", false, "Remove the bottle block in addition to bumping the revision.")
	bumpRevisionCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	bumpRevisionCmd.Flags().Bool("write-only", false, "Make the expected file modifications without taking any Git actions.")
	rootCmd.AddCommand(bumpRevisionCmd)

	carapace.Gen(bumpRevisionCmd).PositionalAnyCompletion(
		brew.ActionAllFormulae(),
	)
}
