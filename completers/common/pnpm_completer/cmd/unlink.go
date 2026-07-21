package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var unlinkCmd = &cobra.Command{
	Use:     "unlink",
	Aliases: []string{"dislink"},
	Short:   "Removes the link created by `pnpm link` and reinstalls package if it is saved in `package.json`",
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unlinkCmd).Standalone()

	unlinkCmd.Flags().BoolP("recursive", "r", false, "Unlink in every package found in subdirectories or in every workspace package")
	rootCmd.AddCommand(unlinkCmd)

	carapace.Gen(unlinkCmd).PositionalAnyCompletion(
		pnpm.ActionDependencyNames(), // TODO complete only linked dependencies
	)
}
