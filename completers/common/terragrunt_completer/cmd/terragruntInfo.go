package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var terragruntInfoCmd = &cobra.Command{
	Use:     "terragrunt-info",
	Short:   "Emits limited terragrunt state on stdout and exits",
	GroupID: "terragrunt",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terragruntInfoCmd).Standalone()

	rootCmd.AddCommand(terragruntInfoCmd)
}
