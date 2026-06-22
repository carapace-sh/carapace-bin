package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/common"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new Rails application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newCmd).Standalone()

	carapace.Gen(newCmd).PositionalCompletion(
		carapace.ActionValues().Usage("application name"),
	)

	newCmd.Flags().BoolP("help", "h", false, "Show help")

	common.AddAppGeneratorFlags(newCmd)
}
