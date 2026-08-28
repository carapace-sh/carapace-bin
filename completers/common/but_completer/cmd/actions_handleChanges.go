package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var actions_handleChangesCmd = &cobra.Command{
	Use:   "handle-changes",
	Short: "Automatically handles the changes in the repository, creating a commit with the provided context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(actions_handleChangesCmd).Standalone()

	actions_handleChangesCmd.Flags().String("desc", "", "A context describing the changes that are currently uncommitted")
	actions_handleChangesCmd.Flags().String("description", "", "A context describing the changes that are currently uncommitted")
	actions_handleChangesCmd.Flags().String("handler", "simple", "Which handler is to be used for the operation. Different handles would have different behavior")
	actions_handleChangesCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	actions_handleChangesCmd.Flag("desc").Hidden = true
	actions_handleChangesCmd.MarkFlagRequired("desc")
	actions_handleChangesCmd.MarkFlagRequired("description")
	actionsCmd.AddCommand(actions_handleChangesCmd)
}