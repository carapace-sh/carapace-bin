package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var label_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new label",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_createCmd).Standalone()
	label_createCmd.Flags().StringP("color", "c", "", "Color of the label, if not specified one will be selected at random")
	label_createCmd.Flags().StringP("description", "d", "", "Description of the label")
	labelCmd.AddCommand(label_createCmd)

	carapace.Gen(label_createCmd).FlagCompletion(carapace.ActionMap{
		"color": os.ActionHexColors(),
	})
}
