package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gum"
	"github.com/spf13/cobra"
)

var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Join text vertically or horizontally",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(joinCmd).Standalone()
	joinCmd.Flags().String("align", "", "Text alignment")
	joinCmd.Flags().Bool("horizontal", false, "Join (potentially multi-line) strings horizontally")
	joinCmd.Flags().Bool("vertical", false, "Join (potentially multi-line) strings vertically")
	rootCmd.AddCommand(joinCmd)

	carapace.Gen(joinCmd).FlagCompletion(carapace.ActionMap{
		"align": gum.ActionAlignments(),
	})
}
