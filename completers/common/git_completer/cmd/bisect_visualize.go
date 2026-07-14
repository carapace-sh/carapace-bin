package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bisect_visualizeCmd = &cobra.Command{
	Use:     "visualize",
	Aliases: []string{"view"},
	Short:   "show bisect status in gitk",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisect_visualizeCmd).Standalone()
	bisect_visualizeCmd.Flags().Bool("p", false, "show patches")
	bisect_visualizeCmd.Flags().Bool("stat", false, "show diffstat")

	bisectCmd.AddCommand(bisect_visualizeCmd)

	carapace.Gen(bisect_visualizeCmd).DashAnyCompletion(
		carapace.ActionPositional(bisect_visualizeCmd),
	)
}
