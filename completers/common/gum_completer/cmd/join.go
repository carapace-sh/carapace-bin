package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gum_completer/cmd/common"
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

	common.AddFlagCompletion(joinCmd)
}
