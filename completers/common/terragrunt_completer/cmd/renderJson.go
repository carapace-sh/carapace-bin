package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var renderJsonCmd = &cobra.Command{
	Use:     "render-json",
	Short:   "Render the final terragrunt config as json",
	GroupID: "terragrunt",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(renderJsonCmd).Standalone()

	rootCmd.AddCommand(renderJsonCmd)
}
