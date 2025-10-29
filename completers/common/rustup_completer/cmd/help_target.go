package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_targetCmd = &cobra.Command{
	Use:   "target",
	Short: "Modify a toolchain's supported targets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_targetCmd).Standalone()

	helpCmd.AddCommand(help_targetCmd)
}
