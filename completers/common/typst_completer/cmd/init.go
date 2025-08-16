package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/typst_completer/cmd/common"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new project from a template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().BoolP("help", "h", false, "Print help (see a summary with '-h')")

	common.AddPackageFlags(initCmd)

	rootCmd.AddCommand(initCmd)

	// TODO: PositionalCompletion for templates
}
